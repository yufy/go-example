package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yufy/go-example/swagger/types"
)

func User(rw http.ResponseWriter, r *http.Request, param httprouter.Params) {
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

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(types.Response{
		Code: 0,
		Data: types.User{
			ID:        ID,
			Name:      "test",
			Email:     "test@test.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}
