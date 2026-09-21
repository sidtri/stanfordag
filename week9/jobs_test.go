package main

import "testing"

func TestWeightedCompletionTime(t *testing.T) {
	tests := []struct {
		name string
		jobs []Job
		want int
	}{
		{
			// catches: not ordering by decreasing (weight - length)
			name: "orders by decreasing difference",
			jobs: []Job{{weight: 3, length: 1}, {weight: 2, length: 2}},
			want: 9, // (3,1):3*1=3; (2,2):2*3=6
		},
		{
			// catches: breaking equal-difference ties by length instead of weight
			name: "breaks difference ties by higher weight first",
			jobs: []Job{{weight: 4, length: 2}, {weight: 5, length: 3}},
			want: 35, // (5,3):5*3=15; (4,2):4*5=20
		},
		{
			// catches: assuming differences are non-negative
			name: "handles negative differences",
			jobs: []Job{{weight: 1, length: 5}, {weight: 3, length: 2}},
			want: 13, // (3,2):3*2=6; (1,5):1*7=7
		},
		{
			// catches: using the ratio w/l instead of the difference
			name: "uses difference not ratio",
			jobs: []Job{{weight: 2, length: 1}, {weight: 5, length: 3}},
			want: 23, // (5,3):5*3=15; (2,1):2*4=8  (ratio order would give 22)
		},
		{
			// catches: default/empty return
			name: "single job",
			jobs: []Job{{weight: 4, length: 3}},
			want: 12, // 4*3
		},
		{
			name: "no jobs",
			jobs: []Job{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobs := tt.jobs
			if got := weightedCompletionTime(&jobs); got != tt.want {
				t.Errorf("weightedCompletionTime(%v) = %d, want %d", tt.jobs, got, tt.want)
			}
		})
	}
}
