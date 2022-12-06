package adapters

import (
	"encoding/json"
	"net/http"

	core "github.com/argSea/portfolio_blog_api/argHex/core/user"
	"github.com/gorilla/mux"
)

//FROM USER TO APP
type userMuxAdapter struct {
	user core.UserCRUDService
}

type userResponseObject struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Users  []interface{} `json:"users"`
}

type erroredResponseObject struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type userlessResponseObject struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func NewUserMuxAdapter(user core.UserCRUDService, m *mux.Router) *userMuxAdapter {
	u := userMuxAdapter{
		user: user,
	}

	m.HandleFunc("/", u.Create).Methods("POST")
	m.HandleFunc("/{id}/", u.Get).Methods("GET")
	m.HandleFunc("/{id}/", u.Update).Methods("PUT")
	m.HandleFunc("/{id}/", u.Delete).Methods("DELETE")

	return &u
}

func (u userMuxAdapter) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//does some stuff

	resp := userlessResponseObject{
		Status: "ok",
		Code:   200,
	}

	json.NewEncoder(w).Encode(resp)

	r.Body.Close()
}

func (u userMuxAdapter) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	user_data := u.user.Read(id)

	response := userResponseObject{
		Status: "ok",
		Code:   200,
	}

	response.Users = append(response.Users, user_data)

	json.NewEncoder(w).Encode(response)

	r.Body.Close()
}

func (u userMuxAdapter) Update(w http.ResponseWriter, r *http.Request) {

}

func (u userMuxAdapter) Delete(w http.ResponseWriter, r *http.Request) {

}
