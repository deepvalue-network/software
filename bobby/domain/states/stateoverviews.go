package states

import "github.com/deepvalue-network/software/bobby/domain/states/overviews"

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
