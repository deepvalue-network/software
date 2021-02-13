package disks

import (
	"errors"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/links"
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
			internalServiceLink.DeleteByMinedLinkHash(ins.Hash())
			return nil
		}

		return errors.New("the event data was expected to be a mined link instance")
	}).Now()

	if err != nil {
		return nil, err
	}

	minedLinkOnLinkDelete, err := builder.Create().WithIdentifier(EventLinkDelete).OnEnter(func(data interface{}, event events.Event) error {
		if ins, ok := data.(links.Link); ok {
			internalServiceLinkMined.DeleteByLink(ins)
			return nil
		}

		return errors.New("the event data was expected to be a link instance")
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
		minedLinkOnLinkDelete,
	})

	if err != nil {
		return nil, err
	}

	return manager, nil
}
