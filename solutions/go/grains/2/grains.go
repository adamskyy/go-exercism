package grains

import "errors"

func Square(number int) (uint64, error) {
    if number < 1 || number > 64 {
        return 0, errors.New("fatal error")
    }
    return uint64(1) << (number - 1), nil
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
