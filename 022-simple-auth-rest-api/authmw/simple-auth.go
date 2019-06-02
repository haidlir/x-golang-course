// Reference: https://github.com/gorilla/mux/blob/master/example_authentication_middleware_test.go

package authmw

import (
	"net/http"
)

// SimpleAuthMiddleware implement simple token auth
type SimpleAuthMiddleware struct {
	tokenUsers map[string]string
}

// Middleware function, which will be called for each request
func (amw *SimpleAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if user, found := amw.tokenUsers[token]; found {
			r.Header.Set("user", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// NewAuthMiddleware returns authentication middleware
func NewAuthMiddleware() *SimpleAuthMiddleware {
	amw := new(SimpleAuthMiddleware)
	amw.tokenUsers = map[string]string{}
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["11111111"] = "user1"
	amw.tokenUsers["cccccccc"] = "userc"
	return amw
}
