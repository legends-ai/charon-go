package server

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GetMatch(ctx context.Context, in *apb.CharonMatchRequest) (*apb.CharonMatchResponse, error) {
	return nil, nil
}

func (s *Server) GetMatchList(ctx context.Context, in *apb.CharonMatchListRequest) (*apb.CharonMatchListResponse, error) {
	return nil, nil
}

func (s *Server) GetRankings(ctx context.Context, in *apb.CharonRankingsRequest) (*apb.CharonRankingsResponse, error) {
	return nil, nil
}
