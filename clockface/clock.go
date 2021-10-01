package clockface

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

const clockCenterX = 150
const clockCenterY = 150
const secondHandLength = 90
const minutesHandLength = 80
const hoursHandLength = 50

func secondHand(t time.Time) Point {
	return makeEndPoint(secondsInRadians(t), secondHandLength)
}

func minutesHand(t time.Time) Point {
	return makeEndPoint(minutesInRadians(t), minutesHandLength)
}

func hoursHand(t time.Time) Point {
	return makeEndPoint(hoursInRadians(t), hoursHandLength)
}

func makeEndPoint(angle float64, length float64) Point {
	return Point{
		clockCenterX + length*math.Sin(angle),
		clockCenterY - length*math.Cos(angle),
	}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return secondsInRadians(t)/60 + math.Pi/(30/float64(t.Minute()))
}

func hoursInRadians(t time.Time) float64 {
	return minutesInRadians(t)/12 + math.Pi/(6/float64(t.Hour()%12))
}
