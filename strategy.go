package main

func IsSameColor(pokers ...Poker) bool {
	if len(pokers) <= 1 {
		return true
	}
	p0 := pokers[0]
	for idx, p := range pokers {
		if idx > 0 && !p.SameColor(p0) {
			return false
		}
	}
	return true
}

func IsSameValue(pokers ...Poker) bool {
	if len(pokers) <= 1 {
		return true
	}
	p0 := pokers[0]
	for idx, p := range pokers {
		if idx > 0 && !p.SameValue(p0) {
			return false
		}
	}
	return true
}

func SameValueNum() int8 {
	return 0
}

func SameColorNum() int8 {
	return 0
}
