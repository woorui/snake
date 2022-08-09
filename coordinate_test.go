package main

import (
	"reflect"
	"testing"
)

func TestNewCoordinate(t *testing.T) {
	type args struct {
		minX   int
		maxX   int
		minY   int
		maxY   int
		ink    byte
		forbid []Coordinate
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "new_coordinate",
			args: args{
				minX: 0,
				maxX: 2,
				minY: 0,
				maxY: 2,
				ink:  0,
				forbid: []Coordinate{
					{ink: 0, x: 1, y: 1},
					{ink: 0, x: 3, y: 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(b *testing.T) {
			for i := 0; i < 100; i++ {
				got, err := NewCoordinate(tt.args.minX, tt.args.maxX, tt.args.minY, tt.args.maxY, tt.args.ink, tt.args.forbid)
				if err != nil {
					t.Errorf("NewCoordinate() error = %v, wantErr %v", err, ErrTooManyRandTimes)
					return
				}
				for _, item := range tt.args.forbid {
					if got.x == item.x && got.y == item.y {
						t.Errorf("NewCoordinate() failed, got = %v, don't want %v", got, tt.args.forbid)
					}
				}
			}
		})
	}
}

func TestRemoveFront(t *testing.T) {
	type args struct {
		list []Coordinate
	}
	tests := []struct {
		name string
		args args
		want []Coordinate
	}{
		{
			name: "normal",
			args: args{
				list: []Coordinate{{ink: 0, x: 2, y: 2}, {ink: 0, x: 1, y: 1}},
			},
			want: []Coordinate{
				{ink: 0, x: 1, y: 1},
			},
		},
		{
			name: "empty",
			args: args{
				list: []Coordinate{},
			},
			want: []Coordinate{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFront(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type args struct {
		list   []Coordinate
		others []Coordinate
	}
	tests := []struct {
		name string
		args args
		want []Coordinate
	}{
		{
			name: "normal",
			args: args{
				list:   []Coordinate{{ink: 0, x: 1, y: 1}},
				others: []Coordinate{{ink: 0, x: 2, y: 2}},
			},
			want: []Coordinate{{ink: 0, x: 1, y: 1}, {ink: 0, x: 2, y: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Append(tt.args.list, tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncludes(t *testing.T) {
	type args struct {
		list    []Coordinate
		element Coordinate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "includes",
			args: args{
				list:    []Coordinate{{ink: 0, x: 1, y: 1}, {ink: 0, x: 2, y: 2}},
				element: Coordinate{ink: 0, x: 1, y: 1},
			},
			want: true,
		},
		{
			name: "not_includes",
			args: args{
				list:    []Coordinate{{ink: 0, x: 1, y: 1}, {ink: 0, x: 2, y: 2}},
				element: Coordinate{ink: 0, x: 1, y: 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Includes(tt.args.list, tt.args.element); got != tt.want {
				t.Errorf("Includes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	type rang struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args rang
		want rang
	}{
		{
			name: "random",
			args: rang{min: 0, max: 1},
			want: rang{min: 0, max: 1},
		},
		{
			name: "min > max",
			args: rang{min: 2, max: 1},
			want: rang{min: -1, max: 1}, // equal 0
		},
		{
			name: "equals",
			args: rang{min: 2, max: 2},
			want: rang{min: 1, max: 3}, // equal 2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomInt(tt.args.min, tt.args.max); got < tt.want.min || got >= tt.want.max {
				t.Errorf("RandomInt() = %v, want range min = %v, max = %v", got, tt.args.min, got >= tt.args.max)
			}
		})
	}
}

func TestCoordinate_CoordinateList(t *testing.T) {
	type fields struct {
		ink byte
		x   int
		y   int
	}
	tests := []struct {
		name   string
		fields fields
		want   []Coordinate
	}{
		{
			name:   "normal",
			fields: fields{ink: 0, x: 1, y: 1},
			want:   []Coordinate{{ink: 0, x: 1, y: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Coordinate{
				ink: tt.fields.ink,
				x:   tt.fields.x,
				y:   tt.fields.y,
			}
			if got := c.CoordinateList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coordinate.CoordinateList() = %v, want %v", got, tt.want)
			}
		})
	}
}
