package forstring

func shiftASCII(s string, step int) string {
	res_str := ""
	for i := 0; i < len(s); i++ {
		res_str += string(s[i]+byte(step))
	}
	return res_str
}