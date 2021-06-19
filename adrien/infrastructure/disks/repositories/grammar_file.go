package repositories

import (
    "github.com/deepvalue-network/software/adrien/domain/grammars"
     "io/ioutil"
     "path/filepath"
      "encoding/json"
)

type grammarFileContent struct {
    Root string `json:"root"`
    Tokens string `json:"tokens"`
    Rules string `json:"rules"`
    Channels string `json:"channels"`
}

type grammarFile struct {
    fileBuilder grammars.FileBuilder
    basePath string
}

func createGrammarFile(
    fileBuilder grammars.FileBuilder,
    basePath string,
    ) grammars.FileRepository {
    out := grammarFile{
        fileBuilder: fileBuilder,
        basePath: basePath,
    }

    return &out
}

// Retrieve retrieves a File instance from a relative path
func (app *grammarFile) Retrieve(relativePath string) (grammars.File, error) {
    path := filepath.Join(app.basePath, relativePath)
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    ins := new(grammarFileContent)
    err = json.Unmarshal(data, ins)
    if err != nil {
        return nil, err
    }

    tokensPath := filepath.Join(app.basePath, ins.Tokens)
    rulesPath := filepath.Join(app.basePath, ins.Rules)
    builder := app.fileBuilder.Create().WithRoot(ins.Root).WithTokensPath(tokensPath).WithRulesPath(rulesPath)
    if ins.Channels != "" {
        channelsPath := filepath.Join(app.basePath, ins.Channels)
        builder.WithChannelsPath(channelsPath)
    }

    return builder.Now()
}
