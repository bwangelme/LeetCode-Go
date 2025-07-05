package l169

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
		want int
	}

	tests := []args{
		{
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
