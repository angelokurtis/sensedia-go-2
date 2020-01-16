package clima

type Medidor interface {
	AferirMinhaLocalização() (*Clima, error)
	AferirCoordenadas(lat float64, long float64) (*Clima, error)
}
