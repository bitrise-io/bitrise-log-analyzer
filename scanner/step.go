package scanner

import "time"

type Color string

const (
	Red   Color = "\x1b[31;1m"
	Green Color = "\x1b[32;1m"
)

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
