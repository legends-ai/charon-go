package server

import (
	"strconv"

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
	res, err := s.Client.Region(in.Region).Match(strconv.Itoa(int(in.MatchId.Id)))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get match: %v", err)
	}

	var summoners []*apb.SummonerId
	for _, p := range res.ParticipantIdentities {
		summoners = append(summoners, &apb.SummonerId{
			Region: in.Region,
			Id:     uint32(p.Player.SummonerID),
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
	return nil, nil
}

func (s *Server) GetRankings(ctx context.Context, in *apb.CharonRankingsRequest) (*apb.CharonRankingsResponse, error) {
	return nil, nil
}
