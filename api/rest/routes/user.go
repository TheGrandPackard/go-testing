package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thegrandpackard/go-testing/cases"
	"github.com/thegrandpackard/go-testing/models"
)

func userGetHandler(w http.ResponseWriter, r *http.Request, c *cases.Cases) {

	req := &models.GetUserRequest{}

	vars := mux.Vars(r)
	userIDString := vars["userID"]
	userID, _ := strconv.ParseInt(userIDString, 10, 32)
	req.ID = int(userID)

	resp, err := c.GetUser(req)
	if err != nil {
		log.Printf("Error getting user: %s", err.Error())
		w.Write([]byte("{ \"error\": \"Error getting user: " + err.Error() + "\" }"))
		return
	}

	response, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("{ \"error\": \"Error marshalling response: " + err.Error() + "\" }"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status", strconv.FormatInt(http.StatusOK, 10))
	w.Write(response)

	return
}
