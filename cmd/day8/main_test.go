package main

import "testing"

func TestDecodeImage(t *testing.T) {
	t.Run("Sample 1", func(t *testing.T) {
		imageData := "123456789012"
		width := 3
		height := 2
		expected := 1
		result := verifyImage(imageData, width, height)
		if result != expected {
			t.Errorf("Expected %d, but was %d", expected, result)
		}
	})
}

func TestBuildImage(t *testing.T) {
	buildImage("0222112222120000", 2,2)
}