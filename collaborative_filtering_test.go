package ml

import (
	"testing"
	"fmt"
)

func TestCollFilteringCostFunc(t *testing.T) {
	fmt.Println("Testing Collavorative Fitlers Cost Function...")

	cf := &CollaborativeFilter {
		Ratings: [][]float64{
			[]float64{5, 4, 0, 0},
			[]float64{3, 0, 0, 0},
			[]float64{4, 0, 0, 0},
			[]float64{3, 0, 0, 0},
			[]float64{3, 0, 0, 0},
		},
		AvailableRatings: [][]float64 {
			[]float64{1, 1, 0, 0},
			[]float64{1, 0, 0, 0},
			[]float64{1, 0, 0, 0},
			[]float64{1, 0, 0, 0},
			[]float64{1, 0, 0, 0},
		},
		ItemsTheta: [][]float64{
			[]float64{1.048686, -0.400232, 1.194119},
			[]float64{0.780851, -0.385626, 0.521198},
			[]float64{0.641509, -0.547854, -0.083796},
			[]float64{0.453618, -0.800218, 0.680481},
			[]float64{0.937538, 0.106090, 0.361953},
		},
		Theta: [][]float64{
			[]float64{0.28544, -1.68427, 0.26294},
			[]float64{0.50501, -0.45465, 0.31746},
			[]float64{-0.43192, -0.47880, 0.84671},
			[]float64{0.72860, -0.27189, 0.32684},
		},
	}

	j, _, err := cf.CostFunction(0.0, false)
	if err != nil {
		panic(err)
	}

	if j != 22.224626092180294 {
		t.Error("Expected cost: 22.224626092180294 , but", j, "returned")
	}
}

func TestCollFiltering(t *testing.T) {
	fmt.Println("Testing Collavorative Fitlers...")
	cf, err := NewCollFilterFromCsv(
		"test_data/collaborative_filtering/votes.csv",
		"test_data/collaborative_filtering/available_votes.csv",
		"test_data/collaborative_filtering/x.csv",
		"test_data/collaborative_filtering/theta.csv",
	)

	if err != nil {
		t.Error("Error loading the info from test files:", err)
	}

	//cf.InitializeThetas(10)
	//cf.Ratings = cf.Normalize()

	Fmincg(cf, 10, 100, true)
}
