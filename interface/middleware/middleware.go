package middleware

import (
	"context"
	"message_api/domain/output"
	"message_api/lib/jwt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Auth .
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-User-Token")
		currentUserID := jwt.FetchUserID(token)
		if currentUserID == 0 {
			output.Error{ErrorMessage: "authentication error"}.WriteResponseJSON(w)
			return
		}

		ctx := context.WithValue(r.Context(), "current_user_id", uint(currentUserID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GroupCtx .
func GroupCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "group_id")
		id, _ := strconv.Atoi(param)

		ctx := context.WithValue(r.Context(), "group_id", uint(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserCtx .
func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "user_id")
		id, _ := strconv.Atoi(param)

		ctx := context.WithValue(r.Context(), "user_id", uint(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GroupMessageCtx .
func GroupMessageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "group_message_id")
		id, _ := strconv.Atoi(param)

		ctx := context.WithValue(r.Context(), "group_message_id", uint(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DirectMessageCtx .
func DirectMessageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "direct_message_id")
		id, _ := strconv.Atoi(param)

		ctx := context.WithValue(r.Context(), "direct_message_id", uint(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
