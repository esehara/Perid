package casting

import (
	"../miner"
	"fmt"
	"testing"
)

var ores []miner.Ore

func InitMining() []miner.Ore {
	src := "print 'test'"
	ores = miner.Mining(src)
	return ores
}

func TestTokenGet(t *testing.T) {
	ores = InitMining()
	if len(ores) != 2 {
		t.Errorf("Ores should mining 2.")
	}
}

func TestToToken(t *testing.T) {
	var ore miner.Ore
	var tokenize int

	ore = miner.Ore{Token: miner.STRING}
	tokenize = toToken(ore)
	fmt.Println(oretokens)
	if tokenize != STRING {
		t.Errorf("cannot tokenize: %d, should be STRING %d", tokenize, STRING)
	}
}
func generateCoin(src string) *lex{
	var l *lex
	var ores []miner.Ore
	ores = miner.Mining(src)
	fmt.Println(ores)
	l = &lex{
		ores,
		Val{},
	}
	yyParse(l)
	fmt.Println(l)
	return l
}

func TestMiningToCasting(t *testing.T) {
	generateCoin("print 'Hello'")
}

func TestAtomOnly(t *testing.T) {
	fmt.Println("=== atom only")
	generateCoin("print")
}

func TestCastingToExpose(t *testing.T) {
	fmt.Println("=== can refer function test")
	l := Caster("print 'Hello'")
	fmt.Println(l)
}

func TestOperator(t *testing.T) {
	fmt.Println("=== operation test")
	generateCoin("print 'Hello, ' + 'And, Im born here.'")
}

func TestDefine(t *testing.T) {
	fmt.Println("=== bind atom")
	generateCoin("foobar = 'Hello, Perid'")
}

func TestAtomAndAtom(t *testing.T) {
	fmt.Println("=== atom and atom")
	generateCoin("print foobar")
}
