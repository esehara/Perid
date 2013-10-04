// -*- mode:go-mode -*-
%%{
    #
    # Define Analizer Machine for Perid
    #
    machine miner;
    
    # Parse Rule
    
    # ---- Literal ----
    string = ("'" . (any - "'")* . "'")
            |('"' . (any - '"')* . '"');
	integer = digit+;
    atom = (alpha (alpha | digit)* );

	# ---- Operator ----
	operator = "+" | "-" | "*" | "/";
	binder = "=";
    
    # ---- parins ----
    roundparen_o = "(";
    roundparen_c = ")";

    #Mining Token
    main := |*
        atom => {
            ores = appendOre(ATOM, ores, data, ts, te)
        };
        string => {
            ores = appendOre(STRING, ores, data, ts, te)
		};
		integer => {
			ores = appendOre(INTEGER, ores, data, ts, te)
		};
		operator => {
            ores = appendOre(OPERATOR, ores, data, ts, te)
		};
		binder => {
		    ores = appendOre(BINDER, ores, data, ts, te)
        };
        roundparen_o => {
            ores = appendOre(ROUNDPAREN_O, ores, data, ts, te)
        };
        roundparen_c => {
            ores = appendOre(ROUNDPAREN_C, ores, data, ts, te)
        };
        space;
    *|;
}%%

package miner

%% write data;

type Token int


const(
    //Special Token
    ILLEGAL Token = iota
    EOF
    COMMMENT

    begin_define_literal
        ATOM
        STRING
		INTEGER
    end_define_literal
	OPERATOR
	BINDER
	ROUNDPAREN_O
    ROUNDPAREN_C
    )

var tokens = [...]string {
    STRING: "STRING",
    ATOM: "ATOM",
	INTEGER: "INTEGER",
	OPERATOR: "OPERATOR",
	BINDER: "BINDER",
    ROUNDPAREN_O: "(",
    ROUNDPAREN_C: ")",
}

var keywords map[string]Token
type Any interface {}

/*
 * "Ore" called "Token" by other language (Lexical Parser).
 */
type Ore struct {
    Token Token
    Name string
    Value Any
}
var material Ore

func appendOre (miningore Token, bag []Ore, data string, ts, te int) []Ore {
	material := Ore {
	    Token: miningore,
		Value: data[ts:te],
	}
	material.Research()
	return append(bag, material)
}

func (tok Token) ToString() string {
    return tokens[tok]
}

func (ore *Ore) Research() {
    ore.Name = ore.Token.ToString()
}

/*
 * Perid Language, "Tokenize" called "Mining".
 */
func Mining (data string) []Ore {
    cs, p, pe := 0, 0, len(data)
    eof := len(data)
    var ts, te, act int
    var ores []Ore

    %% write init;
    %% write exec;

    _, _, _ = ts, te, act
    return ores
}
// vim: set syntax=go:
