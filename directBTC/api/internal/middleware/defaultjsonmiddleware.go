// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"
)

type DefaultJsonMiddleware struct {
}

func NewDefaultJsonMiddleware() *DefaultJsonMiddleware {
	return &DefaultJsonMiddleware{}
}

func (m *DefaultJsonMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" && r.Method == http.MethodPost {
			r.Header.Set("Content-Type", "application/json")
		}
		next(w, r)
	}
}
