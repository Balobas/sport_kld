package main

import (
	"sport_kld/test_app_api/events"
	"sport_kld/test_app_api/places"
	"sport_kld/test_app_api/users"
)

func main() {
	places.TestGetPlaceByUID()
	places.TestGetPlacesByUIDs()

	events.TestCreateEvent()
	users.TestCreateUser()
}
