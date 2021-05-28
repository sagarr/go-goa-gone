package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("test array sum", func (t *testing.T)  {
		nums := [...]int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		if want != got {
			t.Errorf("got %d, want %d, input %v", got, want, nums)
		}	
	})

	t.Run("test slice", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := SumSlice(nums)
		want := 15

		if want != got {
			t.Errorf("got %d, want %d, input %v", got, want, nums)
		}	
	})
	
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{2, 3})
	want := []int{3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v got, %v want", got, want)
	}

}

func TestSumAllTails(t *testing.T) {

	checkSums := func (t testing.TB, got, want []int)  {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v got, %v want", got, want)
		}
	}
	t.Run("sum the slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3})
		want := []int{5, 5}
		checkSums(t, got, want)
	})
	
	t.Run("safely sum the empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}

func TestSlices(t *testing.T) {

	foo := func (x []int)  {
		x[0] = 0
	}

	a := []int{5, 5}
	fmt.Println(a)
	foo(a)
	fmt.Println(a)
}
