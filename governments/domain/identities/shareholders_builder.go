package identities

import (
	"errors"
	"fmt"
)

type shareHoldersBuilder struct {
	list []ShareHolder
}

func createShareHoldersBuilder() ShareHoldersBuilder {
	out := shareHoldersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *shareHoldersBuilder) Create() ShareHoldersBuilder {
	return createShareHoldersBuilder()
}

// WithShareHolders add shareHolders to the builder
func (app *shareHoldersBuilder) WithShareHolders(shareHolders []ShareHolder) ShareHoldersBuilder {
	app.list = shareHolders
	return app
}

// Now builds a new ShareHolders instance
func (app *shareHoldersBuilder) Now() (ShareHolders, error) {
	if app.list == nil {
		app.list = []ShareHolder{}
	}

	mp := map[string]ShareHolder{}
	for _, oneHolder := range app.list {
		keyname := oneHolder.Government().ID().String()
		mp[keyname] = oneHolder
	}

	if len(mp) != len(app.list) {
		str := fmt.Sprintf("there is ShareHolder(s) that have the same Government, only 1 ShareHolder per Government is allowed")
		return nil, errors.New(str)
	}

	return createShareHolders(mp, app.list), nil
}
