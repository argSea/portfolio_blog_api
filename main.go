package main

import (
	"log"
	"net/http"
	"time"

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
	router := mux.NewRouter()
	router.Use(corsMiddleWare)

	//Cache credentials
	// mHost := viper.GetString("mongo.host") + ":" + viper.GetString("mongo.port")
	// mUser := viper.GetString("mongo.user")
	// mPass := viper.GetString("mongo.pass")
	// mDB := viper.GetString("mongo.dbName")

	// userTable := "users"
	// projectTable := "projects"
	// resumeTable := "resume"

	//User
	// userRepo := repo.NewUserRepo(argStore.NewMordor(mHost, mUser, mPass, mDB, userTable))
	// userRepo := repo.NewUserRepo(argStore.NewTestStore())
	// userPres := presenter.NewUserPresenter()
	// userCase := usecase.NewAPIUserCase(userRepo, userPres)

	//Project
	// projRepo := repo.NewProjectRepo(argStore.NewMordor(mHost, mUser, mPass, mDB, projectTable))
	// projPres := presenter.NewProjectPresenter()
	// projCase := usecase.NewAPIProjectCase(projRepo, projPres)

	//Resume
	// resumeRepo := repo.NewResumeRepo(argStore.NewMordor(mHost, mUser, mPass, mDB, resumeTable))
	// resumePres := presenter.NewResumePresenter()
	// resumeCase := usecase.NewAPIResumeCase(resumeRepo, resumePres)

	//user
	// userRouter := router.PathPrefix("/api/1/user/").Subrouter()
	// service.NewUserService(userRouter, userCase)

	// //Project
	// projRouter := router.PathPrefix("/api/1/project/").Subrouter()
	// service.NewProjectService(projRouter, projCase)

	// //Resume
	// resumeRouter := router.PathPrefix("/api/1/resume/").Subrouter()
	// service.NewResumeService(resumeRouter, resumeCase)

	//User
	// userDrivenAdapter := userAdapters.NewUserFakeOutAdapter()
	// userService := userService.NewUserCRUDService(userDrivenAdapter)
	// userAdapters.NewUserMuxAdapter(userService, userRouter)

	// resumeDrivenAdapter := resumeAdapters.NewResumeFakeOutAdapter()

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

func corsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
