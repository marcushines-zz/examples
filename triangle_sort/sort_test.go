package sort

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		desc string
		in   []int
		out  Triangle
		min  int
		max  int
		sort []int
	}{{
		desc: "nil",
	}, {
		desc: "empty",
		in:   []int{},
		out:  Triangle{},
	}, {
		desc: "single",
		in:   []int{1},
		out:  Triangle{1},
		min:  1,
		max:  1,
		sort: []int{1},
	}, {
		desc: "valid",
		in:   []int{1, 5, 6, 8, 10, 7, 2},
		out:  Triangle{1, 5, 6, 8, 10, 7, 2},
		min:  1,
		max:  10,
		sort: []int{1, 2, 5, 6, 7, 8, 10},
	}, {
		desc: "valid 3",
		in:   []int{1, 2, 3},
		out:  Triangle{1, 2, 3},
		min:  1,
		max:  3,
		sort: []int{1, 2, 3},
	}, {
		desc: "invalid",
		in:   []int{1, 2, 4, 2, 3},
		out:  nil,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			l := New(tt.in)
			if !reflect.DeepEqual(l, tt.out) {
				t.Fatalf("New(%v) failed: got %v, want %v", tt.in, l, tt.out)
			}
			if len(l) == 0 {
				return
			}
			if *l.Min() != tt.min {
				t.Fatalf("Min() failed: got %v, want %v", l.Min(), tt.min)
			}
			if *l.Max() != tt.max {
				t.Fatalf("Max() failed: got %v, want %v", l.Max(), tt.max)
			}
			if !reflect.DeepEqual(l.Sort(), tt.sort) {
				t.Fatalf("Sort() failed: got %v, want %v", l.Sort(), tt.sort)
			}
		})
	}
}
