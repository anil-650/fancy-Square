package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestSquare(t *testing.T) {
	t.Run("2 square 4", func(t *testing.T) {
		if Square(2) != 4 {
			t.Error("Failed with 2", 2)
		}
	})

	t.Run("15 square 225", func(t *testing.T) {
		if Square(15) != 225 {
			t.Error("Failed with 15", 15)
		}
	})

	t.Cleanup(func() {
		fmt.Println("Cleanup cache...")
		cleanning := exec.Command("go", "clean", "-cache")
		err := cleanning.Run()
		if err != nil {
			fmt.Println(err)
		}
	})
}
