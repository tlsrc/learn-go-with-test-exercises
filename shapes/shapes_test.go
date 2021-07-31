package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	assertEquals(got, want, t)
}

func assertEquals(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("Wanted %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	type AreaTest struct {
		name         string
		shape        Shape
		expectedArea float64
	}

	areaTests := []AreaTest{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expectedArea: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actualArea := tt.shape.Area()
			if actualArea != tt.expectedArea {
				t.Errorf("Given %#v, area should be %g but got %g", tt.shape, tt.expectedArea, actualArea)
			}
		})
	}

}
