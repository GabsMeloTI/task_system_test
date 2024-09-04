package models

import "errors"

type Priority string

const (
	PriorityDisaster Priority = "disaster"
	PriorityHigh     Priority = "high"
	PriorityAverage  Priority = "average"
	PriorityWarning  Priority = "warning"
)

func (p Priority) IsValid() bool {
	switch p {
	case PriorityDisaster, PriorityHigh, PriorityAverage, PriorityWarning:
		return true
	}
	return false
}

func (p Priority) Validate() error {
	if !p.IsValid() {
		return errors.New("invalid priority")
	}
	return nil
}
