package stringdivisor

type EnhancedString string

func (s *EnhancedString) isDividedBy(d string) bool {
	ls := len(string(*s))
	ld := len(string(d))
	if ls < ld || ls%ld != 0 {
		return false
	}
	ss := []rune(*s)
	c := ls / ld
	for i := 0; i < c; i++ {
		if string(ss[i*ld:(i+1)*ld]) != d {
			return false
		}
	}
	return true
}

func (s EnhancedString) Cut(start, end int) EnhancedString {
	ss := []rune(string(s))
	return (EnhancedString)(string(ss[start:end]))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	var es, ed, eod EnhancedString
	var ld int

	if len(str1) < len(str2) {
		es = EnhancedString(str2)
		ed = EnhancedString(str1)
		eod = EnhancedString(str1)
		ld = len(str1)
	} else {
		es = EnhancedString(str1)
		ed = EnhancedString(str2)
		eod = EnhancedString(str2)
		ld = len(str2)
	}
	for i := ld; i > 0; i-- {
		ed = ed.Cut(0, i)
		if es.isDividedBy(string(ed)) && eod.isDividedBy(string(ed)) {
			return string(ed)
		}
	}	
	return ""
}