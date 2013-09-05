%%{
    #
    # Define Analizer Machine for Perid
    #
    machine miner;
    
    # Parse Rule
    
    # keyword = "print";
    
    #Literal ---------------
    string = ("'" . (any - "'")* . "'")
            |('"' . (any - '"')* . '"');
    
    #literal = string;
    #Literal Define END
    
    atom = (alpha (alpha | digit)* );
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
            ores = append(ores, material) };
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
    end_define_literal
)

var tokens = [...]string {
    STRING: "String",
    ATOM: "Atom",
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
