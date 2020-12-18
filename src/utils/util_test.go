package utils

import (
	"fmt"
	"testing"
)

func TestDoubleAverage(t *testing.T) {
	testDoubleAverage(10, 10000)
	type args struct {
		count  int64
		amount int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name:"1",
			args:args{10, 10000},
			want:10000,
		},
		{
			name:"2",
			args:args{10, 10000},
			want:10000,
		},
		{
			name:"3",
			args:args{10, 10000},
			want:10000,
		},
		{
			name:"4",
			args:args{10, 10000},
			want:10000,
		},
		{
			name:"1",
			args:args{1, 10000},
			want:10001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleAverage(tt.args.count, tt.args.amount); got > tt.want {
				t.Errorf("DoubleAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testDoubleAverage(count, amount int64) {
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := DoubleAverage(count - i, amount)
		fmt.Printf("用户%d获得了红包为%f元\n", i + 1, float64(x)/float64(100))
		sum += x
		amount -= x
	}
	fmt.Printf("合计发红包为%d分",sum)
}