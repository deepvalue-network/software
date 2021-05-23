package bundles

import (
	"github.com/deepvalue-network/software/pangolin/application"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewPangolin creates a new pangolin application instance
func NewPangolin(grammarFilePath string, currentDirPath string) application.Application {
	wsEvent, err := parsers.NewWhiteSpaceEvent("_whiteSpace")
	if err != nil {
		panic(err)
	}

	app, err := application.NewBuilder().Create().WithGrammarFilePath(grammarFilePath).WithEvents([]lexers.Event{
		wsEvent,
	}).WithCurrentDirPath(currentDirPath).Now()

	if err != nil {
		panic(err)
	}

	return app
}
