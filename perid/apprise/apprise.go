package apprise
import (
	"fmt"
	"../casting"
)

func Apprise(toplevel casting.Val) {
	eval(toplevel)
}

func eval(node casting.Val) casting.Value{
	var line casting.Value
	line = node[0]
	return evalroot(line)
}

func evalroot(line casting.Value) casting.Value{
	if box, ok := line.(casting.Box); ok {
		return evalbox(box)
	}
	var result casting.Value
	return result
}

func evalbox(box casting.Box) casting.Value {
	switch box.Type {
	case casting.ATOM:
		evalatom(box)
	case casting.STRING:
		return box.Value 
	}
	var result casting.Value
	return result
}

func evalatom(box casting.Box) {
	var result casting.Value
	if len(box.Box) > 0 {
		result = eval(box.Box)
	}

	if box.Value == "print" {
		if str,ok := result.(string); ok {
			result = str[1:len(str) - 1]
		}
		fmt.Println(result)
	}
}
