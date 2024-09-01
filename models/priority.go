package models

import "errors"

type Priority string

// Constants for Priority
const (
	PriorityDisaster Priority = "disaster"
	PriorityHigh     Priority = "high"
	PriorityAverage  Priority = "average"
	PriorityWarning  Priority = "warning"
)

// IsValid checks if the priority is valid
func (p Priority) IsValid() bool {
	switch p {
	case PriorityDisaster, PriorityHigh, PriorityAverage, PriorityWarning:
		return true
	}
	return false
}

// Validate returns an error if the priority is invalid
func (p Priority) Validate() error {
	if !p.IsValid() {
		return errors.New("invalid priority")
	}
	return nil
}
