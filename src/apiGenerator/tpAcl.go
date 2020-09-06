package main

var tpAcl = `
/***
	Author: Leong Kai Khee (Kurogami)
	Date: 2020

	Generated by DeviserGO
***/

package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// ACLMiddleware checks user ACL permission
func ACLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permitted := false

		ctx := r.Context().Value("Authorization")
		token, ok := ctx.(*jwt.Token)
		if ok {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				role, ok := claims["role"].(string)
				if ok {
					api := r.URL.Path

					acl, _ := DBAclRetrieveCondition("role = '" + role + "' and type = 'API' and key = '" + api + "'")
					permitted = len(acl) != 1
				} else {
					LogError("[acl.go] Error casting claims.role")
				}
			} else {
				LogError("[acl.go] Error casting jwt.MapClaims")
			}
		} else {
			LogError("[acl.go] Error casting jwt.Token")
		}

		if permitted {
			next.ServeHTTP(w, r)
		} else {
			result := DeviserResponse{HTTPStatus: 400, Result: "Invalid user access"}
			result.DoResponse(w)
			return
		}
	})
}
`
