package resumeAdapters

import (
	"encoding/json"
	"net/http"

	core "github.com/argSea/portfolio_blog_api/argHex/core/resume"
	"github.com/gorilla/mux"
)

//FROM USER TO APP
type resumeMuxAdapter struct {
	resume core.ResumeCRUDService
}

type resumeResponseObject struct {
	Status  string        `json:"status"`
	Code    int           `json:"code"`
	Resumes []interface{} `json:"resumes"`
}

type erroredResponseObject struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type resumelessResponseObject struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func NewResumeMuxAdapter(resume core.ResumeCRUDService, m *mux.Router) *resumeMuxAdapter {
	u := resumeMuxAdapter{
		resume: resume,
	}

	m.HandleFunc("/", u.Create).Methods("POST")
	m.HandleFunc("/{id}/", u.Get).Methods("GET")
	m.HandleFunc("/{id}/", u.Update).Methods("PUT")
	m.HandleFunc("/{id}/", u.Delete).Methods("DELETE")

	return &u
}

func (res resumeMuxAdapter) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	//does some stuff

	resp := resumelessResponseObject{
		Status: "ok",
		Code:   200,
	}

	json.NewEncoder(w).Encode(resp)

	r.Body.Close()
}

func (res resumeMuxAdapter) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	resume_data := res.resume.Read(id)

	response := resumeResponseObject{
		Status: "ok",
		Code:   200,
	}

	response.Resumes = append(response.Resumes, resume_data)

	json.NewEncoder(w).Encode(response)

	r.Body.Close()
}

func (res resumeMuxAdapter) Update(w http.ResponseWriter, r *http.Request) {

}

func (res resumeMuxAdapter) Delete(w http.ResponseWriter, r *http.Request) {

}
