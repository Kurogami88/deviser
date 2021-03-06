package main

var apiPurgeToken = `/***
	Author: Leong Kai Khee (Kurogami)
	Date: 2020

	Generated by DeviserGO
***/

package main

import (
	"net/http"
	"os"
	"time"
)

//PurgeToken is an API to purge tokens that are expired
func PurgeToken(w http.ResponseWriter, r *http.Request) {
	result := DeviserResponse{HTTPStatus: 200, Result: "Successfully purged tokens"}

	envAuthDuration, err := time.ParseDuration(os.Getenv("JWT_EXPIRY_MIN") + "m")
	if err != nil {
		result = DeviserResponse{HTTPStatus: 400, Result: "Error purging tokens"}
	}

	err = DBTokenDeleteUnscopeCondition("` + "`created_at`" + ` < '" + time.Now().Add(-envAuthDuration).UTC().Format("2006-01-02 15:04:05") + "'")
	if err != nil {
		result = DeviserResponse{HTTPStatus: 400, Result: "Error purging tokens"}
	}
	
	result.DoResponse(w)
	return
}
`
