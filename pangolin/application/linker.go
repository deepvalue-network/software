package application

import (
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type linker struct {
	lnk linkers.Linker
}

func createLinker(
	lnk linkers.Linker,
) Linker {
	out := linker{
		lnk: lnk,
	}

	return &out
}

// Execute takes a parsed program and links it
func (app *linker) Execute(parsed parsers.Program) (linkers.Program, error) {
	return app.lnk.Execute(parsed)
}
