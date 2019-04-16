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

func addressGetHandler(w http.ResponseWriter, r *http.Request, c *cases.Cases) {

	req := &models.GetAddressRequest{}

	vars := mux.Vars(r)
	addressIDString := vars["addressID"]
	addressID, _ := strconv.ParseInt(addressIDString, 10, 32)
	req.ID = int(addressID)

	resp, err := c.GetAddress(req)
	if err != nil {
		log.Printf("Error getting address: %s", err.Error())
		w.Write([]byte("{ \"error\": \"Error getting address: " + err.Error() + "\" }"))
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
