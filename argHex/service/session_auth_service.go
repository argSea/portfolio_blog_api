package service

import (
	"encoding/json"
	"time"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/argSea/portfolio_blog_api/argHex/out_port"
	"github.com/google/uuid"
)

type sessionAuthData struct {
	Expires time.Time `json:"expires"`
	Roles   []string  `json:"roles"`
	Id      string    `json:"id"`
}

type sessionAuthService struct {
	repo   out_port.AuthRepo
	secret []byte
}

func NewSessionAuthService(repo out_port.AuthRepo, secret []byte) in_port.AuthService {
	return sessionAuthService{
		repo:   repo,
		secret: secret,
	}
}

// NewAuth
func (s sessionAuthService) Generate(id string) (string, error) {
	token := uuid.New().String()
	expires := time.Now().Add(time.Hour * 24 * 7) // 7 days
	roles := []string{"user"}

	data := sessionAuthData{
		Expires: expires,
		Roles:   roles,
		Id:      id,
	}

	// json_data, json_err := json.Marshal(data)

	// if nil != json_err {
	// 	return "", json_err
	// }

	// store token, id, and expiration in json format
	err := s.repo.Store(token, data)

	if nil != err {
		return "", err
	}

	return token, nil
}

// Validate
func (s sessionAuthService) Validate(token string) (data_objects.AuthValidationResponseObject, error) {
	// get token from redis
	data := s.repo.Get(token)

	// check if token exists
	if "" == data {
		return data_objects.AuthValidationResponseObject{}, nil
	}

	// unmarshal data
	var authData sessionAuthData
	json_err := json.Unmarshal([]byte(data), &authData)

	if nil != json_err {
		return data_objects.AuthValidationResponseObject{}, json_err
	}

	// check if token is expired
	if time.Now().After(authData.Expires) {
		return data_objects.AuthValidationResponseObject{}, nil
	}

	// return auth data
	return data_objects.AuthValidationResponseObject{
		Valid:  true,
		UserID: authData.Id,
		Role:   authData.Roles[0],
	}, nil
}

// check if user is authorized
func (s sessionAuthService) IsAuthorized(id string, token string, roles ...string) bool {
	// check if token is valid
	authData, err := s.Validate(token)

	if nil != err {
		return false
	}

	// check if user id matches
	if id != authData.UserID {
		return false
	}

	// check if user has required role
	for _, role := range roles {
		if authData.Role == role {
			return true
		}
	}

	return false
}
