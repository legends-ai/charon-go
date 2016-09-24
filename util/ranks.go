package util

import (
	"fmt"

	apb "github.com/asunaio/charon/gen-go/asuna"
)

const (
	TierChallenger = "CHALLENGER"
	TierMaster     = "MASTER"
	TierDiamond    = "DIAMOND"
	TierPlatinum   = "PLATINUM"
	TierGold       = "GOLD"
	TierSilver     = "SILVER"
	TierBronze     = "BRONZE"

	DivisionI   = "I"
	DivisionII  = "II"
	DivisionIII = "III"
	DivisionIV  = "IV"
	DivisionV   = "V"
)

// ParseRank parses a tier and division to return a Rank.
func ParseRank(tier, division string) (*apb.Rank, error) {
	ti := parseTier(tier)
	if ti == 0 {
		return nil, fmt.Errorf("invalid tier %s", tier)
	}
	di := parseDivision(division)
	if di == 0 {
		return nil, fmt.Errorf("invalid division %s", division)
	}
	return &apb.Rank{
		Tier:     ti,
		Division: di,
	}, nil
}

func parseTier(s string) uint32 {
	switch s {
	case TierChallenger:
		return 0x70
	case TierMaster:
		return 0x60
	case TierDiamond:
		return 0x50
	case TierPlatinum:
		return 0x40
	case TierGold:
		return 0x30
	case TierSilver:
		return 0x20
	case TierBronze:
		return 0x10
	default:
		return 0
	}
}

func parseDivision(s string) uint32 {
	switch s {
	case DivisionI:
		return 0x50
	case DivisionII:
		return 0x40
	case DivisionIII:
		return 0x30
	case DivisionIV:
		return 0x20
	case DivisionV:
		return 0x10
	default:
		return 0
	}
}
