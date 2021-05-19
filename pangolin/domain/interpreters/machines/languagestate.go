package machines

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
)

type languageState struct {
	currentToken lexers.NodeTree
}

func createLanguageState() LanguageState {
	out := languageState{
		currentToken: nil,
	}

	return &out
}

// ChangeCurrentToken changes the current token
func (app *languageState) ChangeCurrentToken(tok lexers.NodeTree) {
	app.currentToken = tok
}

// CurrentToken returns the current token, if any
func (app *languageState) CurrentToken() lexers.NodeTree {
	return app.currentToken
}
