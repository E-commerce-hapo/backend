package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kiem-toan/infrastructure/errorx"

	"github.com/dgrijalva/jwt-go"

	"github.com/kiem-toan/infrastructure/idx"
)

const (
	AuthorizationHeader         = "Authorization"
	AuthorizationScheme         = "Bearer"
	TokenExpiresDuration        = 60
	RefreshTokenExpiresDuration = 24 * 7
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type SessionInfo struct {
	UserID idx.ID
}

type JWTCustomClaims struct {
	UserID idx.ID
	jwt.StandardClaims
}

func GenerateToken(userID idx.ID) (*TokenDetails, error) {
	td := &TokenDetails{}

	var err error
	//Creating Access Token
	td.AtExpires = time.Now().Add(time.Minute * TokenExpiresDuration).Unix()
	atClaims := &JWTCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Minute * RefreshTokenExpiresDuration).Unix()
	rtClaims := &JWTCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func extractToken(r *http.Request) (string, error) {
	s := r.Header.Get(AuthorizationHeader)
	if s == "" {
		return "", errorx.New(http.StatusUnauthorized, nil, "Missing authorization string")
	}
	splits := strings.SplitN(s, " ", 2)
	if len(splits) < 2 {
		return "", errorx.New(http.StatusUnauthorized, nil, "Bad authorization string")
	}
	if splits[0] != AuthorizationScheme && splits[0] != AuthorizationScheme {
		return "", errorx.New(http.StatusUnauthorized, nil, "Request unauthenticated with "+AuthorizationScheme)
	}
	return splits[1], nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, err := extractToken(r)
	if err != nil {
		return nil, err
	}
	atClaims := &JWTCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, atClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errorx.New(http.StatusUnauthorized, nil, "ValidationErrorMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errorx.New(http.StatusUnauthorized, nil, "Token have already expried")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errorx.New(http.StatusUnauthorized, nil, "ValidationErrorNotValidYet")
			} else {
				return nil, errorx.New(http.StatusUnauthorized, nil, "TokenInvalid")
			}
		}
	}
	//if time.Now().Unix() > atClaims.ExpiresAt {
	//	return nil, errorx.New(http.StatusUnauthorized, nil, "Token have already expried")
	//}
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetaData(r *http.Request) (*JWTCustomClaims, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, errorx.New(http.StatusUnauthorized, err, "Can not verify token. Token does not valid")
	}

	cl, ok := token.Claims.(*JWTCustomClaims)
	if ok && token.Valid {
		return cl, nil
	}
	return nil, nil
}

func RefreshToken(refreshToken string) (*TokenDetails, error) {
	// Verify the token
	var customClaims JWTCustomClaims
	token, err := jwt.ParseWithClaims(refreshToken, customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	// If there is an error, the token must have expired
	if err != nil {
		return nil, errorx.New(http.StatusUnauthorized, nil, "Refresh token expired")
	}

	// Check token valid
	if _, ok := token.Claims.(JWTCustomClaims); !ok || !token.Valid {
		return nil, errorx.New(http.StatusUnauthorized, err, "Refresh token does not valid")
	}

	// Generate new tokens
	customClaims, _ = token.Claims.(JWTCustomClaims)
	userID := customClaims.UserID
	if userID == 0 {
		return nil, errorx.New(http.StatusUnprocessableEntity, nil, "Unauthorized")
	}

	//Create new pairs of refresh and access tokens
	ts, err := GenerateToken(userID)
	if err != nil {
		return nil, errorx.New(http.StatusForbidden, nil, err.Error())
	}
	return ts, nil
}
func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(JWTCustomClaims); !ok || !token.Valid {
		return err
	}
	return nil
}
