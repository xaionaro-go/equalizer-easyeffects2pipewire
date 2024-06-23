package equalizer

type BandType int

const (
	UndefinedBandType = BandType(iota)
	BandTypeLowShelf
	BandTypeHighShelf
	BandTypePeaking
)

type Band struct {
	Type      BandType
	Frequency float64
	Q         float64
	Gain      float64
}

type Preset struct {
	Bands []Band
}
