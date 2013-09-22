/*
Enviroment in Perid.

Enviroment remember define atom and binding, it's singleton.
And if access enviroment, use AccessEnv function.
*/
package world

type env interface{}
type envs []env
var environment envs

// Access enviroment singleton
func AccessEnv() *envs {
	return &environment
}

// Access enviroment return array pointer,
// and array pointer not use len buildin function.
// then, "not pointer" valiable use.
func (environ *envs) Len() int {
	return len(environment)
}

// show Len() function.
// this access to "envs not pointer" from "envs pointer"
// and append any data(bind atom, etc).
func (environ *envs) Append(value env) {
	environment = append(environment, value)
}
