package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yufy/go-example/swagger/types"
)

func UpdateUser(rw http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var req types.UserRequest

	rw.Header().Set("Content-Type", "application/json")

	ID, err := strconv.Atoi(param.ByName("id"))
	if ID <= 0 || err != nil {
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(types.Response{
			Code: -1,
			Msg:  "Resource not found",
		})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(rw).Encode(types.Response{
			Code: -1,
			Msg:  "request parameter error",
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(types.Response{
		Code: 0,
		Data: types.User{
			ID:        ID,
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}
