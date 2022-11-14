package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
	"github.com/gorilla/mux"
)

//Handler
type projectService struct {
	projCase entity.ProjectUsecase
}

func NewProjectService(m *mux.Router, user entity.ProjectUsecase) {
	handler := &projectService{
		projCase: user,
	}

	m.HandleFunc("/", handler.Create).Methods("POST")
	m.HandleFunc("/", handler.GetMany).Methods("GET")
	m.HandleFunc("/{id}/", handler.Get).Methods("GET")
	m.HandleFunc("/byUser/{userID}/", handler.GetByUser).Methods("GET")
	m.HandleFunc("/{id}/", handler.Update).Methods("PUT")
	m.HandleFunc("/{id}/", handler.Delete).Methods("DELETE")
}

func (p *projectService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	//Decode
	newProject := entity.Project{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProject)

	//Make model
	finalModel := struct {
		Status  string      `json:"status"`
		Code    int         `json:"code"`
		Message string      `json:"message,omitempty"`
		Items   interface{} `json:"items"`
	}{
		Status: "ok",
		Code:   200,
	}

	createdUser, err := p.projCase.Save(ctx, newProject)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = createdUser
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)
}

func (p *projectService) GetMany(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalModel := struct {
		Status  string      `json:"status"`
		Code    int         `json:"code"`
		Total   int         `json:"total"`
		Count   int         `json:"currentCount"`
		Message string      `json:"message,omitempty"`
		Items   interface{} `json:"items"`
	}{
		Status: "ok",
		Code:   200,
	}

	queries := r.URL.Query()

	limit, lErr := strconv.Atoi(queries.Get("limit"))

	if nil != lErr {
		limit = 20
	}

	fmt.Println(limit)

	offset, oErr := strconv.Atoi(queries.Get("offset"))

	if nil != oErr {
		offset = 0
	}

	fmt.Println(offset)

	sortProj := entity.ProjectSort{
		Priority: 1,
	}

	projects, totalCount, err := p.projCase.GetProjects(ctx, int64(limit), int64(offset), sortProj)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = projects
		finalModel.Count = len(*projects)
		finalModel.Total = int(totalCount)
	}

	json.NewEncoder(w).Encode(finalModel)

	r.Body.Close()
}

func (p *projectService) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)
	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["id"]

	tempProject, err := p.projCase.GetByProjectID(ctx, id)
	// tempProject, err := u.userCase.GetUserByUserName(ctx, "saltosk")

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = tempProject
	} else {
		finalModel.Items = tempProject
	}

	json.NewEncoder(w).Encode(finalModel)

	r.Body.Close()
}

func (p *projectService) GetByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalModel := struct {
		Status  string      `json:"status"`
		Code    int         `json:"code"`
		Total   int         `json:"total"`
		Count   int         `json:"currentCount"`
		Message string      `json:"message,omitempty"`
		Items   interface{} `json:"items"`
	}{
		Status: "ok",
		Code:   200,
	}

	queries := r.URL.Query()

	limit, lErr := strconv.Atoi(queries.Get("limit"))

	if nil != lErr {
		limit = 20
	}

	fmt.Println(limit)

	offset, oErr := strconv.Atoi(queries.Get("offset"))

	if nil != oErr {
		offset = 0
	}

	fmt.Println(offset)

	sortProj := entity.ProjectSort{
		Priority: 1,
	}

	userID := mux.Vars(r)["userID"]

	projects, totalCount, err := p.projCase.GetProjectsByUserID(ctx, userID, int64(limit), int64(offset), sortProj)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = projects
		finalModel.Count = len(*projects)
		finalModel.Total = int(totalCount)
	}

	json.NewEncoder(w).Encode(finalModel)

	r.Body.Close()
}

func (p *projectService) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	newProjectDetails := entity.Project{}
	json.NewDecoder(r.Body).Decode(&newProjectDetails)
	newProjectDetails.Id = mux.Vars(r)["id"]

	updatedProject, err := p.projCase.Update(ctx, newProjectDetails)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = updatedProject
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)

	r.Body.Close()
}

func (p *projectService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["id"]

	err := p.projCase.Delete(ctx, id)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Message = "Deleted"
		finalModel.Items = nil
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)

	r.Body.Close()
}
