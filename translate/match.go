package translate

import (
	"fmt"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func Match(match *models.RiotMatch) (*apb.Charon_Match, error) {
	//cmpb := &apb.Charon_Match{}
	fmt.Printf("%+v", match)
	return nil, nil
}
