package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelection(t *testing.T) {
	tests := map[string]struct {
		toSort []int
		want   []int
	}{
		"success": {
			toSort: []int{5, 2, 3},
			want:   []int{2, 3, 5},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			Selection(tt.toSort)
			assert.Equal(t, tt.toSort, tt.want)
		})
	}
}
