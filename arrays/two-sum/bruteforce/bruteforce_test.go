package bruteforce

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return [3, 4]",
			args: args{
				nums:   []int{1, 3, 7, 9, 2},
				target: 11,
			},
			want: []int{3, 4},
		},
		{
			name: "should return [0, 4]",
			args: args{
				nums:   []int{10, 8, 2, 9, 11},
				target: 21,
			},
			want: []int{0, 4},
		},
		{
			name: "should return nil",
			args: args{
				nums:   []int{10, 8, 2, 9, 11},
				target: 40,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
