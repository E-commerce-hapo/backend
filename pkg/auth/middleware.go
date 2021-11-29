package auth

import (
	"net/http"
	"strings"

	"github.com/kiem-toan/infrastructure/authorize/auth"

	"github.com/kiem-toan/infrastructure/errorx"

	"github.com/gin-gonic/gin"
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

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//err := TokenValid(c.Request)
		//if err != nil {
		//	c.JSON(http.StatusUnauthorized, err)
		//	c.Abort()
		//	return
		//}
		//claims, err := ExtractTokenMetaData(c.Request)
		//_err, _ := err.(*errorx.Errorx)
		//if err != nil {
		//	pp.Println(claims)
		//	c.JSON(_err.StatusCode, _err)
		//	c.Abort()
		//	return
		//}
		// TODO: Get user info and pass into SessionInfo
		rolesAfterGetFromDB := auth.Roles{"admin", "accountant"}
		urlFromRequest := "GetProduct"
		methodFromRequest := "post"
		action := strings.Join([]string{urlFromRequest, methodFromRequest}, ":")
		e := auth.New()

		if e.Check(rolesAfterGetFromDB, action) {
			c.JSON(http.StatusUnauthorized, errorx.New(http.StatusUnauthorized, nil, "Không tìm thấy hoặc cần quyền truy cập.."))
			return
		}
		var sessionInfo *SessionInfo
		c.Set("SS", sessionInfo)
		c.Next()
	}
}
