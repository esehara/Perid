package world

import (
	"../casting"
	"../apprise"
)

func LineEval(lines []string) {
	for _, line := range lines {
		var coins casting.Val
		coins = casting.Caster(line)
		apprise.Apprise(coins)
	}
}
