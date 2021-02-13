package disks

import (
	"errors"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	mined_link "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/events"
)

func initEventManager() (events.Manager, error) {

	builder := events.NewBuilder()

	// mined block:
	minedBlockOnBlockDelete, err := builder.Create().WithIdentifier(EventBlockDelete).OnEnter(func(data interface{}, event events.Event) error {
		if ins, ok := data.(blocks.Block); ok {
			internalServiceBlockMined.DeleteByBlock(ins)
			return nil
		}

		return errors.New("the event data was expected to be a block instance")
	}).Now()

	if err != nil {
		return nil, err
	}

	// link:
	linkOnBlockDelete, err := builder.Create().WithIdentifier(EventBlockDelete).OnEnter(func(data interface{}, event events.Event) error {
		if ins, ok := data.(blocks.Block); ok {
			internalServiceLink.DeleteByBlock(ins)
			return nil
		}

		return errors.New("the event data was expected to be a block instance")
	}).Now()

	if err != nil {
		return nil, err
	}

	linkOnMinedLinkDelete, err := builder.Create().WithIdentifier(EventLinkMinedDelete).OnEnter(func(data interface{}, event events.Event) error {
		if ins, ok := data.(mined_link.Link); ok {
			return internalServiceLink.DeleteByMinedLinkHash(ins.Hash())
		}

		return errors.New("the event data was expected to be a mined link instance")
	}).Now()

	if err != nil {
		return nil, err
	}

	// creates the manager:
	manager := events.NewManagerFactory().Create()

	// add the events:
	err = manager.AddList([]events.Event{
		minedBlockOnBlockDelete,
		linkOnBlockDelete,
		linkOnMinedLinkDelete,
	})

	if err != nil {
		return nil, err
	}

	return manager, nil
}
