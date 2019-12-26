package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuick(t *testing.T) {
	tests := map[string]struct {
		toSort []int
		want   []int
	}{
		"success": {
			toSort: []int{3, 1, 4, 2},
			want:   []int{1, 2, 3, 4},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			Quick(tt.toSort)
			assert.ElementsMatch(t, tt.toSort, tt.want)
		})
	}
}
