package casting

func operate_string(prev, oper, next Value) Value {
	var result Value
	prev_str, prev_ok := prev.(string)
	next_str, next_ok := next.(string)
	if prev_ok && next_ok { 
		switch oper {
		case "+":
			return add_str(prev_str, next_str)
		}
	}
	return result
}

func add_str(prev, next string) string {
	var new_p, new_n string
	new_p = prev[1: len(prev) - 1]
	new_n = next[1: len(next) - 1]
	return "'" + new_p + new_n + "'"
}
