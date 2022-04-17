package utils

import "testing"

type Transformable struct {
	first  string
	second string
}

type Result struct {
	third  string
	fourth string
}

func TestMap(t *testing.T) {
	var initial []Transformable
	initial = append(initial,
		Transformable{
			first:  "first_0",
			second: "second_0",
		},
		Transformable{
			first:  "first_1",
			second: "second_2",
		},
	)
	var mapped = Map(initial, func(t Transformable) Result {
		return Result{
			third:  t.first,
			fourth: t.second,
		}
	})
	// Map should result in same-sized slice
	actualLen := len(mapped)
	expectedLen := 2

	if actualLen != expectedLen {
		t.Errorf("got %q, wanted %q", actualLen, expectedLen)
	}

	// a linear 1:1 transform should happen
	actualTransform := (mapped[0].third == initial[0].first && mapped[0].fourth == initial[0].second) && (mapped[1].third == initial[1].first && mapped[1].fourth == initial[1].second)
	if !actualTransform {
		t.Errorf("transform was not linear")
	}
}
