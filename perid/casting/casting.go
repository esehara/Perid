package casting

import (
	"../miner"
)

func Caster(src string) Val {
	var ores []miner.Ore
	ores = miner.Mining(src)
	l := &lex{
		ores,
		Val{},
	}
	yyParse(l)
	return l.root
}
