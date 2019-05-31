package main

import (
	"bufio"
	"os"
	"testing"
)

var screen = bufio.NewWriter(os.Stdout)

func Test_ScreenClear(t *testing.T) {
	n, err := ScreenClear(screen)
	if err != nil || n != 4 {
		t.Error("ScreenClear exec error")
	}
	t.Logf("ScreenClear return n = %d", n)
}

func Test_ScreenWrite(t *testing.T) {
	n, err := ScreenWrite(screen, []byte("b"))
	if err != nil || n != 2 {
		t.Error("ScreenWrite exec error")
	}
	t.Logf("ScreenWrite return n = %d", n)
}

func Test_ScreenFlush(t *testing.T) {
	if err := ScreenFlush(screen); err != nil {
		t.Error("ScreenFlush exec error")
	}
}

func Test_cleanScreen_deCleanScreen(t *testing.T) {
	cleanScreen()
	deCleanScreen()
	t.Log("cleanScreen and deCleanScreen exec success")
}
