package event_controller

import "sport_kld/app/models/event_model"

func PutEventInfoPost(post event_model.EventInfoPost) (string, error) {
	uid, err := event_model.PutEventInfoPost(post)
	return string(uid), err
}
