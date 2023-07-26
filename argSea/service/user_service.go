package service

import (
	"encoding/json"
	"net/http"

	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/gorilla/mux"
)

//Handler
type userService struct {
	userCase core.UserUsecase
}

func NewUserService(m *mux.Router, user core.UserUsecase) {
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

	// //Decode
	// newUser := entity.User{}
	// decoder := json.NewDecoder(r.Body)
	// decoder.Decode(&newUser)

	// //Make model
	// finalModel := &BaseResponse{
	// 	Status: "ok",
	// 	Code:   200,
	// }

	// createdUser, err := u.userCase.Save(newUser)

	// if nil != err {
	// 	finalModel.Code = 404
	// 	finalModel.Status = "error"
	// 	finalModel.Message = err.Error()
	// 	finalModel.Items = nil
	// } else {
	// 	finalModel.Items = createdUser
	// }

	// encoder := json.NewEncoder(w)
	// encoder.SetIndent("", "    ")
	// encoder.Encode(finalModel)

	r.Body.Close()
}

func (u *userService) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	view, err := u.userCase.GetUserByID(id)

	if nil != err {
		finalModel := &BaseResponse{
			Status:  "ok",
			Code:    400,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(finalModel)
	} else {
		json.NewEncoder(w).Encode(view)
	}

	r.Body.Close()
}

func (u *userService) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// finalModel := &BaseResponse{
	// 	Status: "ok",
	// 	Code:   200,
	// }

	// newUserDetails := entity.User{}
	// json.NewDecoder(r.Body).Decode(&newUserDetails)
	// newUserDetails.Id = mux.Vars(r)["id"]

	// updatedUser, err := u.userCase.Update(newUserDetails)

	// if nil != err {
	// 	finalModel.Code = 404
	// 	finalModel.Status = "error"
	// 	finalModel.Message = err.Error()
	// 	finalModel.Items = nil
	// } else {
	// 	finalModel.Items = updatedUser
	// }

	// encoder := json.NewEncoder(w)
	// encoder.SetIndent("", "    ")
	// encoder.Encode(finalModel)

	r.Body.Close()
}

func (u *userService) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// //Make model
	// finalModel := &BaseResponse{
	// 	Status: "ok",
	// 	Code:   200,
	// }

	// id := mux.Vars(r)["id"]

	// err := u.userCase.Delete(id)

	// if nil != err {
	// 	finalModel.Code = 404
	// 	finalModel.Status = "error"
	// 	finalModel.Message = err.Error()
	// 	finalModel.Items = nil
	// } else {
	// 	finalModel.Message = "Deleted"
	// 	finalModel.Items = nil
	// }

	// encoder := json.NewEncoder(w)
	// encoder.SetIndent("", "    ")
	// encoder.Encode(finalModel)

	r.Body.Close()
}
