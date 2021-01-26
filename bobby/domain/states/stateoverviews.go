package states

import "github.com/steve-care-software/products/bobby/domain/states/overviews"

type stateOverviews struct {
	state     State
	overviews []overviews.Overview
}

func createStateOverviews(
	state State,
	overviews []overviews.Overview,
) *stateOverviews {
	out := stateOverviews{
		state:     state,
		overviews: overviews,
	}

	return &out
}
