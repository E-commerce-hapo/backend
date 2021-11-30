package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/kiem-toan/pkg/httpx"

	"github.com/kiem-toan/pkg/authorize/auth"

	"github.com/kiem-toan/pkg/errorx"
)

func CORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")
		switch {
		case
			origin == "ionic://localhost",
			origin == "capacitor://localhost",
			origin == "http://localhost",
			origin == "http://localhost:8080",
			origin == "http://localhost:8100",
			strings.HasSuffix(origin, ".d.etop.vn"):
			w.Header().Set("Access-Control-Allow-Origin", origin)

		//case cmenv.IsSandBox(), cmenv.IsDevOrStag():
		//	w.Header().Set("Access-Control-Allow-Origin", "*")
		//
		//case cmenv.IsProd():

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

func TokenAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			httpx.WriteError(ctx, w, errorx.Errorf(http.StatusUnauthorized, nil, "Không tìm thấy hoặc cần quyền truy cập."))
			return
		}
		sessionInfo := &SessionInfo{}
		if claims != nil {
			sessionInfo.UserID = claims.UserID
		}
		ctx = context.WithValue(ctx, "ss", sessionInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
