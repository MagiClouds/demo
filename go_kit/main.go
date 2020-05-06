package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"net/http"
	"start/go_kit/services"
)

func main()  {
	user := services.UserService{}
	endp := services.GenUserEndpoint(user)
	serverHandler := httptransport.NewServer(endp, services.DecodeUserRequest, services.EncodeUserResponse)

	r := mymux.NewRouter()
	r.Methods("GET").Path(`/user/{uid:\d+}`).Handler(serverHandler)

	http.ListenAndServe(":8080", r)
}
