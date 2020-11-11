package bruteforce

import "testing"

func TestGetTotalWater(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 8",
			args: args{
				heights: []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTotalWater(tt.args.heights); got != tt.want {
				t.Errorf("GetTotalWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			name: "should return 10",
			args: args{nums: []int{10, 5}},
			wantMax: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := Max(tt.args.nums...); gotMax != tt.wantMax {
				t.Errorf("Max() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
	}{
		{
			name: "should return 5",
			args: args{nums: []int{10, 5}},
			wantMin: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := Min(tt.args.nums...); gotMin != tt.wantMin {
				t.Errorf("Min() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}