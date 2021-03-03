package server

import (
	"github.com/gin-gonic/gin"
	"sport_kld/app/api"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/place", gin.WrapF(api.GetPlaceByUID))
	router.GET("/places", gin.WrapF(api.GetPlacesByUIDs))
	router.GET("/places_by_fields", gin.WrapF(api.GetPlacesByFields))
	router.GET("/organization_places", gin.WrapF(api.GetOrganizationPlaces))
	router.GET("/places_by_tag", gin.WrapF(api.GetPlacesByTag))

	router.GET("/organization", gin.WrapF(api.GetOrganizationByUID))
	router.GET("/organizations", gin.WrapF(api.GetOrganizationsByUIDs))
	router.GET("/organizations_by_fields", gin.WrapF(api.GetOrganizationsByFields))
	router.GET("/place_organization", gin.WrapF(api.GetPlaceOrganization))
	router.GET("/organizations_by_tag", gin.WrapF(api.GetOrganizationsByTag))

	router.GET("/event", gin.WrapF(api.GetEventByUid))
	router.GET("/events_in_place", gin.WrapF(api.GetEventsByPlace))
	router.GET("/event_user_role", gin.WrapF(api.GetEventUserRole))
	router.GET("/event_info_post", gin.WrapF(api.GetEventInfoPost))
	router.GET("/event_info_posts", gin.WrapF(api.GetEventInfoPosts))


	router.POST("/create_event", gin.WrapF(api.CreateEvent))
	router.POST("/join_event", gin.WrapF(api.JoinEvent))
	router.POST("/update_event", gin.WrapF(api.UpdateEvent))
	router.POST("/change_event_privacy", gin.WrapF(api.ChangeEventPrivateStatus))
	router.POST("/change_user_event_role", gin.WrapF(api.ChangeUserEventRole))
	router.POST("/event_info_post", gin.WrapF(api.PutEventInfoPost))


	router.DELETE("/event", gin.WrapF(api.DeleteEvent))
	router.DELETE("/user_from_event", gin.WrapF(api.DeleteUserFromEvent))
	router.DELETE("/event_info_post", gin.WrapF(api.DeleteEventInfoPost))

}