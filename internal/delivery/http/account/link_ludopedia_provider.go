package account_delivery_http

import "net/http"

func (self *accountController) LinkLudopediaProvider(w http.ResponseWriter, r *http.Request) {
	reqId := self.idAdapter.GenReqId()

	logger := self.logger.With().
		Str("dmn", "Account").
		Str("mtd", "GetProfileByHandle").
		Str("reqId", reqId).
		Logger()

	// Add logic here

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
