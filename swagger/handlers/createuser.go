package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yufy/go-example/swagger/types"
)

func CreateUser(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req types.UserRequest

	rw.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(rw).Encode(types.Response{
			Code: -1,
			Msg:  "request parameter error",
		})
		return
	}

	resp := types.Response{
		Code: 0,
		Data: types.User{
			ID:        1,
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(resp)
}
