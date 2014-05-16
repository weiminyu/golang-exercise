package mathgen

import (
	"fmt"
	"io"
	"math/rand"
)

func Generate(out io.Writer, count, bound int) {
	fmt.Fprintf(out, "<!DOCTYPE html><html><head><style>p.ex1 {margin-top:0.1cm; }")
	fmt.Fprintf(out, "p.small {line-height:70%;}</style></head><body>")
	fmt.Fprintln(out, "<pre><b><font size=\"5\"><p class=\"ex1\"><p class=\"small\">")

	for _, p := range makeAll(count, bound) {
		fmt.Fprintf(out, "%2d %v %2d =   <br>%s", p.left, p.op, p.right, "\n")
	}
	fmt.Fprintln(out, "</p></p></font></b></pre></body></html>")
}

func makeAll(count, bound int) map[Problem]*Problem {
	rand := rand.New(rand.NewSource(99))
	problems := make(map[Problem]*Problem)

	for len(problems) < count {
		ptr := makeOne(rand, bound)
		problems[*ptr] = ptr
	}
	return problems
}

func makeOne(rand *rand.Rand, bound int) *Problem {
	op_val := rand.Float32()
	f := int(rand.Float32()*float32(bound)) + 1
	s := int(rand.Float32()*float32(bound)) + 1

	var op Operator

	if op_val < float32(0.5) {
		op = ADD
	} else {
		op = SUBTRACT
	}

	if f == s && op == SUBTRACT {
		op = ADD
	}

	if op == ADD && f+s > bound && s > 1 {
		s = s >> 1
	}
	if op == SUBTRACT && f < s {
		f = f ^ s
		s = f ^ s
		f = f ^ s
	}

	return &Problem{op, f, s}
}

type Operator int

const (
	ADD Operator = iota // 1
	SUBTRACT
)

func (o Operator) String() string {
	switch o {
	case ADD:
		return "+"
	case SUBTRACT:
		return "-"
	default:
		return "" // TODO: error
	}
}

type Problem struct {
	op          Operator
	left, right int
}
