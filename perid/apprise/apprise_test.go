package apprise

import (
	"fmt"
	"testing"
	"../casting"
	"../world"
)

func ExamplePrint() {
	var coins casting.Val
	coins = casting.Caster("print 'Hello'")
	Apprise(coins)
	//output: Hello
}

func ExampleBindAtom() {
	env := world.AccessEnv()
	env.UnsafeReset()
	var coins casting.Val
	coins = casting.Caster("foobar = 'Hello'")
	Apprise(coins)
	coins = casting.Caster("print foobar")
	Apprise(coins)
	//output: Hello
}

func TestBindAtomAtom(t *testing.T) {
	env := world.AccessEnv()
	env.UnsafeReset()
	var coins casting.Val
	coins = casting.Caster("foobar = print 'Hello'")
	Apprise(coins)
	fmt.Println("=== Result Enviroment binding Atom to Atom")
	fmt.Println(env)
	coins = casting.Caster("foobar")
	Apprise(coins)
}

func TestBindAtom(t *testing.T) {
	env := world.AccessEnv()
	env.UnsafeReset()
	var coins casting.Val
	coins = casting.Caster("foobar = 'hoge'")
	Apprise(coins)
	if env.Len() != 1 {
		t.Errorf("Enviroment saves binding %s", env.Len())
	}
	fmt.Println("=== Result Enviroment binding")
	fmt.Println(env)
}
