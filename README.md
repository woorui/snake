# Game snake written in go, No third-dependence

[![Build Status](https://cloud.drone.io/api/badges/woorui/snake/status.svg)](https://cloud.drone.io/woorui/snake)
[![Docker](https://github.com/woorui/snake/actions/workflows/docker-publish.yml/badge.svg?branch=main)](https://github.com/woorui/snake/actions/workflows/docker-publish.yml)
[![Docker Image CI](https://github.com/woorui/snake/actions/workflows/docker-image.yml/badge.svg?branch=main)](https://github.com/woorui/snake/actions/workflows/docker-image.yml)

The snake is controlled with `w`, `a`, `s` and `d`, It don't support windows (but docker supported).

## Install

```bash
go install github.com/woorui/snake
```

## Flags

```bash
> snake -h
Usage of snake:
  -height int
        game stage height (default 12)
  -speed int
        game speed, duration between two frames (default 120)
  -width int
        game stage width (default 25)
```

## Run

```bash
snake
```

![Show the running result](snake_run.gif)

## Run with docker

see here: https://github.com/woorui/snake/pkgs/container/snake
