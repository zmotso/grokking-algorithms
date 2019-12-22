package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinary(t *testing.T) {
	tests := map[string]struct {
		search  int
		in      []int
		want    int
		wantErr error
	}{
		"10": {
			search:  10,
			in:      []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			want:    8,
			wantErr: nil,
		},
		"3": {
			search:  3,
			in:      []int{2, 3, 4, 5},
			want:    1,
			wantErr: nil,
		},
		"not found": {
			search:  33,
			in:      []int{2, 3, 4, 5},
			want:    0,
			wantErr: ErrNotFount,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Binary(tt.search, tt.in)
			if assert.Equal(t, tt.wantErr, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
