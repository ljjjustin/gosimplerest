package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ljjjustin/gosimplerest/models"
)

var userRoutes = Routes{
	Route{"UserIndex", "GET", "/users", UserIndex},
	Route{"UserDetail", "GET", "/users/{userId}", UserDetail},
	Route{"UserUpdate", "PUT", "/users/{userId}", UserUpdate},
	Route{"UserCreate", "POST", "/users", UserCreate},
	Route{"UserDelete", "DELETE", "/users/{userId}", UserDelete},
}

func init() {
	for i := range userRoutes {
		routes = append(routes, userRoutes[i])
	}
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["userId"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	user := models.User{Id: id}
	if _, err := models.UserGet(&user); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["userId"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	user := models.User{Id: id}
	if _, err := models.UserGet(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 65535))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if _, err = models.UserUpdate(&user); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.UserGetAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user models.User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 65535))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if err = models.UserCreate(&user); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["userId"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	user := models.User{Id: id}
	if err = models.UserDelete(&user); err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}
