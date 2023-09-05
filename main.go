package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/argSea/portfolio_blog_api/argHex/data_objects"
	"github.com/argSea/portfolio_blog_api/argHex/in_adapter"
	"github.com/argSea/portfolio_blog_api/argHex/out_adapter"
	"github.com/argSea/portfolio_blog_api/argHex/service"
	"github.com/argSea/portfolio_blog_api/argHex/stores"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()

	if nil != err {
		panic(err)
	}

	//Possibly add debugging?
}

func main() {
	//logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log_file, log_file_err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0775)
	if nil != log_file_err {
		log.Fatal(log_file_err)
	}

	//signal to kill and print final info
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("Shutting down argSea API")
		fmt.Println("Shutting down argSea API")
		os.Exit(0)
	}()

	log.SetOutput(log_file)

	//mux
	router := mux.NewRouter()
	router.Use(baseMiddleWare)
	router.StrictSlash(true)

	//Cache credentials
	mHost := viper.GetString("mongo.host") + ":" + viper.GetString("mongo.port")
	mUser := viper.GetString("mongo.user")
	mPass := viper.GetString("mongo.pass")
	mDB := viper.GetString("mongo.dbName")
	jSecret := []byte(viper.GetString("jwt.secret"))

	mongo_db, mongo_err := stores.NewMongoStore(mUser, mPass, mHost, mDB)

	defer mongo_db.Client.Disconnect(context.Background())

	if nil != mongo_err {
		fmt.Fprintf(os.Stderr, "error: %v\n", mongo_err)
		log.Fatal(mongo_err)
		os.Exit(1)
	}

	user_table := "users"
	projectTable := "projects"
	resumeTable := "resume"

	//user
	userRouter := router.PathPrefix("/1/user/").Subrouter()
	projRouter := router.PathPrefix("/1/project/").Subrouter()
	resumeRouter := router.PathPrefix("/1/resume/").Subrouter()

	//resume
	log.Println("Initializing resume")
	// resumeDrivenAdapter := out_adapter.NewResumeFakeOutAdapter()
	resumeMordor := stores.NewMordor(mongo_db.DB.Collection(resumeTable), context.Background())
	resumeMongoAdapter := out_adapter.NewResumeMongoAdapter(resumeMordor)
	resumeService := service.NewResumeCRUDService(resumeMongoAdapter)
	in_adapter.NewResumeMuxAdapter(resumeService, resumeRouter)

	//project
	log.Println("Initializing project")
	// projectDrivenAdapter := out_adapter.NewProjectFakeOutAdapter()
	projectMordor := stores.NewMordor(mongo_db.DB.Collection(projectTable), context.Background())
	projectMongoAdapter := out_adapter.NewProjectMongoAdapter(projectMordor)
	projectService := service.NewProjectCRUDService(projectMongoAdapter)
	in_adapter.NewProjectMuxAdapter(projectService, projRouter)

	//User
	log.Println("Initializing user")
	// userDrivenAdapter := out_adapter.NewUserFakeOutAdapter()
	userMordor := stores.NewMordor(mongo_db.DB.Collection(user_table), context.Background())
	userMongoAdapter := out_adapter.NewUserMongoAdapter(userMordor)
	userService := service.NewUserCRUDService(userMongoAdapter)
	userResumeService := service.NewUserResumeService(resumeMongoAdapter)
	userProjectService := service.NewUserProjectService(projectMongoAdapter)
	userAuthService := service.NewUserAuthService(userMongoAdapter)
	in_adapter.NewUserMuxAdapter(userService, userAuthService, userResumeService, userProjectService, jSecret, userRouter)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8181",
		Handler:      router,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func baseMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		fmt.Println(r.URL)
		fmt.Println(r.Method)

		exemptedPaths := []string{
			"/1/user/login/",
			"/1/user/signup/",
		}

		// if not POST, PUT or DELETE, just continue
		if r.Method != "POST" && r.Method != "PUT" && r.Method != "DELETE" {
			next.ServeHTTP(w, r)
			return
		}

		// if path is exempted, just continue
		for _, path := range exemptedPaths {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		// check if jwt is present
		token := r.Header.Get("Authorization")

		if token == "" {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    401,
				Message: "Unauthorized",
			}
			json.NewEncoder(w).Encode(response)

			return
		}

		// parse jwt
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.secret")), nil
		})

		if err != nil {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    401,
				Message: "Unauthorized",
			}
			json.NewEncoder(w).Encode(response)

			return
		}

		// check if jwt is expired
		exp := claims["exp"].(float64)
		expTime := time.Unix(int64(exp), 0)

		if time.Now().After(expTime) {
			response := data_objects.ErroredResponseObject{
				Status:  "error",
				Code:    401,
				Message: "Unauthorized",
			}
			json.NewEncoder(w).Encode(response)

			return
		}

		// get userID from jwt
		userID := claims["userID"].(string)

		// check if userID is present in the URL or in the body
		if r.Method == "POST" || r.Method == "PUT" {
			var body map[string]interface{}
			json.NewDecoder(r.Body).Decode(&body)

			if body["userID"] != userID {
				response := data_objects.ErroredResponseObject{
					Status:  "error",
					Code:    401,
					Message: "Unauthorized",
				}
				json.NewEncoder(w).Encode(response)

				return
			}
		} else {
			if mux.Vars(r)["id"] != userID {
				response := data_objects.ErroredResponseObject{
					Status:  "error",
					Code:    401,
					Message: "Unauthorized",
				}
				json.NewEncoder(w).Encode(response)

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func contains(exemptedPaths []string, s string) {
	panic("unimplemented")
}
