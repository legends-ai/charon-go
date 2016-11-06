package models

type RiotStaticImage struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	H      uint32 `json:"h"`
	Sprite string `json:"sprite"`
	W      uint32 `json:"w"`
	X      uint32 `json:"x"`
	Y      uint32 `json:"y"`
}
