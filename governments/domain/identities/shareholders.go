package identities

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/governments/domain/governments"
)

type shareHolders struct {
	mp   map[string]ShareHolder
	list []ShareHolder
}

func createShareHolders(
	mp map[string]ShareHolder,
	list []ShareHolder,
) ShareHolders {
	out := shareHolders{
		mp:   mp,
		list: list,
	}

	return &out
}

// All returns the shareHolder list
func (obj *shareHolders) All() []ShareHolder {
	return obj.list
}

// Fetch fetches a shareHolder by government
func (obj *shareHolders) Fetch(gov governments.Government) (ShareHolder, error) {
	keyname := gov.ID().String()
	if holder, ok := obj.mp[keyname]; ok {
		return holder, nil
	}

	str := fmt.Sprintf("there is no ShareHolder for the given Government (ID: %s)", keyname)
	return nil, errors.New(str)
}
