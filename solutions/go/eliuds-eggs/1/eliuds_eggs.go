package eliudseggs

func EggCount(displayValue int) int {
    sum := 0
    for displayValue >= 1 {
        if isLastBitEqualOne(displayValue) {
        	sum++
    	}
        displayValue >>= 1
    }
    return sum + displayValue
}

func isLastBitEqualOne(n int) bool {
    return n & 1 == 1
}