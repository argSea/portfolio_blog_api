package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
	"github.com/gorilla/mux"
)

//Handler
type resumeService struct {
	resumeCase entity.ResumeUseCase
}

func NewResumeService(m *mux.Router, resume entity.ResumeUseCase) {
	handler := &resumeService{
		resumeCase: resume,
	}

	m.HandleFunc("/", handler.Create).Methods("POST")
	m.HandleFunc("/{id}/", handler.Get).Methods("GET")
	m.HandleFunc("/byUser/{userID}/", handler.GetByUser).Methods("GET")
	m.HandleFunc("/{id}/", handler.Update).Methods("PUT")
	m.HandleFunc("/{id}/", handler.Delete).Methods("DELETE")
}

func (resume *resumeService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	//Decode
	newResume := entity.Resume{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newResume)

	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	createdResume, err := resume.resumeCase.Save(ctx, newResume)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = createdResume
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)

	r.Body.Close()
}

func (resume *resumeService) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)
	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["id"]

	tempResume, err := resume.resumeCase.GetResumeByID(ctx, id)
	// tempResume, err := resume.resumeCase.GetUserByUserName(ctx, "saltosk")

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = tempResume
	} else {
		finalModel.Items = tempResume
	}

	json.NewEncoder(w).Encode(finalModel)

	r.Body.Close()
}

func (resume *resumeService) GetByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)
	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["userID"]

	tempResume, err := resume.resumeCase.GetResumeByUserID(ctx, id)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = tempResume
	} else {
		finalModel.Items = tempResume
	}

	json.NewEncoder(w).Encode(finalModel)

	r.Body.Close()
}

func (resume *resumeService) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	newResumeDetails := entity.Resume{}
	json.NewDecoder(r.Body).Decode(&newResumeDetails)
	newResumeDetails.Id = mux.Vars(r)["id"]

	updatedUser, err := resume.resumeCase.Update(ctx, newResumeDetails)

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = nil
	} else {
		finalModel.Items = updatedUser
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(finalModel)

	r.Body.Close()
}

func (resume *resumeService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["id"]

	err := resume.resumeCase.Delete(ctx, id)

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
