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
type userService struct {
	userCase entity.UserUsecase
}

func NewUserService(m *mux.Router, user entity.UserUsecase) {
	handler := &userService{
		userCase: user,
	}

	m.HandleFunc("/", handler.Create).Methods("POST")
	m.HandleFunc("/{id}/", handler.Get).Methods("GET")
	m.HandleFunc("/{id}/", handler.Update).Methods("PUT")
	m.HandleFunc("/{id}/", handler.Delete).Methods("DELETE")
}

func (u *userService) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	//Decode
	newUser := entity.User{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newUser)

	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	createdUser, err := u.userCase.Save(ctx, newUser)

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

	r.Body.Close()
}

func (u *userService) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	// finalModel :=
	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["id"]

	tempUser, err := u.userCase.GetUserByID(ctx, id)
	// tempUser, err := u.userCase.GetUserByUserName(ctx, "saltosk")

	if nil != err {
		finalModel.Code = 404
		finalModel.Status = "error"
		finalModel.Message = err.Error()
		finalModel.Items = tempUser
	} else {
		finalModel.Items = tempUser
	}

	json.NewEncoder(w).Encode(finalModel)

	r.Body.Close()
}

func (u *userService) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	newUserDetails := entity.User{}
	json.NewDecoder(r.Body).Decode(&newUserDetails)
	newUserDetails.Id = mux.Vars(r)["id"]

	updatedUser, err := u.userCase.Update(ctx, newUserDetails)

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

func (u *userService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	//Make model
	finalModel := &BaseResponse{
		Status: "ok",
		Code:   200,
	}

	id := mux.Vars(r)["id"]

	err := u.userCase.Delete(ctx, id)

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
