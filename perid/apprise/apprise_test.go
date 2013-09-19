package apprise

import (
	"../casting"
)

func ExamplePrint() {
	var coins casting.Val
	coins = casting.Caster("print 'Hello'")
	Apprise(coins)
	//output: Hello
}
