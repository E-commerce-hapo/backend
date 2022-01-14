package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/E-commerce-hapo/backend/core/config"
	"github.com/E-commerce-hapo/backend/pkg/env"

	"github.com/E-commerce-hapo/backend/pkg/httpx"

	"github.com/E-commerce-hapo/backend/pkg/authorize/auth"

	"github.com/E-commerce-hapo/backend/pkg/errorx"
)

func CORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")
		switch {
		case config.GetAppConfig().Env == env.EnvDev,
			origin == "ionic://localhost",
			origin == "capacitor://localhost",
			origin == "http://localhost",
			origin == "http://localhost:8080",
			origin == "http://localhost:8100":
			w.Header().Set("Access-Control-Allow-Origin", origin)

		case config.GetAppConfig().Env == env.EnvStaging:
			w.Header().Set("Access-Control-Allow-Origin", "*")

		case config.GetAppConfig().Env == env.EnvProd:

		default:
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Add("Access-Control-Max-Age", "86400")
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	}
}

func TokenAuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := ValidateToken(r)
		if err != nil {
			httpx.WriteError(ctx, w, err)
			return
		}
		claims, err := GetCustomClaimsFromRequest(r)
		if err != nil {
			httpx.WriteError(ctx, w, err)
			return
		}
		// TODO: Get user info and pass into SessionInfo
		rolesAfterGetFromDB := auth.Roles{"admin", "shipper"}
		urlFromRequest := r.RequestURI
		methodFromRequest := r.Method
		action := strings.Join([]string{urlFromRequest, methodFromRequest}, ":")
		e := auth.New()

		if !e.Check(rolesAfterGetFromDB, action) {
			httpx.WriteError(ctx, w, errorx.ErrAuthFailure(errors.New(fmt.Sprintf("Không tìm thấy hoặc cần quyền truy cập."))))
			return
		}
		sessionInfo := &SessionInfo{}
		if claims != nil {
			sessionInfo.UserID = claims.UserID
		}
		ctx = context.WithValue(ctx, "ss", sessionInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
