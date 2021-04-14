package api

import "net/http"

var OnlyWithAuth = map[string]bool{
	http.MethodGet + "/place":               false,
	http.MethodGet + "/places":              false,
	http.MethodGet + "/places_by_fields":    false,
	http.MethodGet + "/organization_places": false,
	http.MethodGet + "/places_by_tag":       false,

	http.MethodGet + "/organization":            false,
	http.MethodGet + "/organizations":           false,
	http.MethodGet + "/organizations_by_fields": false,
	http.MethodGet + "/place_organization":      false,
	http.MethodGet + "/organizations_by_tag":    false,

	http.MethodGet + "/event":            false,
	http.MethodGet + "/events_in_place":  false,
	http.MethodGet + "/event_user_role":  true,
	http.MethodGet + "/event_info_post":  true,
	http.MethodGet + "/event_info_posts": true,

	http.MethodPost + "/event":                  true,
	http.MethodPost + "/join_event":             true,
	http.MethodPost + "/update_event":           true,
	http.MethodPost + "/change_event_privacy":   true,
	http.MethodPost + "/change_user_event_role": true,
	http.MethodPost + "/event_info_post":        true,

	http.MethodDelete + "/event":           true,
	http.MethodDelete + "/user_from_event": true,

	http.MethodGet + "/user":          false,
	http.MethodGet + "/user_by_login": false,

	http.MethodPost + "/user":        false,
	http.MethodPost + "/update_user": true,

	http.MethodPost + "/login":   false,
	http.MethodPost + "/logout":  true,
	http.MethodPost + "/refresh": false,
}
