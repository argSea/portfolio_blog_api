package in_adapter

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	auth "github.com/argSea/portfolio_blog_api/argHex/utility"
	"github.com/gorilla/mux"
)

//FROM USER TO APP
type userMuxAdapter struct {
	user    in_port.UserCRUDService
	resume  in_port.UserResumeService
	project in_port.UserProjectService
	media   in_port.MediaService
}

func NewUserMuxAdapter(u in_port.UserCRUDService, r in_port.UserResumeService, p in_port.UserProjectService, m in_port.MediaService, router *mux.Router) {
	adapter := &userMuxAdapter{
		user:    u,
		resume:  r,
		project: p,
		media:   m,
	}

	//user service
	router.HandleFunc("/", adapter.Create).Methods("POST")
	router.HandleFunc("/{id}/", adapter.Get).Methods("GET")
	router.HandleFunc("/{id}/", adapter.Update).Methods("PUT")
	router.HandleFunc("/{id}/", adapter.Delete).Methods("DELETE")

	//resume service
	router.HandleFunc("/{id}/resumes/", adapter.GetResumes).Methods("GET")

	//project service
	router.HandleFunc("/{id}/projects/", adapter.GetProjects).Methods("GET")
}

func (u userMuxAdapter) Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
	}()

	// check auth
	authorized, auth_err := u.checkAuth(r)

	if nil != auth_err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: auth_err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)

		return
	}

	if !authorized {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    401,
			Message: "Unauthorized",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)

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
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp = data_objects.NewUserResponseObject{
			Status: "ok",
			Code:   200,
			UserID: new_id,
		}

		w.WriteHeader(http.StatusOK)
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
	}()

	// check for json errors in r.body
	body, body_err := ioutil.ReadAll(r.Body)

	if nil != body_err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: body_err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)

		return
	}

	// check for empty body
	if "" == string(body) {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: "Empty body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	log.Println(string(body))

	// parse body into user
	user := domain.User{}
	json.Unmarshal(body, &user)

	log.Println(user)

	id := mux.Vars(r)["id"]
	user.Id = id

	// check auth
	authorized, auth_err := u.checkAuth(r)

	if nil != auth_err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: auth_err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)

		return
	}

	if !authorized {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    401,
			Message: "Unauthorized",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)

		return
	}

	if "" != user.Password {
		// hash password
		hashed_pass, pass_err := auth.HashPassword(string(user.Password))

		if nil != pass_err {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: pass_err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)

			return
		}

		user.Password = domain.Password(hashed_pass)
	}

	for i := 0; i < len(user.Contacts); i++ {
		// check if icon is file data or url
		if "" == user.Contacts[i].Icon {
			continue
		}

		if "data:" == user.Contacts[i].Icon[:5] {
			// upload file
			mime_type := user.Contacts[i].Icon[5:strings.Index(user.Contacts[i].Icon, ";")]
			encoded_data := user.Contacts[i].Icon[strings.Index(user.Contacts[i].Icon, ",")+1:]

			decoded_data, decode_err := base64.StdEncoding.DecodeString(encoded_data)

			if nil != decode_err {
				response := data_objects.ErroredResponseObject{
					Status:  "error",
					Code:    500,
					Message: decode_err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)

				return
			}

			// get file type from mime type
			file_type := mime_type[strings.Index(mime_type, "/")+1:]

			// upload file
			upload_res, upload_err := u.media.UploadMedia(file_type, decoded_data)

			if nil != upload_err {
				response := data_objects.ErroredResponseObject{
					Status:  "error",
					Code:    500,
					Message: upload_err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)

				return
			}

			user.Contacts[i].Icon = upload_res
		}
	}

	updated_err := u.user.Update(user)

	var resp interface{}

	if nil != updated_err {
		resp = data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: updated_err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp = data_objects.ItemLessResponseObject{
			Status: "ok",
			Code:   200,
		}
		w.WriteHeader(http.StatusOK)
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
	}()

	user := domain.User{}

	id := mux.Vars(r)["id"]
	user.Id = id

	// check auth
	authorized := true //u.checkAuth(r, w, id)

	if !authorized {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    401,
			Message: "Unauthorized",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)

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
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp = data_objects.ItemLessResponseObject{
			Status: "ok",
			Code:   200,
		}
		w.WriteHeader(http.StatusOK)
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
			w.WriteHeader(http.StatusInternalServerError)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (u userMuxAdapter) GetProjects(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			w.WriteHeader(http.StatusInternalServerError)
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

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (u userMuxAdapter) checkAuth(r *http.Request) (bool, error) {
	// check auth
	validate_endpoint := "https://api.argsea.com/1/auth/validate/"

	// pass along all cookies
	cookies := r.Cookies()
	cookie_string := ""

	for i := 0; i < len(cookies); i++ {
		cookie_string += cookies[i].Name + "=" + cookies[i].Value + ";"
	}

	log.Println(cookie_string)

	req, req_err := http.NewRequest("GET", validate_endpoint, nil)

	if nil != req_err {
		return false, req_err
	}

	req.Header.Add("Cookie", cookie_string)

	val_res, val_err := http.DefaultClient.Do(req)

	if nil != val_err {
		return false, val_err
	}

	defer val_res.Body.Close()

	val_body, val_body_err := ioutil.ReadAll(val_res.Body)

	if nil != val_body_err {
		return false, val_body_err
	}

	var val_data map[string]interface{}
	json.Unmarshal(val_body, &val_data)

	if "ok" != val_data["status"] {
		return false, nil
	}

	return true, nil
}
