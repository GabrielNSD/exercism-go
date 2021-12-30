package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	valueUnit, existsUnit := units[unit]
	valueBill, existsBill := bill[item]
	if !existsUnit {
		return false
	}

	if existsBill {
		bill[item] = valueBill + valueUnit
	} else {
		bill[item] = valueUnit
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valueUnit, existsUnit := units[unit]
	valueBill, existsBill := bill[item]
	if !existsBill || !existsUnit {
		return false
	}
	newQuantity := valueBill - valueUnit
	if newQuantity < 0 {
		//bill[item] = 0
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
	valueBill, existsBill := bill[item]
	return valueBill, existsBill
}
