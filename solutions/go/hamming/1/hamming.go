package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("Input should be equal")
    }
    seq_len := len(a)
    sum := 0
    for i := 0; i < seq_len; i++ {
        if a[i] != b[i] {
            sum++
        }
    }
    return sum, nil
}
