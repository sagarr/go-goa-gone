package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{Widhth: 10.0, Height: 10.0}, 100.0},
		{Circle{Radius: 10.0}, 314.1592653589793},
		{Triangle{Height: 12, Base: 6}, 36},
	}
	
	for _, test := range areaTests {
		got := test.shape.Area()
		if test.want != got {
			t.Errorf("%#v, got %g, want %g", test, got, test.want)
		}
	}

}
