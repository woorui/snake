# Game snake written in go, No third-dependence

[![Build Status](https://cloud.drone.io/api/badges/woorui/snake/status.svg)](https://cloud.drone.io/woorui/snake)

The snake is controlled with `w`, `a`, `s` and `d`, It don't support windows (but docker support)

## Running in the project directory without building, just for testing

```bash
go run -race $(ls -1 *.go | grep -v _test.go)
```

## Install

```bash
go get -u github.com/woorui/snake
```

## Flag

```bash
snake --help
  -debug
        debug mode.
  -height int
        Game stage height. (default 12)
  -width int
        Game stage width. (default 25)
```

## Run

```bash
snake
```

![Show the running result](snake_run.gif)

## Run with docker

see here: https://github.com/woorui/snake/pkgs/container/snake
