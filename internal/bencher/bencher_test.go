package bencher

import "testing"

func TestStdDeviation(t *testing.T) {
	b := New("Bubble Sort", 50)

	tests := []struct {
		name     string
		nums     []int64
		expected int64
	}{
		{
			name:     "test 1",
			nums:     []int64{982803, 3287510, 32891274, 29875985, 2910934},
			expected: 14255517,
		},
		{
			name:     "test 2",
			nums:     []int64{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "test 3",
			nums:     []int64{1, 100, 1000, 10000, 100000},
			expected: 39069,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avg := int64(0)
			for _, num := range tt.nums {
				avg += num
			}
			avg /= int64(len(tt.nums))
			t.Logf("avg: %d", avg)

			got := b.getStandardDeviation(tt.nums, avg)
			if got != tt.expected {
				t.Errorf("b.getStandardDeviation() = %v, want %v", got, tt.expected)
			}
		})
	}
}
