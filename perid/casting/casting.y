%{
package casting

import (
    "fmt"
    "../miner"
)
%}
%union {
    tok int
    val interface{}
    values Val
    box Box
}

%token ATOM STRING
%type <val> ATOM, STRING, action
%type <values> literal
%%

action: ATOM {
                var v Value
                v = Box{Value: ATOM}
                yylex.(*lex).NewBox(v)
            }
      | ATOM literal {
                var v Value
                v = Box{Type: ATOM, Value: $1, Box: $2}
                yylex.(*lex).NewBox(v)
            }

      ;

literal: STRING {
            var newval Val
            newval = Val{Box{Type: STRING, Value: $1}}
            $$ = newval
            }
       ;
%% /* --- start of programs */

var oretokens = [...]int {
    miner.ATOM: ATOM,
    miner.STRING: STRING,
}

type Value interface{}
type Val []Value

type token struct {
    tok int
    val Value
}

type lex struct {
    tokens []miner.Ore
    root Val
}

func toToken(ore miner.Ore) int {
    return oretokens[ore.Token]
}

/* Value Structure
 *
 * root: [
 *           {Box{
 *              val: "print",
 *              box: [{STRING, "'any string'"}]
 *           }, 
 *       ]
 */
type Box struct {
    Type int
    Value Value
    Box Val
}

func (l *lex) NewBox(any Value) {
    l.root = append(l.root, any)
}

func (l *lex) Error (s string) {
}

func (l *lex) Lex(lval *yySymType) int {
    if len(l.tokens) == 0 {
        return 0
    }
    v := l.tokens[0]
    l.tokens = l.tokens[1:]
    lval.val = v.Value
    return toToken(v)
}

// vim: set syntax=go:
