// Package classification go-example API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//	   Host: localhost
//	   BasePath: /v1
//     Version: 1.0.0
//     Contact: yufy<1504960756@qq.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import "github.com/yufy/go-example/swagger/types"

// swagger:parameters createUser
type CreateUserRequest struct {
	// in: body
	Body struct {
		types.UserRequest
	}
}

// swagger:parameters updateUser
type UpdateUserRequest struct {
	// in: path
	// minimum: 1
	ID int `json:"id"`
	// in: body
	Body struct {
		types.UserRequest
	}
}

// swagger:parameters user deleteUser
type IDRequest struct {
	// in: path
	// minimum: 1
	ID int `json:"id"`
}

// swagger:parameters users
type UsersRequest struct {
	// in: query
	// minimum: 1
	Page int `json:"page"`
	// in: query
	// minimum: 1
	PageSize int `json:"page_size"`
	// in: query
	Name string `json:"name"`
	// in: query
	Email string `json:"email"`
}

// swagger:response errorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		types.Response
	}
}

// swagger:response userResponse
type UserResponse struct {
	// in: body
	Body struct {
		types.Response
		Data types.User `json:"data"`
	}
}

// swagger:response usersResponse
type UsersResponse struct {
	// in: body
	Body struct {
		types.Response
		Data []types.User `json:"data"`
	}
}

// swagger:route POST /users 用户 createUser
//
// 添加用户
//
// 添加新用户
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- appliation/json
//
//		Scheme: http
//
//		Responses:
//			200: userResponse
//			400: errorResponse
//			500: errorResponse

// swagger:route GET /users 用户 users
//
// 用户列表
//
// 通过指定参数查询用户
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- appliation/json
//
//		Scheme: http
//
//		Responses:
//			200: usersResponse

// swagger:route GET /users/:id 用户 user
//
// 用户详情
//
// 获取指定ID用户
//
//		Consumes:
//		- urlform
//
//		Produces:
//		- appliation/json
//
//		Scheme: http
//
//		Responses:
//			200: usersResponse
//			404: errorResponse

// swagger:route DELETE /users/:id 用户 deleteUser
//
// 删除用户
//
// 删除指定ID用户
//
//		Consumes:
//		- urlform
//
//		Produces:
//		- appliation/json
//
//		Scheme: http
//
//		Responses:
//			200: usersResponse
//			404: errorResponse
//			500: errorResponse

// swagger:route PUT /users/:id 用户 updateUser
//
// 更新用户
//
// 更新指定ID用户
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- appliation/json
//
//		Scheme: http
//
//		Responses:
//			200: usersResponse
//			404: errorResponse
//			500: errorResponse
