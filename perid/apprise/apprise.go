package apprise
import (
	"fmt"
	"../casting"
	"../goldkey"
	"../world"
)

func Apprise(toplevel casting.Val) {
	eval(toplevel)
}

func eval(node casting.Val) casting.Value{
	var line casting.Value
	if len(node) > 0 {
		line = node[0]
		return evalroot(line)
	}
	return new(casting.Value)
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
		return evalatom(box)
	case casting.STRING:
		return box.Value 
	case casting.BINDER:
		bindatom(box)
	}
	var result casting.Value
	return result
}

func evalatom(box casting.Box) casting.Value {
	var result casting.Value
	if len(box.Box) > 0 {
		result = eval(box.Box)
	}

	if box.Value == "print" {
		if str,ok := result.(string); ok {
			result = goldkey.CastString(str)
		}
		fmt.Println(result)
		result = new(casting.Value)
	} else {
		if str, ok := box.Value.(string); ok {
			env := world.AccessEnv()
			if ok, atomvalue := env.Search(str); ok {
				if parseval, ok := atomvalue.(casting.Val); ok {
					result = eval(parseval)
				}
			}
		}
	}
	return result
}

func bindatom(box casting.Box) {
	var setvalue world.EnvValue
	if str, ok := box.Value.(string); ok {
		setvalue = world.EnvValue{
			Name: str,
			Val: box.Box,
		}
	}
	env := world.AccessEnv()
	env.SetValue(setvalue)
}

func LineEval(lines []string) {
	for _, line := range lines {
		var coins casting.Val
		coins = casting.Caster(line)
		Apprise(coins)
	}
}
