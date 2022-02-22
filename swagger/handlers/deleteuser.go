package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/yufy/go-example/swagger/types"
)

func DeleteUser(rw http.ResponseWriter, r *http.Request, param httprouter.Params) {
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
	})
}
