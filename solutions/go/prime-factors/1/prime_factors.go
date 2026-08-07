package primefactors

func Factors(n int64) []int64 {
    var outputs []int64
	for n != 1 {
        for i := int64(2);; i++ {
            if n % i == 0 {
                n = n / i
                outputs = append(outputs, i)
                break
            }
        }
    }
    return outputs
}
