package main_test

import (
	"gounited/lecture1/solution"
	"testing"

	"github.com/kyokomi/emoji"
)

func TestSolution(t *testing.T) {
	if solution.GetMessage() != emoji.Sprint("Hello :world_map:!") {
		t.Fatal("Wrong solution!")
	}
}
