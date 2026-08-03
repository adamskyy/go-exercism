package grains

import "errors"

func Square(number int) (uint64, error) {
    if number < 1 || number > 64 {
        return uint64(0), errors.New("Fatal Error")
    }
	if number == 1 {
        return uint64(1), nil
    }
    val, _ := Square(number - 1) 
    return uint64(val * 2), nil
}

func Total() uint64 {
	sum := uint64(0)
    for i := 1; i < 65; i++ {
        val, err := Square(i)
        if err == nil {
            sum += val
        }
    }
    return sum
}
