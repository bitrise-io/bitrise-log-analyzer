package build

import "time"

type Status int

const (
	Success Status = iota
	Failed
	Skipped
)

type Step struct {
	ID      string
	Version string
	Status
	Lines    []string
	Duration time.Duration
}

type Log struct {
	steps []Step
}

type Color string

const (
	Red    Color = "\x1b[31;1m"
	Green  Color = "\x1b[32;1m"
	Yellow Color = "\x1b[33;1m"
	Reset  Color = "\x1b[0m"
)
