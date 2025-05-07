package event_delivery_http

import "net/http"

func (self *eventController) EventsAttendance(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
