// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type DefaultJsonMiddleware struct {
}

func NewDefaultJsonMiddleware() *DefaultJsonMiddleware {
	return &DefaultJsonMiddleware{}
}

func (m *DefaultJsonMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			httpx.Error(w, fmt.Errorf("Content-Type must be application/json"))
			return
		}
		next(w, r)
	}
}
