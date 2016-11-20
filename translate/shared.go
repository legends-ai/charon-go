package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func parseImage(image models.StaticImage) *apb.CharonData_Static_Image {
	return &apb.CharonData_Static_Image{
		Full:   image.Full,
		Sprite: image.Sprite,
		Group:  image.Group,
		X:      image.X,
		Y:      image.Y,
		W:      image.W,
		H:      image.H,
	}
}
