package golang

import (
	"math"
)

func angleClock(hour int, minutes int) float64 {
	anglePreHour := 30.0
	hourAngleList := [12]float64{
		0, 30, 60, 90,
		120, 150, 180, 210,
		240, 270, 300, 330,
	}

	minutesBase0 := float64(minutes) * 6
	hourBase0 :=
		hourAngleList[hour%12] + minutesBase0/360*anglePreHour

	return math.Min(
		math.Abs(hourBase0-minutesBase0),
		360.0-math.Abs(hourBase0-minutesBase0),
	)
}
