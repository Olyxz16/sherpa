package middleware

import (
	"context"
	"net/http"

	"github.com/Olyxz16/sherpa/domain/model"
	"github.com/Olyxz16/sherpa/infrastructure/jwt"
	"github.com/Olyxz16/sherpa/infrastructure/persistence/repository"
)

func Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		uid, err := jwt.ParseSessionCookie(cookie)
		if err != nil {
			cookie = jwt.CleanCookie()
			http.SetCookie(w, cookie)
			next.ServeHTTP(w, r)
			return
		}
		
		userRepo := repository.NewUserRepository()
		user, err := userRepo.FindFromID(uid, r.Context())
		if err != nil {
			cookie = jwt.CleanCookie()
			http.SetCookie(w, cookie)
			next.ServeHTTP(w, r)
			return
		}

		refreshCookie := jwt.RefreshCookie(cookie)
		http.SetCookie(w, refreshCookie)
		context := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(context))
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*model.User)
		if !ok {
			http.Redirect(w, r, "/", 302)
			return
		}

		if user.IsNew() {
			http.Redirect(w, r, "/welcome", 302)
		}

		next.ServeHTTP(w, r)
	})
}
