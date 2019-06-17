package main

import (
	"bufio"
	"os"
	"testing"
)

var screen = bufio.NewWriter(os.Stdout)

func Test_screenClear(t *testing.T) {
	n, err := screenClear(screen)
	if err != nil || n != 4 {
		t.Error("screenClear exec error")
	}
	t.Logf("screenClear return n = %d", n)
}

func Test_screenWrite(t *testing.T) {
	n, err := screenWrite(screen, []byte("b"))
	if err != nil || n != 2 {
		t.Error("screenWrite exec error")
	}
	t.Logf("screenWrite return n = %d", n)
}

func Test_screenFlush(t *testing.T) {
	if err := screenFlush(screen); err != nil {
		t.Error("screenFlush exec error")
	}
}

func Test_cleanScreen_deCleanScreen(t *testing.T) {
	err1 := cleanScreen()
	err2 := deCleanScreen()
	if err1 != nil || err2 != nil {
		t.Log("cleanScreen and deCleanScreen exec fail")
	}
	t.Log("cleanScreen and deCleanScreen exec success")
}
