package server

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	epb "github.com/golang/protobuf/ptypes/empty"
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

func (s *Server) GetMatch(ctx context.Context, in *apb.CharonRpc_MatchRequest) (*apb.CharonRpc_MatchResponse, error) {
	if in.Match == nil {
		return nil, grpc.Errorf(codes.FailedPrecondition, "must specify match")
	}

	res, err := s.Client.Region(in.Match.Region).Match(in.Match.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get match: %v", err)
	}

	mpb := translate.Match(res)

	var summoners []*apb.SummonerId
	for _, p := range mpb.ParticipantInfo {
		summoners = append(summoners, &apb.SummonerId{
			Region: mpb.Region,
			Id:     p.Identity.SummonerId,
		})
	}

	return &apb.CharonRpc_MatchResponse{
		MatchInfo: mpb,
		Summoners: summoners,
	}, nil
}

func (s *Server) GetMatchList(ctx context.Context, in *apb.CharonRpc_MatchListRequest) (*apb.CharonRpc_MatchListResponse, error) {
	if in.Summoner == nil {
		return nil, grpc.Errorf(codes.FailedPrecondition, "must specify summoner")
	}

	res, err := s.Client.Region(in.Summoner.Region).MatchList(in.Summoner.Id, in.Queues, in.Seasons)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get match list: %v", err)
	}

	var matches []*apb.CharonRpc_MatchListResponse_MatchInfo
	for _, match := range res.Matches {
		ts := &tspb.Timestamp{
			Seconds: int64(match.Timestamp / 1000),
			Nanos:   int32(match.Timestamp%1000) * 1E7,
		}
		matches = append(matches, &apb.CharonRpc_MatchListResponse_MatchInfo{
			MatchId: &apb.MatchId{
				Region: in.Summoner.Region,
				Id:     match.MatchId,
			},
			Timestamp: ts,
		})
	}

	return &apb.CharonRpc_MatchListResponse{
		Matches: matches,
	}, nil
}

func (s *Server) GetRankings(ctx context.Context, in *apb.CharonRpc_RankingsRequest) (*apb.CharonRpc_RankingsResponse, error) {
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

	return &apb.CharonRpc_RankingsResponse{
		Rankings: rankings,
	}, nil
}

func (s *Server) GetStaticChampions(
	ctx context.Context, in *apb.CharonRpc_StaticChampionsRequest,
) (*apb.CharonRpc_StaticChampionsResponse, error) {
	if in.Version == "" {
		return nil, grpc.Errorf(codes.FailedPrecondition, "must specify version")
	}

	res, err := s.Client.Region(in.Region).StaticChampions(in.Locale, in.Version)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get static champions: %v", err)
	}

	return &apb.CharonRpc_StaticChampionsResponse{
		StaticChampions: translate.StaticChampions(res),
	}, nil
}

func (s *Server) GetStaticVersions(
	ctx context.Context, in *epb.Empty,
) (*apb.CharonRpc_StaticVersionsResponse, error) {
	res, err := s.Client.Region(apb.Region_NA).StaticVersions()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "could not get static versions: %v", err)
	}

	return &apb.CharonRpc_StaticVersionsResponse{
		Versions: res,
	}, nil
}
