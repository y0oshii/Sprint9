package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	result := generateRandomElements(0)

	assert.Len(t, result, 0)
}

func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	size := 10
	result := generateRandomElements(size)

	assert.Len(t, result, size)
}

func TestMaximumEmptySlice(t *testing.T) {
	result := maximum([]int{})
	assert.Equal(t, 0, result)
}

func TestMaximumSingleElement(t *testing.T) {
	result := maximum([]int{7})
	assert.Equal(t, 7, result)
}

func TestMaximumMultipleElements(t *testing.T) {
	result := maximum([]int{1, 3, 9, 2, 5})
	assert.Equal(t, 9, result)
}

func TestChunkEmptySlice(t *testing.T) {
	result := maxChunks([]int{})
	assert.NotNil(t, result)
	assert.Equal(t, result, 0)
}

func TestSingleChunk(t *testing.T) {
	result := maxChunks([]int{32})
	assert.Equal(t, 32, result)
}

func TestMultipleChunks(t *testing.T) {
	result := maxChunks([]int{32, 1, 2, 3, 4, 33})
	assert.Equal(t, 33, result)
}

func TestChunkOdd(t *testing.T) {
	result := maxChunks([]int{12, 22, 3, 5, 1})
	assert.Equal(t, 22, result)
}
