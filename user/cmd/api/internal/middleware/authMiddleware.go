package middleware

import (
	"encoding/json"
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
		userId, err := r.Context().Value("userId").(json.Number).Int64()
		if err != nil {
			fmt.Println("============>>>>>>> ", err)
		}
		println("userId: ========》》》", userId)

		next(w, r)
	}
}
