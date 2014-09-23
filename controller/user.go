package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kosuda/golang-web/db"
	"github.com/kosuda/golang-web/model"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	colname = "user"
)

// UserGet function
func UserGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("access to user get api")

	id := params.ByName("id")

	if id == "" {
		log.Print("get all")

		var users model.UserSlice
		db.Find(colname, nil, &users)
		if json, err := users.MarshalJSON(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	} else {
		log.Printf("get by id: id = %s", id)

		var user model.User
		db.FindByID(colname, id, &user)
		if json, err := user.MarshalJSON(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(json)
		}
	}
}

// UserUpsert function
func UserUpsert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("access to user put api")

	id := params.ByName("id")

	if id == "" {
		log.Print("put without _id")
	} else {
		log.Print("put with id")
	}

	if r.Body != nil {
		defer r.Body.Close()
		var user model.User

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal([]byte(body), &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		if err := db.Upsert(colname, id, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\":\"ok\"}"))
	}
}

// UserDelete function
func UserDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

		if err := db.RemoveByID(colname, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"ok\"}"))
}

// UserUpdate function
func UserUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print("access to user update api")

	id := params.ByName("id")

	defer r.Body.Close()

	if id == "" {
		log.Print("update by query")
	} else {
		log.Print("update by id")
		var user model.User

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal([]byte(body), &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		if err := db.UpdateByID(colname, id, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\":\"ok\"}"))
	}
}
