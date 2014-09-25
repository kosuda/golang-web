package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kosuda/golang-web/db"
	"github.com/kosuda/golang-web/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	musicColname = "music"
)

// MusicGet function
func MusicGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("access to user get api")

	id := params.ByName("id")
	var artist_id string
	var limit int
	var start int
	var title string

	if val, ok := r.URL.Query()["artist_id"]; ok {
		artist_id = val[0]
	} else {
		artist_id = ""
	}

	if val, ok := r.URL.Query()["limit"]; ok {
		limit, _ = strconv.Atoi(val[0])
	} else {
		limit = 100
	}

	if val, ok := r.URL.Query()["title"]; ok {
		title = "/" + val[0] + "/"
	} else {
		title = ""
	}

	if val, ok := r.URL.Query()["start"]; ok {
		start, _ = strconv.Atoi(val[0])
	} else {
		start = 0
	}

	if id == "" {
		log.Print("get all")

		var data model.Music
		db.Find(colname, nil, &data)
		if json, err := data.MarshalJSON(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	} else {
		log.Print("get all")

		var data model.MusicSlice
		db.Find(colname, nil, &data)
		if json, err := data.MarshalJSON(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	}
}

// MusicUpsert function
func MusicUpsert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("access to user put api")

	id := params.ByName("id")

	if id == "" {
		log.Print("put without _id")
	} else {
		log.Print("put with id")
	}

	if r.Body != nil {
		defer r.Body.Close()
		var data model.Music

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal([]byte(body), &data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		if err := db.Upsert(musicColname, data.ID, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\":\"ok\"}"))
	}
}

// MusicDelete function
func MusicDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("access to user delete api")

	id := params.ByName("id")

	if id == "" {
		log.Print("remove by query")

		if err := db.Remove(colname, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Print("remove by id")

		if err := db.RemoveByID(musicColname, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("{\"status\":\"ok\"}"))
}
