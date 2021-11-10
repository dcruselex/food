package router

import (
	"testing"

	"github.com/dcruselex/food/internal/rapidapi"
)

var r1 rapidapi.Recipe = rapidapi.Recipe{
	Title:          "t1",
	ReadyInMinutes: 30,
}
var r2 rapidapi.Recipe = rapidapi.Recipe{
	Title:          "t2",
	ReadyInMinutes: 31,
}
var r3 rapidapi.Recipe = rapidapi.Recipe{
	Title:          "t3",
	ReadyInMinutes: 61,
}

func TestOption0(t *testing.T) {
	foods = []rapidapi.Recipe{r1, r2, r3}

	frs := getFilteredFoods(0)
	if len(frs) != 1 {
		t.Fatalf("expected length 1 received length %v", len(frs))
	}
	if frs[0].Title != "t1" {
		t.Fatalf("expected title t1 received title %v", frs[0].Title)
	}
}

func TestOption1(t *testing.T) {
	foods = []rapidapi.Recipe{r1, r2, r3}

	frs := getFilteredFoods(1)
	if len(frs) != 1 {
		t.Fatalf("expected length 1 received length %v", len(frs))
	}
	if frs[0].Title != "t2" {
		t.Fatalf("expected title t2 received title %v", frs[0].Title)
	}
}

func TestOption2(t *testing.T) {
	foods = []rapidapi.Recipe{r1, r2, r3}

	frs := getFilteredFoods(2)
	if len(frs) != 1 {
		t.Fatalf("expected length 1 received length %v", len(frs))
	}
	if frs[0].Title != "t3" {
		t.Fatalf("expected title t3 received title %v", frs[0].Title)
	}
}
