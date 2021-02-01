package middleware

import (
	"context"
	"fmt"
	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

var wCtxKey = &wContextKey{"w"}

type wContextKey struct {
	name string
}


// Middleware decodes the share session cookie and packs the session into context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header, err := r.Cookie("access_token")

		// put http.ResponseWriter in context for later use
		wCtx := context.WithValue(r.Context(), wCtxKey, &w)
		r = r.WithContext(wCtx)

		// Allow unauthenticated users in
		if err != nil || header == nil {
			fmt.Println("[INFO] Unauthenticated user")
			next.ServeHTTP(w, r)
			return
		}

		//validate jwt token
		tokenStr := header
		id, err := providers.ParseToken(tokenStr.Value)

		if err != nil {
			fmt.Println("[INFO] Invalid cookie")
			http.Error(w, "Invalid cookie", http.StatusForbidden)
			return
		}

		// create user and check if user exists in db
		userId := int(id)
		user := model.User{ID: userId}

		// put it in context
		ctx := context.WithValue(r.Context(), userCtxKey, &user)

		fmt.Println("[INFO] Authenticated user")

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}


// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

func WForContext(ctx context.Context) *http.ResponseWriter {
	raw, _ := ctx.Value(wCtxKey).(*http.ResponseWriter)
	return raw
}

