package miner

import (
    "testing"
    "fmt"
)

var ores []Ore

func MiningResult(title string, ores []Ore) {
    fmt.Println("=====")
    fmt.Println(title)
    fmt.Println("-----")
    for index, ore := range ores {
        fmt.Printf("No. %d. \n", index)
        fmt.Printf("Token: %s\n", ore.Name)
        fmt.Println("Value:", ore.Value)
    }
    fmt.Println("=====")
}

func TestCanMining(t *testing.T) {
    src := "foobar"
    ores = Mining(src)
    MiningResult(
        "TestCanMining",
        ores)
    if len(ores) != 1 {
        t.Errorf("Ores should mining 1")
    }
}

func TestTokenStringGet(t *testing.T) {
    token := STRING
    if token.ToString() != "STRING" {
        t.Errorf("Cannot get token string")
    }
}

func isDefine (src, testname string, wish_token int) bool {
	ores = Mining(src)
	MiningResult("[Test::Define]" + testname, ores)
	if len(ores) != wish_token {
		return false
	}
	return true
}

func TestDefineString(t *testing.T) {
	if !isDefine(
		"print 'Hello!'",
		"String", 2) {
			t.Errorf("Ores should mining 2, but really %d.", len(ores))
		}
}

func TestDefineInteger(t *testing.T) {
	if !isDefine(
		"print 30",
		"Integer", 2) {
			t.Errorf("Ores should mining 2, but really %d.", len(ores))
		}
}

func TestDefineOperator(t *testing.T) {
	if !isDefine(
		"1 + 99",
		"Operator", 3) {
			t.Errorf("Ores should mining 3, but really %d.", len(ores))
		}
}

func TestDefineBinder(t *testing.T) {
	if !isDefine(
		"foobar = 100",
		"Binder", 3) {
			t.Errorf("Ores should mining 3, but really %d.", len(ores))
		}
}
