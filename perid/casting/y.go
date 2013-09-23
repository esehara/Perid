
//line casting.y:3
package casting

import (
    "fmt"
    "../miner"
)

//line casting.y:10
type yySymType struct {
	yys int
    tok int
    val interface{}
    values Val
    box Box
}

const ATOM = 57346
const STRING = 57347
const OPERATOR = 57348
const BINDER = 57349
const DEFINE = 57350

var yyToknames = []string{
	"ATOM",
	"STRING",
	"OPERATOR",
	"BINDER",
	"DEFINE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line casting.y:64
 /* --- start of programs */

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

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 8
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 12

var yyAct = []int{

	7, 6, 10, 4, 9, 6, 3, 2, 5, 1,
	0, 8,
}
var yyPact = []int{

	3, -1000, -4, -1000, 0, -1000, -2, -1000, -1000, -3,
	-1000,
}
var yyPgo = []int{

	0, 9, 8, 6,
}
var yyR1 = []int{

	0, 1, 1, 1, 1, 2, 3, 3,
}
var yyR2 = []int{

	0, 1, 2, 3, 2, 1, 1, 3,
}
var yyChk = []int{

	-1000, -1, 4, -3, 7, -2, 5, 4, -3, 6,
	5,
}
var yyDef = []int{

	0, -2, 1, 2, 0, 4, 6, 5, 3, 0,
	7,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c > 0 && c <= len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return fmt.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return fmt.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		fmt.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		fmt.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				fmt.Printf("%s", yyStatname(yystate))
				fmt.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					fmt.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				fmt.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		fmt.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line casting.y:23
		{
		 var v Value
		 v = Box{Type: ATOM, Value: yyS[yypt-0].val,}
		 yylex.(*lex).NewBox(v)  
	   }
	case 2:
		//line casting.y:28
		{
		 var v Value
		 v = Box{Type: ATOM, Value: yyS[yypt-1].val, Box: yyS[yypt-0].values}
		 yylex.(*lex).NewBox(v)
	  }
	case 3:
		//line casting.y:33
		{
	  var v Value
	  v = Box{Type: BINDER, Value: yyS[yypt-2].val, Box: yyS[yypt-0].values}
	  yylex.(*lex).NewBox(v)
	 }
	case 4:
		//line casting.y:38
		{
	  var v Value
		if vals, ok := yyS[yypt-0].val.(Val); ok {
		v = Box{Type: ATOM, Value: yyS[yypt-1].val, Box: vals}
		yylex.(*lex).NewBox(v)
	   }
	 }
	case 5:
		//line casting.y:47
		{
	  var newval Val
	  newval = Val{Box{Type:ATOM, Value:yyS[yypt-0].val}}
	  yyVAL.val = newval
		}
	case 6:
		//line casting.y:52
		{
	            var newval Val
	            newval = Val{Box{Type: STRING, Value: yyS[yypt-0].val}}
	            yyVAL.values = newval
	         }
	case 7:
		//line casting.y:57
		{
	            var newval Val
	            result := operate_string(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
	            newval = Val{Box{Type: STRING, Value: result}}
	            yyVAL.values = newval
	         }
	}
	goto yystack /* stack new state and value */
}
