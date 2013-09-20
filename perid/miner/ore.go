
// line 1 "ore.rl"
// -*- mode:go-mode -*-

// line 64 "ore.rl"


package miner


// line 12 "ore.go"
var _miner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
}

var _miner_key_offsets []byte = []byte{
	0, 0, 1, 2, 18, 20, 
}

var _miner_trans_keys []byte = []byte{
	34, 39, 32, 34, 39, 45, 47, 61, 
	9, 13, 42, 43, 48, 57, 65, 90, 
	97, 122, 48, 57, 48, 57, 65, 90, 
	97, 122, 
}

var _miner_single_lengths []byte = []byte{
	0, 1, 1, 6, 0, 0, 
}

var _miner_range_lengths []byte = []byte{
	0, 0, 0, 5, 1, 3, 
}

var _miner_index_offsets []byte = []byte{
	0, 0, 2, 4, 16, 18, 
}

var _miner_trans_targs []byte = []byte{
	3, 1, 3, 2, 3, 1, 2, 3, 
	3, 3, 3, 3, 4, 5, 5, 0, 
	4, 3, 5, 5, 5, 3, 3, 3, 
	
}

var _miner_trans_actions []byte = []byte{
	5, 0, 5, 0, 11, 0, 0, 7, 
	7, 9, 11, 7, 0, 0, 0, 0, 
	0, 15, 0, 0, 0, 13, 15, 13, 
	
}

var _miner_to_state_actions []byte = []byte{
	0, 0, 0, 1, 0, 0, 
}

var _miner_from_state_actions []byte = []byte{
	0, 0, 0, 3, 0, 0, 
}

var _miner_eof_trans []byte = []byte{
	0, 0, 0, 0, 23, 24, 
}

const miner_start int = 3
const miner_first_final int = 3
const miner_error int = 0

const miner_en_main int = 3


// line 69 "ore.rl"

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

    
// line 132 "ore.go"
	{
	cs = miner_start
	ts = 0
	te = 0
	act = 0
	}

// line 126 "ore.rl"
    
// line 142 "ore.go"
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_miner_from_state_actions[cs])
	_nacts = uint(_miner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _miner_actions[_acts - 1] {
		case 1:
// line 1 "NONE"

ts = p

// line 166 "ore.go"
		}
	}

	_keys = int(_miner_key_offsets[cs])
	_trans = int(_miner_index_offsets[cs])

	_klen = int(_miner_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _miner_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _miner_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_miner_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _miner_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _miner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
_eof_trans:
	cs = int(_miner_trans_targs[_trans])

	if _miner_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_miner_trans_actions[_trans])
	_nacts = uint(_miner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _miner_actions[_acts-1] {
		case 2:
// line 30 "ore.rl"

te = p+1
{
            material := Ore{
                Token: STRING,
                Value: data[ts:te],
            }
            material.Research()
            ores = append(ores, material)
		}
		case 3:
// line 46 "ore.rl"

te = p+1
{
			material := Ore {
				Token: OPERATOR,
				Value: data[ts:te],
			}
			material.Research()
			ores = append(ores, material)
		}
		case 4:
// line 54 "ore.rl"

te = p+1
{
			material := Ore {
				Token: BINDER,
				Value: data[ts:te],
			}
			material.Research()
			ores = append(ores, material)
		}
		case 5:
// line 62 "ore.rl"

te = p+1

		case 6:
// line 22 "ore.rl"

te = p
p--
{
            material := Ore{
                Token: ATOM,
                Value: data[ts:te],
            }
            material.Research()
            ores = append(ores, material) 
        }
		case 7:
// line 38 "ore.rl"

te = p
p--
{
			material := Ore {
				Token: INTEGER,
				Value: data[ts:te],
			}
			material.Research()
			ores = append(ores, material)
		}
// line 302 "ore.go"
		}
	}

_again:
	_acts = int(_miner_to_state_actions[cs])
	_nacts = uint(_miner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _miner_actions[_acts-1] {
		case 0:
// line 1 "NONE"

ts = 0

// line 317 "ore.go"
		}
	}

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _miner_eof_trans[cs] > 0 {
			_trans = int(_miner_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

// line 127 "ore.rl"

    _, _, _ = ts, te, act
    return ores
}
// vim: set syntax=go:
