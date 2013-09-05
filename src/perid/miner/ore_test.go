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
    if token.ToString() != "String" {
        t.Errorf("Cannot get token string")
    }
}

func TestDefineLiteral(t *testing.T) {
    src := "print 'foobar'"
    ores = Mining(src)
    MiningResult(
        "TestDefineLiteral",
        ores)
    if len(ores) != 2 {
        t.Errorf("Ores should mining 2, but really %d.", len(ores))
    }
}
