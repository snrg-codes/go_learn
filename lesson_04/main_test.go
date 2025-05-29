package main
import (
	"testing"
)


func TestArea(t *testing.T) {
	r := Rectangle{Length: 5, Width: 10}
	expectedArea := 50
	if r.Area() != expectedArea {
		t.Errorf("Expected area %d, got %d", expectedArea, r.Area())
	}
}
func TestPerimeter(t *testing.T) {	
	r := Rectangle{Length: 5, Width: 10}
	expectedPerimeter := 30
	if r.Perimeter() != expectedPerimeter {
		t.Errorf("Expected perimeter %d, got %d", expectedPerimeter, r.Perimeter())
	}
}
func TestScale(t *testing.T) {	
	r := Rectangle{Length: 5, Width: 10}
	r.Scale(2)
	expectedLength := 10
	expectedWidth := 20
	if r.Length != expectedLength || r.Width != expectedWidth {
		t.Errorf("Expected Length %d and Width %d, got Length %d and Width %d", expectedLength, expectedWidth, r.Length, r.Width)
	}
}
