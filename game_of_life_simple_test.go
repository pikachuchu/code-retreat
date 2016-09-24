package game_of_life_simple

import (
	"testing"
)

var stableGrid = [][]int{
	{0, 0, 0, 0},
	{0, 1, 1, 0},
	{0, 1, 1, 0},
	{0, 0, 0, 0},
}
var flippingGrid0 = [][]int{
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 1, 1, 1, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
}
var flippingGrid1 = [][]int{
	{0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
}

func TestNextStatus(t *testing.T) {
	expected := []int{1, 0, 1, 0}
	actual := []int{
		nextStatus(1, 2),
		nextStatus(0, 2),
		nextStatus(0, 3),
		nextStatus(1, 4),
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected %v, found %v\n", expected[i], actual[i])
		}
	}
}

func TestCountNeighbors(t *testing.T) {
	expected := []int{1, 3, 2}
	actual := []int{
		countNeighbors(stableGrid, Coord{0, 0}),
		countNeighbors(stableGrid, Coord{1, 1}),
		countNeighbors(stableGrid, Coord{3, 1}),
	}
	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("expected %v, found %v\n", expected[i], actual[i])
		}
	}
}

func gridEquals(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}
	for i := range a {
		for j := range b {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestNextGen(t *testing.T) {
	expected := [][][]int{
		stableGrid,
		flippingGrid1,
		flippingGrid0,
	}
	actual := [][][]int{
		nextGen(stableGrid),
		nextGen(flippingGrid0),
		nextGen(flippingGrid1),
	}
	for i := range expected {
		if !gridEquals(expected[i], actual[i]) {
			t.Errorf("expected %v, found %v\n", expected[i], actual[i])
		}
	}
}

func TestRunGame(t *testing.T) {
	expected := [][][]int{
		stableGrid,
		flippingGrid1,
		flippingGrid0,
		flippingGrid0,
	}
	actual := [][][]int{
		runGame(stableGrid, 1000),
		runGame(flippingGrid0, 55),
		runGame(flippingGrid1, 7),
		runGame(flippingGrid0, 64),
	}
	for i := range expected {
		if !gridEquals(expected[i], actual[i]) {
			t.Errorf("expected %v, found %v\n", expected[i], actual[i])
		}
	}
}
