package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/sahandPgr/Email-verifier/controllers"
	"github.com/sahandPgr/Email-verifier/models"
)

// /check-domain handler
func FormHandler(w http.ResponseWriter, r *http.Request) {
	var domainUrl models.DomainUrl
	json.NewDecoder(r.Body).Decode(&domainUrl)
	w.Header().Set("Content-Type", "application/json")
	domainVar := controllers.CheckDomain(domainUrl.DomainUrl)
	json.NewEncoder(w).Encode(domainVar)
}
