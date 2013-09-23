// -*- mode:go-mode -*-
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
%token ATOM STRING OPERATOR BINDER DEFINE
%type <val> ATOM, STRING, OPERATOR, BINDER, DEFINE
%type <val> action, AtomExpression
%type <values> StringExpression
%%

   action:
   ATOM {
	 var v Value
	 v = Box{Type: ATOM, Value: $1,}
	 yylex.(*lex).NewBox(v)  
   }
|  ATOM StringExpression {
	 var v Value
	 v = Box{Type: ATOM, Value: $1, Box: $2}
	 yylex.(*lex).NewBox(v)
  }
| ATOM BINDER StringExpression {
  var v Value
  v = Box{Type: BINDER, Value: $1, Box: $3}
  yylex.(*lex).NewBox(v)
 }
| ATOM AtomExpression {
  var v Value
	if vals, ok := $2.(Val); ok {
	v = Box{Type: ATOM, Value: $1, Box: vals}
	yylex.(*lex).NewBox(v)
   }
 }
;

AtomExpression: ATOM {
  var newval Val
  newval = Val{Box{Type:ATOM, Value:$1}}
  $$ = newval
	};
StringExpression: STRING {
            var newval Val
            newval = Val{Box{Type: STRING, Value: $1}}
            $$ = newval
         }
       | STRING OPERATOR STRING {
            var newval Val
            result := operate_string($1, $2, $3)
            newval = Val{Box{Type: STRING, Value: result}}
            $$ = newval
         }
       ;
%% /* --- start of programs */

var oretokens = [...]int {
    miner.ATOM: ATOM,
    miner.STRING: STRING,
    miner.OPERATOR: OPERATOR,
	miner.BINDER: BINDER,
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
