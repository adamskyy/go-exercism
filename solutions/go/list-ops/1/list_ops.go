package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	elemLen := len(s)
    result := initial
    for i := 0; i < elemLen; i++ {
        result = fn(initial, s[i])
        initial = result
    }
    return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	elemLen := len(s) - 1
    result := initial
    for i := elemLen; i >= 0; i-- {
        result = fn(s[i], initial)
        initial = result
    }
    return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var output IntList;
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
    for i := 0; i < len(s); i++ {
        s[i] = fn(s[i])
    }
    return s
}

func (s IntList) Reverse() IntList {
	var output IntList;
    for i := len(s) - 1; i >= 0; i-- {
        output = append(output, s[i])
    }
    return output
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
    var output IntList;
    output = s
    for _, v := range lists {
        output = append(output, v...)
    }
	return output
}
