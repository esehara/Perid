package casting

import (
    "../miner"
    "testing"
    "fmt"
)

var ores []miner.Ore

func InitMining() []miner.Ore{
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

/*
func TestUsingLex(t *testing.T) {

    l := &lex{
        []token{{ATOM, "print"}, {STRING, "'Hello'"}},
        Val{},
    }
    yyParse(l)
    fmt.Println(l)
}
*/

func TestToToken(t *testing.T) {
    var ore miner.Ore
    var tokenize int
    
    ore = miner.Ore{Token:miner.STRING}
    tokenize = toToken(ore)
    fmt.Println(oretokens)
    if tokenize != STRING {
        t.Errorf("cannot tokenize: %d, should be STRING %d", tokenize, STRING)
    }
}

func TestMiningToCasting(t *testing.T) {
    var ores []miner.Ore
    ores = miner.Mining("print 'Hello'")
    fmt.Println(ores)
    l := &lex {
        ores,
        Val{},
    }
    yyParse(l)
    fmt.Println(l)
}
