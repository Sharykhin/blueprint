package trace

import (
	"fmt"
	"io"
)

type (
	// Tracer is the interface that describes an object capable of
	// tracing events throughout code.
	Tracer interface {
		Trace(...interface{})
	}

	tracer struct {
		out io.Writer
	}

	nilTracer struct{}
)

func (t tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *nilTracer) Trace(a ...interface{}) {}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &nilTracer{}
}
