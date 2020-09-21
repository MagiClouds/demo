package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"start/go_kit/services"
	"start/go_kit/util"
	"syscall"
)

func main() {
	user := services.UserService{}
	endp := services.GenUserEndpoint(user)
	serverHandler := httptransport.NewServer(endp, services.DecodeUserRequest, services.EncodeUserResponse)

	r := mymux.NewRouter()
	r.Methods("GET").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	r.Methods("Get").Path("/health").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})

	errChan := make(chan error)

	go func() {
		util.RegisterUserService()
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Println(err.Error())
			errChan <- err
		}
	}()

	go func() {
		signC := make(chan os.Signal)
		signal.Notify(signC, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-signC)
	}()

	getErr := <-errChan
	log.Println(getErr)

	util.DeregisterUserService()
}
