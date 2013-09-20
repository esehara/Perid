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

    #Mining Token
    main := |*
        atom => {
            material := Ore{
                Token: ATOM,
                Value: data[ts:te],
            }
            material.Research()
            ores = append(ores, material) 
        };
        string => {
            material := Ore{
                Token: STRING,
                Value: data[ts:te],
            }
            material.Research()
            ores = append(ores, material)
		};
		integer => {
			material := Ore {
				Token: INTEGER,
				Value: data[ts:te],
			}
			material.Research()
			ores = append(ores, material)
		};
		operator => {
			material := Ore {
				Token: OPERATOR,
				Value: data[ts:te],
			}
			material.Research()
			ores = append(ores, material)
		};
		binder => {
			material := Ore {
				Token: BINDER,
				Value: data[ts:te],
			}
			material.Research()
			ores = append(ores, material)
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
	)

var tokens = [...]string {
    STRING: "STRING",
    ATOM: "ATOM",
	INTEGER: "INTEGER",
	OPERATOR: "OPERATOR",
	BINDER: "BINDER",
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
