package servers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/hash"
)

func fetchHashFromParams(hashAdapter hash.Adapter, w http.ResponseWriter, r *http.Request, keyname string) *hash.Hash {
	vars := mux.Vars(r)
	if hashAsStr, ok := vars[keyname]; ok {
		hash, err := hashAdapter.FromString(hashAsStr)
		if err != nil {
			renderBadRequest(w, err, []byte(invalidHashErrorOutput))
			return nil
		}

		return hash
	}

	output := fmt.Sprintf(missingParamErrorOutput, keyname)
	renderBadRequest(w, errors.New(output), []byte(output))
	return nil
}

func fetchIDFromParams(w http.ResponseWriter, r *http.Request, keyname string) *uuid.UUID {
	vars := mux.Vars(r)
	if idAsString, ok := vars[keyname]; ok {
		id, err := uuid.FromString(idAsString)
		if err != nil {
			renderBadRequest(w, err, []byte(invalidIDErrorOutput))
			return nil
		}

		return &id
	}

	output := fmt.Sprintf(missingParamErrorOutput, keyname)
	renderBadRequest(w, errors.New(output), []byte(output))
	return nil
}

func renderInsToJSON(w http.ResponseWriter, ins interface{}, err error) {
	panic(errors.New("TODO: finish the renderInsToJSON func in the restapi server helper"))
}

func renderSuccess(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func renderError(w http.ResponseWriter, err error, output []byte) {
	log.Printf("Error: %s\n", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(output)
}

func renderBadRequest(w http.ResponseWriter, err error, output []byte) {
	log.Printf("BadRequest: %s\n", err.Error())
	w.WriteHeader(http.StatusBadRequest)
	w.Write(output)
}
