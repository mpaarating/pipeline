package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cloudnativego/gogo-engine"
	"github.com/google/uuid"
	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		json.Unmarshal(payload, &newMatchRequest)

		newMatch := gogo.NewMatch(newMatchRequest.GridSize, "Sam", "Carol")
		repo.addMatch(newMatch)
		guid := uuid.New().String()
		w.Header().Add("Location", "/matches/"+guid)
		formatter.JSON(w, http.StatusCreated, &newMatchResponse{ID: guid, GridSize: newMatch.GridSize})
	}
}
