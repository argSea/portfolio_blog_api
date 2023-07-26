package in_adapter

import (
	"encoding/json"
	"net/http"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/gorilla/mux"
)

type projectMuxAdatper struct {
	project in_port.ProjectCRUDService
}

func NewProjectMuxAdapter(proj in_port.ProjectCRUDService, m *mux.Router) *projectMuxAdatper {
	p := projectMuxAdatper{
		project: proj,
	}

	m.HandleFunc("/", p.Create).Methods("POST")
	m.HandleFunc("/{id}/", p.Get).Methods("GET")
	m.HandleFunc("/{id}/", p.Update).Methods("PUT")
	m.HandleFunc("/{id}/", p.Delete).Methods("DELETE")

	return &p
}

func (p projectMuxAdatper) Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			json.NewEncoder(w).Encode(response)
		}
	}()

	var project domain.Project
	json.NewDecoder(r.Body).Decode(&project)

	new_id, err := p.project.Create(project)
	var resp interface{}

	if nil != err {
		resp = data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: err,
		}
	} else {
		resp = data_objects.NewProjectResponseObject{
			Status:    "ok",
			Code:      200,
			ProjectID: new_id,
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func (p projectMuxAdatper) Get(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			json.NewEncoder(w).Encode(response)
		}
	}()

	id := mux.Vars(r)["id"]
	project_data := p.project.Read(id)

	response := data_objects.ProjectResponseObject{
		Status: "ok",
		Code:   200,
	}

	response.Projects = append(response.Projects, project_data)

	json.NewEncoder(w).Encode(response)
}

func (p projectMuxAdatper) Update(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			json.NewEncoder(w).Encode(response)
		}
	}()

	var project domain.Project
	json.NewDecoder(r.Body).Decode(&project)

	id := mux.Vars(r)["id"]
	project.Id = id
	updated_err := p.project.Update(project)

	var resp interface{}

	if nil != updated_err {
		resp = data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: updated_err,
		}
	} else {
		resp = data_objects.ItemLessResponseObject{
			Status: "ok",
			Code:   200,
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func (p projectMuxAdatper) Delete(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			json.NewEncoder(w).Encode(response)
		}
	}()

	project := domain.Project{}

	id := mux.Vars(r)["id"]
	project.Id = id
	deleted_err := p.project.Delete(project)

	var resp interface{}

	if nil != deleted_err {
		resp = data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: deleted_err,
		}
	} else {
		resp = data_objects.ItemLessResponseObject{
			Status: "ok",
			Code:   200,
		}
	}

	json.NewEncoder(w).Encode(resp)
}
