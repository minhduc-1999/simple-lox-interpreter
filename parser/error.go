package parser

import (
	"fmt"
)

type RuntimeErr struct {
	// TODO: track error position
	msg string
}

func (e RuntimeErr) Error() string {
	return fmt.Sprintf("Runtime error: %v", e.msg)
}

func NewRuntimeErr(msg string) RuntimeErr {
	return RuntimeErr{
		msg: msg,
	}
}
