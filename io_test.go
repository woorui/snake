package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_ReadDirection(t *testing.T) {
	bu := bytes.NewBuffer([]byte{'w', 'a', 's', 'd'})

	dch := ReadToDirection(bu)

	got := []Direction{}
	for d := range dch {
		got = append(got, d)
	}

	want := []Direction{Up, Left, Down, Right}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadToDirection() got: %v, want: %v", got, want)
	}
}
