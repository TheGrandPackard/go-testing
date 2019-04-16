package routes

import "github.com/thegrandpackard/go-testing/rest"

func RegisterRoutes(r *rest.RESTService) (err error) {

	r.HandleFunc("/user/{userID:[0-9]+}", userGetHandler)
	r.HandleFunc("/address/{addressID:[0-9]+}", addressGetHandler)

	return
}
