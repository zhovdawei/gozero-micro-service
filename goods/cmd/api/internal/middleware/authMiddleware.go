package middleware

import (
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need

		fmt.Println("==============================================================")
		fmt.Println(" auth middleWare execute ......")
		fmt.Println("==============================================================")

		next(w, r)
	}
}
