package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kosuda/golang-web/db"
	"github.com/kosuda/golang-web/model"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

const (
	playlistColname       = "playlist"
	playlistDetailColname = "playlist_detail"
)

// MusicGet function
func PlaylistGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("access to user get api")

	name := params.ByName("name")

	var data model.Playlist
	if err := db.Find(playlistColname, bson.M{"name": name}, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data2 model.PlaylistDetail
	if err := db.Find(playlistDetailColname, bson.M{"name": name}, &data2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if id == "" {
	// 	log.Print("get all")
	//
	// 	var data model.MusicSlice
	// 	db.Find(colname, nil, &data)
	// 	if json, err := data.MarshalJSON(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	} else {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.Write(json)
	// 	}
	// } else {
	// 	log.Printf("get by id: id = %s", id)
	//
	// 	var data model.Artist
	// 	db.FindByID(colname, id, &data)
	// 	if json, err := data.MarshalJSON(); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	} else {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.Write(json)
	// 	}
	// }
}

// MusicUpsert function
func PlaylistUpsert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

// // MusicDelete function
func PlaylistDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
