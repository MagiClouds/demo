package services

import (
	"context"
	"encoding/json"
	"fmt"
	mymux "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	// http://localhost:8080/user/{uid}
	vars := mymux.Vars(r)
	if v, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(v)
		return UserRequest{Uid:uid}, nil
	}

	return nil, fmt.Errorf("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error  {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
