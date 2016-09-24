package server

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Server struct {
	Client *riot.Client `inject:"t"`
}

func (s *Server) GetMatch(ctx context.Context, in *apb.CharonMatchRequest) (*apb.CharonMatchResponse, error) {
	if in.Match == nil {
		return nil, grpc.Errorf(codes.FailedPrecondition, "must specify match")
	}

	res, err := s.Client.Region(in.Match.Region).Match(in.Match.Id)

	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get match: %v", err)
	}

	var summoners []*apb.SummonerId
	for _, p := range res.ParticipantIdentities {
		summoners = append(summoners, &apb.SummonerId{
			Region: in.Match.Region,
			Id:     p.Player.SummonerID,
		})
	}

	return &apb.CharonMatchResponse{
		Payload: &apb.CharonMatchResponse_Payload{
			MatchVersion: res.MatchVersion,
			QueueType:    res.QueueType,
			RawJson:      res.RawJSON,
			Summoners:    summoners,
		},
	}, nil
}

func (s *Server) GetMatchList(ctx context.Context, in *apb.CharonMatchListRequest) (*apb.CharonMatchListResponse, error) {
	if in.Summoner == nil {
		return nil, grpc.Errorf(codes.FailedPrecondition, "must specify summoner")
	}

	res, err := s.Client.Region(in.Summoner.Region).MatchList(in.Summoner.Id, in.Queues, in.Seasons)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get match list: %v", err)
	}

	var matches []*apb.MatchId
	for _, match := range res.Matches {
		matches = append(matches, &apb.MatchId{
			Region: in.Summoner.Region,
			Id:     match.MatchId,
		})
	}

	return &apb.CharonMatchListResponse{
		Payload: &apb.CharonMatchListResponse_Payload{
			Matches: matches,
		},
	}, nil
}

func (s *Server) GetRankings(ctx context.Context, in *apb.CharonRankingsRequest) (*apb.CharonRankingsResponse, error) {
	return nil, nil
}
