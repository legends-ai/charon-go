package server

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/client"
	"github.com/asunaio/charon/translate"
	"github.com/asunaio/charon/util"
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

	mpb, err := translate.Match(res)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not translate riot match response to Charon Match format: %v", err)
	}

	return &apb.CharonMatchResponse{
		Payload: &apb.CharonMatchResponse_Payload{
			MatchInfo: mpb,
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

	var matches []*apb.CharonMatchListResponse_MatchInfo
	for _, match := range res.Matches {
		ts := &tspb.Timestamp{
			Seconds: int64(match.Timestamp / 1000),
			Nanos:   int32(match.Timestamp%1000) * 1E7,
		}
		matches = append(matches, &apb.CharonMatchListResponse_MatchInfo{
			MatchId: &apb.MatchId{
				Region: in.Summoner.Region,
				Id:     match.MatchId,
			},
			Timestamp: ts,
		})
	}

	return &apb.CharonMatchListResponse{
		Matches: matches,
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
			// unranked
			continue
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
