package slicer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestartSlice(t *testing.T) {
	data := []int64{4, 6, 7, 8, 10, 5}
	beginID := 7
	want := []int64{7, 8, 10, 5}
	got := RestartSliceInt(data, int64(beginID))

	assert.ElementsMatch(t, want, got)
}
