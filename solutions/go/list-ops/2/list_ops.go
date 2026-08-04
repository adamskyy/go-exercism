package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
        initial = fn(initial, v)
    }
    return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
    for i := len(s) - 1; i >= 0; i-- {
        initial = fn(s[i], initial)
    }
    return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	output := make(IntList, 0, len(s))
    for _, v := range s {
        if fn(v) {
            output = append(output, v)
        }
    }
    return output
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
    output := make(IntList, len(s))
    for i, v := range s {
        output[i] = fn(v)
    }
    return output
}

func (s IntList) Reverse() IntList {
	output := make(IntList, len(s))
    for i, value := range s {
    	output[len(s)-1-i] = value
    }
    return output
}

func (s IntList) Append(lst IntList) IntList {
    output := make(IntList, 0, len(s) + len(lst))
    output = append(output, s...)
    output = append(output, lst...)
	return output
}

func (s IntList) Concat(lists []IntList) IntList {
    totalLen := len(s)

	for _, list := range lists {
		totalLen += len(list)
	}
	output := make(IntList, 0, totalLen)
	output = append(output, s...)
	for _, list := range lists {
		output = append(output, list...)
	}
	return output
}
