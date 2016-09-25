package server

import (
	"fmt"
	"strconv"
	"time"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot"
	"github.com/asunaio/charon/util"
	"github.com/golang/protobuf/ptypes"
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
	res, err := s.Client.Region(in.Region).League(in.SummonerIds)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get match list: %v", err)
	}

	rankings := map[uint64]*apb.Ranking{}
	for _, summoner := range in.SummonerIds {
		// summoner id
		sid := strconv.FormatUint(summoner, 10)
		dtos := res[sid]
		if dtos == nil {
			// we fucked
			return nil, grpc.Errorf(codes.FailedPrecondition, "Riot missing dtos in ranking response")
		}

		// find dto
		var dto *riot.LeagueDto
		for _, x := range dtos {
			if x.Queue == riot.QueueSolo5x5 {
				dto = x
				break
			}
		}
		if dto == nil {
			// unranked
			continue
		}

		// Find player ranking
		tier := dto.Tier
		var entry *riot.LeagueEntryDto
		for _, x := range dto.Entries {
			if x.PlayerOrTeamID == strconv.FormatUint(summoner, 10) {
				entry = x
				break
			}
		}

		if entry == nil {
			// should not happen
			return nil, grpc.Errorf(codes.FailedPrecondition, "no summoner %d for league %s of %s", summoner, dto.Name, dto.Tier)
		}

		rank, err := util.ParseRank(tier, entry.Division)
		if err != nil {
			return nil, fmt.Errorf("invalid rank: %v", err)
		}

		now, err := ptypes.TimestampProto(time.Now())
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "could not create now timestamp: %v", err)
		}

		rankings[summoner] = &apb.Ranking{
			Summoner: &apb.SummonerId{
				Region: in.Region,
				Id:     summoner,
			},
			Rank: rank,
			Time: now,
		}
	}

	return &apb.CharonRankingsResponse{
		Payload: &apb.CharonRankingsResponse_Payload{
			Rankings: rankings,
		},
	}, nil
}
