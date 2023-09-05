package in_adapter

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

//FROM USER TO APP
type userMuxAdapter struct {
	user      in_port.UserCRUDService
	auth      in_port.UserAuthService
	resume    in_port.UserResumeService
	project   in_port.UserProjectService
	jtwSecret []byte
}

func NewUserMuxAdapter(
	user in_port.UserCRUDService,
	auth in_port.UserAuthService,
	resume in_port.UserResumeService,
	project in_port.UserProjectService,
	jtwSecret []byte,
	m *mux.Router) {

	u := &userMuxAdapter{
		user:      user,
		auth:      auth,
		resume:    resume,
		project:   project,
		jtwSecret: jtwSecret,
	}

	//user service
	m.HandleFunc("/", u.Create).Methods("POST")
	m.HandleFunc("/{id}/", u.Get).Methods("GET")
	m.HandleFunc("/{id}/", u.Update).Methods("PUT")
	m.HandleFunc("/{id}/", u.Delete).Methods("DELETE")

	//user auth service
	m.HandleFunc("/login/", u.Login).Methods("POST")

	//resume service
	m.HandleFunc("/{id}/resumes/", u.GetResumes).Methods("GET")

	//project service
	m.HandleFunc("/{id}/projects/", u.GetProjects).Methods("GET")
}

func (u userMuxAdapter) Login(w http.ResponseWriter, r *http.Request) {
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

	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	user_id, err := u.auth.Login(user)

	if nil != err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// create jwt token
	key := u.jtwSecret
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user_id
	// get time since epoch
	time_now := time.Now()
	// add 30 days
	time_now.AddDate(0, 0, 30)
	claims["exp"] = time_now.Unix()
	//q: any other claims we want to add?
	claims["role"] = "user"

	// later we can add a claim to make my user an admin
	// claims["role"] = "admin"
	tokenString, _ := token.SignedString(key)

	json.NewEncoder(w).Encode(data_objects.LoginResponseObject{
		Status: "ok",
		Code:   200,
		Token:  tokenString,
	})
}

func (u userMuxAdapter) Create(w http.ResponseWriter, r *http.Request) {
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

	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	new_id, err := u.user.Create(user)
	var resp interface{}

	if nil != err {
		resp = data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: err.Error(),
		}
	} else {
		resp = data_objects.NewUserResponseObject{
			Status: "ok",
			Code:   200,
			UserID: new_id,
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func (u userMuxAdapter) Get(w http.ResponseWriter, r *http.Request) {
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
	user_data := u.user.Read(id)

	response := data_objects.UserResponseObject{
		Status: "ok",
		Code:   200,
	}

	response.Users = append(response.Users, user_data)

	json.NewEncoder(w).Encode(response)
}

func (u userMuxAdapter) Update(w http.ResponseWriter, r *http.Request) {
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

	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	id := mux.Vars(r)["id"]
	user.Id = id
	updated_err := u.user.Update(user)

	var resp interface{}

	if nil != updated_err {
		resp = data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: updated_err.Error(),
		}
	} else {
		resp = data_objects.ItemLessResponseObject{
			Status: "ok",
			Code:   200,
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func (u userMuxAdapter) Delete(w http.ResponseWriter, r *http.Request) {
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

	user := domain.User{}

	id := mux.Vars(r)["id"]
	user.Id = id
	deleted_err := u.user.Delete(user)

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

func (u userMuxAdapter) GetResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

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

	userID := mux.Vars(r)["id"]
	user_resumes, count := u.resume.GetResumes(userID)

	response := data_objects.ResumeResponseObject{
		Status: "ok",
		Code:   200,
		Count:  count,
	}

	for i := 0; i < len(user_resumes); i++ {
		response.Resumes = append(response.Resumes, user_resumes[i])
	}

	json.NewEncoder(w).Encode(response)
}

func (u userMuxAdapter) GetProjects(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", "application/json")

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

	userID := mux.Vars(r)["id"]
	user_projects, count := u.project.GetProjects(userID)

	response := data_objects.ProjectResponseObject{
		Status: "ok",
		Code:   200,
		Count:  count,
	}

	for i := 0; i < len(user_projects); i++ {
		response.Projects = append(response.Projects, user_projects[i])
	}

	json.NewEncoder(w).Encode(response)
}
