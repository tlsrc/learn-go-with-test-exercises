package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	input := []int{1, 2, 3}
	got := Sum(input)
	expected := 6

	if got != expected {
		t.Errorf("Expected %d but got %d given %v", expected, got, input)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("Sum tails of regular slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 5})
		want := []int{3, 5}

		assertSliceEquals(got, want, t)
	})
}

func assertSliceEquals(got []int, want []int, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %d but got %d", want, got)
	}
}

func TestSumTails(t *testing.T) {
	t.Run("with regular slices", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{2, 5, 6})
		want := []int{2, 11}

		assertSliceEquals(got, want, t)

	})

	t.Run("with an empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{1, 2})
		want := []int{0, 2}

		assertSliceEquals(got, want, t)
	})

	t.Run("trying out slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := make([]int, 10, 20)
		c := []int{}

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)

		copy(b, a)
		c = append(c, a[:1]...)

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)

	})
}
