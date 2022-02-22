package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yufy/go-example/swagger/types"
)

func Users(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	if pageSize == 0 {
		pageSize = 10
	}
	name := q.Get("name")
	email := q.Get("email")

	{
		// query users
		fmt.Printf("\nQuery parameters: \n")
		fmt.Printf("\tpage: %d\n", page)
		fmt.Printf("\tpageSize: %d\n", pageSize)
		fmt.Printf("\tname: %s\n", name)
		fmt.Printf("\temail: %s\n", email)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(types.Response{
		Code: 0,
		Data: types.Users{
			Count: 2,
			List: []types.User{
				{ID: 1, Name: "user1", Email: "email1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{ID: 1, Name: "user1", Email: "email1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		},
	})
}
