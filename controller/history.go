package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kosuda/golang-web/db"
	"github.com/kosuda/golang-web/model"
	"log"
	"net/http"
	"time"
)

const (
	historyColname = "history"
)

// MusicUpsert function
func HistoryUpsert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("access to user put api")

	id := params.ByName("id")

	if id == "" {
		log.Print("put without _id")
	} else {
		log.Print("put with id")
	}

	if r.Body != nil {
		defer r.Body.Close()
		var data model.PlayHistory

		data.MusicID = id
		data.CreatedAt = time.Now().String()

		if err := db.Upsert(musicColname, data.ID, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\":\"ok\"}"))
	}
}
