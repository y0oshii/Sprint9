package main

import (
	"testing"
)

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	result := generateRandomElements(0)

	if len(result) != 0 {
		t.Errorf("expected empty slice, got %d elements", len(result))
	}
}

func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	size := 10
	result := generateRandomElements(size)

	if len(result) != size {
		t.Errorf("expected slice size %d, got %d", size, len(result))
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("expected 0 for empty slice, got %d", result)
	}
}

func TestMaximumSingleElement(t *testing.T) {
	result := maximum([]int{7})
	if result != 7 {
		t.Errorf("expected 7, got %d", result)
	}
}

func TestMaximumMultipleElements(t *testing.T) {
	result := maximum([]int{1, 3, 9, 2, 5})
	if result != 9 {
		t.Errorf("expected 9, got %d", result)
	}
}

func TestChunkEmptySlice(t *testing.T) {
	result := maxChunks([]int{})
	if result != 0 {
		t.Errorf("expected 0 for empty slice, got %d", result)
	}
}

func TestSingleChunk(t *testing.T) {
	result := maxChunks([]int{32})
	if result != 32 {
		t.Errorf("expected 36, got %d", result)
	}
}

func TestMultipleChunks(t *testing.T) {
	result := maxChunks([]int{32, 1, 2, 3, 4, 33})
	if result != 33 {
		t.Errorf("expected 33, got %d", result)
	}
}

func TestChunkOdd(t *testing.T) {
	result := maxChunks([]int{12, 22, 3, 5, 1})
	if result != 22 {
		t.Errorf("expected 22, got %d", result)
	}
}
