package auth

import "net/http"

type pair struct {
	path string
	method string
}

var OnlyWithAuth = map[pair] bool {
	{"/place", http.MethodGet}: false,
	{"/places", http.MethodGet}: false,
	{"/places_by_fields", http.MethodGet}: false,
	{"/organization_places", http.MethodGet}: false,
	{"/places_by_tag", http.MethodGet}: false,

	{"/organization", http.MethodGet}: false,
	{"/organizations", http.MethodGet}: false,
	{"/organizations_by_fields", http.MethodGet}: false,
	{"/place_organization", http.MethodGet}: false,
	{"/organizations_by_tag", http.MethodGet}: false,

	{"/event", http.MethodGet}: false,
	{"/events_in_place", http.MethodGet}: false,
	{"/event_user_role", http.MethodGet}: true,
	{"/event_info_post", http.MethodGet}: false,
	{"/event_info_posts", http.MethodGet}: false,

	{"/event", http.MethodPost}: true,
	{"/join_event", http.MethodPost}: true,
	{"/update_event", http.MethodPost}: true,
	{"/change_event_privacy", http.MethodPost}: true,
	{"/change_user_event_role", http.MethodPost}: true,
	{"/event_info_post", http.MethodPost}: true,

	{"/event", http.MethodDelete}: true,
	{"/user_from_event", http.MethodDelete}: true,
}

