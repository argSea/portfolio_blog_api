package in_adapter

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type authMuxAdapter struct {
	authService  in_port.AuthService
	loginService in_port.UserLoginService
	secret       []byte
}

func NewAuthMuxAdapter(a in_port.AuthService, l in_port.UserLoginService, s []byte, r *mux.Router) {
	adapter := authMuxAdapter{
		authService:  a,
		loginService: l,
		secret:       s,
	}

	//user auth service
	r.HandleFunc("/login/", adapter.Login).Methods("POST")
	r.HandleFunc("/validate/", adapter.Validate).Methods("GET")

}

func (a authMuxAdapter) Login(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			// set code 500
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
	}()

	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	user, err := a.loginService.Login(user)

	if nil != err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    400,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	token, token_error := a.setSession(user, w, r)

	if nil != token_error {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: token_error.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data_objects.LoginResponseObject{
		Status:   "ok",
		Code:     200,
		UserName: user.UserName,
		UserID:   user.Id,
		Token:    token,
	})
}

func (a authMuxAdapter) Validate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    500,
				Message: err,
			}
			// set code 500
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
	}()

	// check if auth-token cookie exists
	session, session_err := sessions.NewCookieStore(a.secret).Get(r, "auth-token")

	if nil != session_err {
		log.Println("Error getting session: ", session_err)
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: session_err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("Session data: ", session)

	token := session.Values["token"].(string)

	// check auth
	v_response, v_err := a.authService.Validate(token)

	if nil != v_err {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    500,
			Message: v_err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	if !v_response.Valid {
		response := data_objects.ErroredResponseObject{
			Status:  "error",
			Code:    401,
			Message: "Unauthorized",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println("User is authorized! " + v_response.UserID)

	userID := v_response.UserID
	// role := v_response.Role

	// get user
	user := a.getUserDetails(userID)

	// return user details
	response := data_objects.UserResponseObject{
		Status: "ok",
		Code:   200,
	}

	response.Users = append(response.Users, user)

	// todo: add role to response

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (a authMuxAdapter) setSession(user domain.User, w http.ResponseWriter, r *http.Request) (string, error) {
	expires := time.Now().Add(time.Hour * 24)
	roles := []string{"user"}
	token, auth_error := a.authService.Generate(user.Id, expires, roles)

	if nil != auth_error {
		return "", auth_error
	}

	sess_options := &sessions.Options{
		// Domain:   "argsea.com",
		Path:     "/",
		MaxAge:   24 * 60 * 60,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}

	session, session_err := sessions.NewCookieStore(a.secret).Get(r, "auth-token")
	session.Options = sess_options

	if nil != session_err {
		return "", session_err
	}

	session.Values["token"] = token
	session.Values["iat"] = time.Now().Unix()
	session.Save(r, w)
	log.Println("Cookie set: ", session)

	return token, nil
}

func (a authMuxAdapter) checkAuth(r *http.Request, w http.ResponseWriter, userID string) bool {
	// token := r.Header.Get("Authorization")
	session, session_err := sessions.NewCookieStore(a.secret).Get(r, "auth-token")

	if nil != session_err {
		log.Println("Error getting session: ", session_err)
		return false
	}

	token := session.Values["token"].(string)

	// check if user is authorized
	authorized := a.authService.IsAuthorized(userID, token, in_port.PERM_USER, in_port.PERM_ADMIN)

	if !authorized {
		log.Println("User not authorized! " + userID)
		return false
	}

	user := a.getUserDetails(userID)

	// if authorized, refresh token
	a.setSession(user, w, r)

	return true
}

func (a authMuxAdapter) getUserDetails(userID string) domain.User {
	user_call := "https://api.argsea.com/1/user/" + userID + "/"
	user_resp, user_err := http.Get(user_call)

	if nil != user_err {
		log.Println("Error getting user: ", user_err)
		return domain.User{}
	}

	var user_response data_objects.UserResponseObject
	json.NewDecoder(user_resp.Body).Decode(&user_response)

	if 0 == len(user_response.Users) {
		log.Println("User not found! " + userID)
		return domain.User{}
	}

	user := user_response.Users[0].(domain.User)

	return user
}
