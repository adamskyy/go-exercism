package primefactors

func Factors(n int64) []int64 {
    var outputs []int64
    for i := int64(2); i * i <= n; {
        if n % i == 0 {
            n /= i
            outputs = append(outputs, i)
        } else {
            i++
        }
    }

    if n != 1 {
        outputs = append(outputs, n)
    }
    return outputs
}
