package resistorcolorduo

var colorsMap = map[string]int{
    "black": 0,
    "brown": 1,
    "red": 2,
    "orange": 3,
    "yellow": 4,
    "green": 5,
    "blue": 6,
    "violet": 7,
    "grey": 8,
    "white": 9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
    output := 0
    limit := 2
    if len(colors) < 2 {
        limit = len(colors)
    }
	for i, j := 0, 10; i < limit; i, j = i+1, j/10 {
        color, ok := colorsMap[colors[i]]
        if !ok {
            return 0
        }
        output += color * j
    }
    return output
}
