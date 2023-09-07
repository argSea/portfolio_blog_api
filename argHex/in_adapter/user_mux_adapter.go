package in_adapter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type UserMuxServices struct {
	User    in_port.UserCRUDService
	Login   in_port.UserLoginService
	Resume  in_port.UserResumeService
	Project in_port.UserProjectService
	Auth    in_port.AuthService
	Secret  []byte
}

//FROM USER TO APP
type userMuxAdapter struct {
	user    in_port.UserCRUDService
	login   in_port.UserLoginService
	resume  in_port.UserResumeService
	project in_port.UserProjectService
	auth    in_port.AuthService
	secret  []byte
}

func NewUserMuxAdapter(muxService UserMuxServices, m *mux.Router) {
	u := &userMuxAdapter{
		user:    muxService.User,
		login:   muxService.Login,
		resume:  muxService.Resume,
		project: muxService.Project,
		auth:    muxService.Auth,
		secret:  muxService.Secret,
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

	user_id, err := u.login.Login(user)

	if nil != err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	token, auth_error := u.auth.Generate(user_id)

	if nil != auth_error {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: auth_error.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// write token to http only cookie
	session, session_err := sessions.NewCookieStore(u.secret).Get(r, "auth-token")

	if nil != session_err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: session_err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	session.Values["token"] = token
	session.Save(r, w)
	log.Println("Cookie set: ", session)

	json.NewEncoder(w).Encode(data_objects.LoginResponseObject{
		Status: "ok",
		Code:   200,
		Token:  token,
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

	// check auth
	authorized := u.checkAuth(r, w, "")

	if !authorized {
		return
	}

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

	// check auth
	if !u.checkAuth(r, w, id) {
		return
	}

	// hash password
	hashed_pass, pass_err := u.login.HashPassword(string(user.Password))

	if nil != pass_err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: pass_err.Error(),
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	user.Password = domain.Password(hashed_pass)

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

	// check auth
	if !u.checkAuth(r, w, id) {
		return
	}

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

func (u userMuxAdapter) checkAuth(r *http.Request, w http.ResponseWriter, userID string) bool {
	// token := r.Header.Get("Authorization")
	session, session_err := sessions.NewCookieStore(u.secret).Get(r, "auth-token")

	if nil != session_err {
		return false
	}

	token := session.Values["token"].(string)

	// check if user is authorized
	authorized := u.auth.IsAuthorized(userID, token, in_port.PERM_USER, in_port.PERM_ADMIN)

	if !authorized {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    401,
			Message: "Unauthorized",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return false
	}

	return true
}
