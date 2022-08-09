package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_Frame(t *testing.T) {
	bu := bytes.NewBuffer([]byte{})

	frame := NewFramer(bu, 4, 4)

	c, _ := NewCoordinate(1, 1, 1, 1, 'a', []Coordinate{})

	frame.RenderWithCoordinate(c)

	frame.Clear()

	want := []byte{45, 45, 45, 45, 10, 124, 97, 32, 124, 10, 124, 32, 32, 124, 10, 45, 45, 45, 45, 10, 27, 91, 50, 74}

	if got := bu.Bytes(); string(got) != string(want) {
		t.Errorf("test frame got = %v, want %v", got, want)
	}
}

func Test_genCantorMap(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want map[float64]int
	}{
		{
			name: "genCantorMap",
			args: args{
				width:  2,
				height: 2,
			},
			want: map[float64]int{0: 0, 1: 1, 1.5: 3, 3: 2, 3.5: 4, 6.5: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genCantorMap(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genCantorMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
