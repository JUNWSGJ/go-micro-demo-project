// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type UserHandler interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateReply, error)

	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)

	RegisterOrLogin(context.Context, *RegisterOrLoginRequest) (*GetUserReply, error)
}

func NewUserHandler(srv UserHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/user.service.v1.User/GetUser", func(w http.ResponseWriter, r *http.Request) {
		var in GetUserRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetUserReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/user.service.v1.User/RegisterOrLogin", func(w http.ResponseWriter, r *http.Request) {
		var in RegisterOrLoginRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterOrLogin(ctx, req.(*RegisterOrLoginRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetUserReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/user.service.v1.User/Authenticate", func(w http.ResponseWriter, r *http.Request) {
		var in AuthenticateRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Authenticate(ctx, req.(*AuthenticateRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		out, err := next(r.Context(), &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*AuthenticateReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	return r
}
