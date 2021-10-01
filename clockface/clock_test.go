package clockface

import (
	"math"
	"testing"
	"time"
)

type AngleTestCase struct {
	time          time.Time
	expectedAngle float64
}

func TestSecondsInRadians(t *testing.T) {
	cases := []AngleTestCase{
		{hoursMinsSecs(0, 0, 30), math.Pi},
		{hoursMinsSecs(0, 0, 0), 0},
		{hoursMinsSecs(0, 0, 45), math.Pi * 45 / 30},
		{hoursMinsSecs(0, 0, 7), math.Pi * 7 / 30},
	}
	runAll(cases, secondsInRadians, t)
}

func TestMinutesInRadians(t *testing.T) {
	cases := []AngleTestCase{
		{hoursMinsSecs(0, 30, 0), math.Pi},
		{hoursMinsSecs(0, 0, 7), 7 * math.Pi / (60 * 30)},
	}
	runAll(cases, minutesInRadians, t)
}

func TestHoursInRadians(t *testing.T) {
	cases := []AngleTestCase{
		{hoursMinsSecs(9, 0, 0), math.Pi * 1.5},
		{hoursMinsSecs(15, 0, 0), math.Pi * 0.5},
		{hoursMinsSecs(6, 30, 0), math.Pi + math.Pi/12},
	}
	runAll(cases, hoursInRadians, t)

}

func runAll(cases []AngleTestCase, funcUnderTest func(time.Time) float64, t *testing.T) {
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := funcUnderTest(c.time)
			if got != c.expectedAngle {
				t.Fatalf("Wanted %v radians, but got %v", c.expectedAngle, got)
			}
		})
	}
}

func testName(time time.Time) string {
	return time.Format("11:12:13")
}

func hoursMinsSecs(h, m, s int) time.Time {
	return time.Date(2000, time.January, 5, h, m, s, 0, time.UTC)
}
