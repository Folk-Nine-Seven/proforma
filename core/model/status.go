package model

type Status int8

const (
	None Status = iota
	Open
	Blocked
	InProgress
	Closed
	Completed
)
