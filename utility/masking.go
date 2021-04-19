package utility

func BookbankMasking(in string) string {
	if in == "" {
		return ("*** ****** *")
	}
	if len(in) < 10 {
		return ("*** ****** *")
	}
	return ("***" + " ***" + in[6:9] + " " + in[9:10])
}
