package main

import (
	"./perid/world"
)

func ExamplePrint() {
	lines := read_fileline("./case/print.pld")
	world.LineEval(lines)
	//Output:
	//Hello, Perid.
	//And, I'm born here!!
}
