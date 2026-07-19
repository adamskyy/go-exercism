package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMeasurs := map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
	}
    return unitMeasurs
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unit_retr, exists := units[unit]
    if exists == false {
        return false
    }
    bill[item] = bill[item] + unit_retr
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	givItem, existsItem := bill[item]
    givUnit, existsUnit := units[unit]
    if existsItem == false || existsUnit == false {
        return false
    }
    newQuantity := givItem - givUnit
    if newQuantity < 0 {
        return false
    } else if newQuantity == 0 {
        delete(bill, item)
    } else {
        bill[item] = newQuantity
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	givItem, existsItem := bill[item]
    if existsItem == false {
        return 0, false
    }
    return givItem, true
}
