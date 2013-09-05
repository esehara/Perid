
// line 1 "ore.rl"

// line 38 "ore.rl"


package miner


// line 11 "ore.go"
var _miner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 
}

var _miner_key_offsets []byte = []byte{
	0, 0, 1, 2, 11, 
}

var _miner_trans_keys []byte = []byte{
	34, 39, 32, 34, 39, 9, 13, 65, 
	90, 97, 122, 48, 57, 65, 90, 97, 
	122, 
}

var _miner_single_lengths []byte = []byte{
	0, 1, 1, 3, 0, 
}

var _miner_range_lengths []byte = []byte{
	0, 0, 0, 3, 3, 
}

var _miner_index_offsets []byte = []byte{
	0, 0, 2, 4, 11, 
}

var _miner_indicies []byte = []byte{
	1, 0, 1, 2, 3, 0, 2, 3, 
	5, 5, 4, 5, 5, 5, 6, 
}

var _miner_trans_targs []byte = []byte{
	1, 3, 2, 3, 0, 4, 3, 
}

var _miner_trans_actions []byte = []byte{
	0, 5, 0, 7, 0, 0, 9, 
}

var _miner_to_state_actions []byte = []byte{
	0, 0, 0, 1, 0, 
}

var _miner_from_state_actions []byte = []byte{
	0, 0, 0, 3, 0, 
}

var _miner_eof_trans []byte = []byte{
	0, 0, 0, 0, 7, 
}

const miner_start int = 3
const miner_first_final int = 3
const miner_error int = 0

const miner_en_main int = 3


// line 43 "ore.rl"

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

    
// line 123 "ore.go"
	{
	cs = miner_start
	ts = 0
	te = 0
	act = 0
	}

// line 94 "ore.rl"
    
// line 133 "ore.go"
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

// line 157 "ore.go"
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
	_trans = int(_miner_indicies[_trans])
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
// line 29 "ore.rl"

te = p+1
{
            material := Ore{
                Token: STRING,
                Value: data[ts:te],
            }
            material.Research()
            ores = append(ores, material) }
		case 3:
// line 36 "ore.rl"

te = p+1

		case 4:
// line 21 "ore.rl"

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
// line 256 "ore.go"
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

// line 271 "ore.go"
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

// line 95 "ore.rl"

    _, _, _ = ts, te, act
    return ores
}
// vim: set syntax=go:
