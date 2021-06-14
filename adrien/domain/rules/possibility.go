package rules

type possibility struct {
	list   []string
	amount Amount
}

func createPossibility(
	list []string,
	amount Amount,
) Possibility {
	return createPossibilityInternally(list, amount)
}

func createPossibilityInternally(
	list []string,
	amount Amount,
) Possibility {
	out := possibility{
		list:   list,
		amount: amount,
	}

	return &out
}

// List returns the list
func (obj *possibility) List() []string {
	return obj.list
}

// Amount returns the amount
func (obj *possibility) Amount() Amount {
	return obj.amount
}
