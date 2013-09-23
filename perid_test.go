package main

import (
	"./perid/apprise"
)

func ExamplePrint() {
	lines := read_fileline("./case/print.prd")
	apprise.LineEval(lines)
	//Output:
	//Hello, Perid.
	//And, I'm born here!!
}
