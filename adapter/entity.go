package adapter

type GEOPos struct {
	Mem      string
	Distance float64
	Position Pos
}

type Pos struct {
	Longitude float64
	Latitude  float64
}
