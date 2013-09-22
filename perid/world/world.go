package world

import (
	"../casting"
	"../apprise"
)

type env interface{}
type envs []env
var environment envs

/*
Enviroment in Perid.

Enviroment remember define atom and binding, it's singleton.
And if access enviroment, use AccessEnv function.
*/

func AccessEnv() *envs {
	return &environment
}

func (environ *envs) Len() int {
	return len(environment)
}

func (environ *envs) Append(value env) {
	environment = append(environment, value)
}

func LineEval(lines []string) {
	for _, line := range lines {
		var coins casting.Val
		coins = casting.Caster(line)
		apprise.Apprise(coins)
	}
}
