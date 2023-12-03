package in_adapter

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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

	m.HandleFunc("", p.GetAll).Methods("GET")
	m.HandleFunc("", p.Create).Methods("POST")
	m.HandleFunc("/{id}", p.Get).Methods("GET")
	m.HandleFunc("/{id}", p.Update).Methods("PUT")
	m.HandleFunc("/{id}", p.Delete).Methods("DELETE")

	m.HandleFunc("/", p.GetAll).Methods("GET")
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

	json.NewEncoder(w).Encode(response.Projects[0])
}

func (p projectMuxAdatper) GetAll(w http.ResponseWriter, r *http.Request) {
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

	limit := int64(0)
	offset := int64(0)
	sort := ""

	if nil != r.URL.Query()["limit"] {
		// convert string to int64
		i, ierr := strconv.ParseInt(r.URL.Query()["limit"][0], 10, 64)

		if nil != ierr {
			// do nothing
		} else {
			limit = i
		}
	}

	if nil != r.URL.Query()["offset"] {
		// convert string to int64
		i, ierr := strconv.ParseInt(r.URL.Query()["offset"][0], 10, 64)

		if nil != ierr {
			// do nothing
		} else {
			offset = i
		}
	}

	if nil != r.URL.Query()["sort"] {
		sort = r.URL.Query()["sort"][0]

		if "" == sort {
			sort = "nil"
		}
	}

	// if limit and offset are 0, check for range query string
	if 0 == limit && 0 == offset {
		if nil != r.URL.Query()["range"] {
			// convert [0, 10] to limit = 10, offset = 0
			range_str := r.URL.Query()["range"][0]
			range_str = strings.Replace(range_str, "[", "", -1)
			range_str = strings.Replace(range_str, "]", "", -1)

			range_arr := strings.Split(range_str, ",")
			limit, _ = strconv.ParseInt(range_arr[1], 10, 64)
			offset, _ = strconv.ParseInt(range_arr[0], 10, 64)
		}
	}

	// look for filter param and decode it
	query := r.URL.Query().Get("filter")

	// log query
	log.Println(query)

	if query == "" {
		// throw 404
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    404,
			Message: "No filter provided",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var filter map[string]string
	json.Unmarshal([]byte(query), &filter)

	// check for userID in filter
	userID, ok := filter["userID"]

	if !ok {
		// throw 404
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    404,
			Message: "No userID provided",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	projects, count, err := p.project.GetByUserID(userID)

	if nil != err {
		// throw 404
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    404,
			Message: err,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := data_objects.ProjectResponseObject{
		Status: "ok",
		Code:   200,
		Count:  count,
	}

	for i := 0; i < len(projects); i++ {
		response.Projects = append(response.Projects, projects[i])
	}

	total := len(response.Projects)

	w.Header().Add("Content-Range", "users "+strconv.FormatInt(offset, 10)+"-"+strconv.FormatInt(offset+limit, 10)+"/"+strconv.FormatInt(int64(total), 10))
	w.Header().Add("range", "users "+strconv.FormatInt(offset, 10)+"-"+strconv.FormatInt(offset+limit, 10)+"/"+strconv.FormatInt(int64(total), 10))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.Projects)

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
