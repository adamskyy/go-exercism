package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var output int
    switch card {
        case "ace":
        	output = 11
        case "two":
        	output = 2
        case "three":
        	output = 3
        case "four":
        	output = 4
        case "five":
        	output = 5
        case "six":
        	output = 6
        case "seven":
        	output = 7
        case "eight":
        	output = 8
        case "nine":
        	output = 9
        case "ten", "jack", "queen", "king":
        	output = 10
        default:
        	output = 0    
    }
    return output
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
    var decision string
    switch {
        case value == 22:
        	decision = "P"
    	case value == 21 && dealerValue < 10:
        	decision = "W"
        case value == 21 && dealerValue >= 10:
        	decision = "S"
        case value >= 17 && value <= 20:
        	decision = "S"
        case (value >= 12 && value <= 16) && dealerValue >= 7:
        	decision = "H"
        case (value >= 12 && value <= 16) && dealerValue < 7:
        	decision = "S"
        default:
        	decision = "H"
    }
    return decision
}
