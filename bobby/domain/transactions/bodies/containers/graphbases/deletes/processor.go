package deletes

import (
	"github.com/steve-care-software/products/bobby/domain/structures"
	"github.com/steve-care-software/products/bobby/domain/structures/graphbases"
)

type processor struct {
	structureBuilder    structures.Builder
	structureRepository structures.Repository
}

func createProcessor(
	structureBuilder structures.Builder,
	structureRepository structures.Repository,
) Processor {
	out := processor{
		structureBuilder:    structureBuilder,
		structureRepository: structureRepository,
	}

	return &out
}

// Execute processes a transaction
func (app *processor) Execute(trx Transaction) ([]structures.Structure, error) {
	selector := trx.Graphbase()
	list, err := app.structureRepository.Search(selector)
	if err != nil {
		return nil, err
	}

	deleted := []structures.Structure{}
	mustBeEmpty := trx.MustBeGraphbaseEmpty()
	for _, oneStructure := range list {
		content := oneStructure.Content()
		if !content.IsGraphbase() {
			//err
		}

		base := content.Graphbase()
		if mustBeEmpty {
			if !app.isGraphbaseEmpty(base) {
				// error not empty
			}
		}

		ins, err := app.structureBuilder.Create().WithGraphbase(base).IsDeleted().Now()
		if err != nil {
			return nil, err
		}

		deleted = append(deleted, ins)
	}

	return deleted, nil
}

func (app *processor) isGraphbaseEmpty(graphbase graphbases.Graphbase) bool {
	return true
}
