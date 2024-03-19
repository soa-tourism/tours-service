package model

type TransportationType string

const (
	Walking TransportationType = "walking"
	Driving TransportationType = "driving"
	Cycling TransportationType = "cycling"
)

type TourTime struct {
	TimeInSeconds  float64
	Distance       float64
	Transportation TransportationType
}

func NewTourTime(timeInSeconds, distance float64, transportation TransportationType) *TourTime {
	return &TourTime{
		TimeInSeconds:  timeInSeconds,
		Distance:       distance,
		Transportation: transportation,
	}
}
