package router

import (
	"github.com/gin-gonic/gin"
	"sport_kld/app/api"
	serverAuthApi "sport_kld/server/auth/api"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/place", api.GetPlaceByUID)
	router.GET("/places", api.GetPlacesByUIDs)
	router.GET("/places_by_fields", api.GetPlacesByFields)
	router.GET("/organization_places", api.GetOrganizationPlaces)
	router.GET("/places_by_tag", api.GetPlacesByTag)

	router.GET("/organization", api.GetOrganizationByUID)
	router.GET("/organizations", api.GetOrganizationsByUIDs)
	router.GET("/organizations_by_fields", api.GetOrganizationsByFields)
	router.GET("/place_organization", api.GetPlaceOrganization)
	router.GET("/organizations_by_tag", api.GetOrganizationsByTag)

	router.GET("/event", api.GetEventByUid)
	router.GET("/events_in_place", api.GetEventsByPlace)
	router.GET("/event_user_role", api.GetEventUserRole)
	router.GET("/event_info_post", api.GetEventInfoPost)
	router.GET("/event_info_posts", api.GetEventInfoPosts)

	router.GET("/user", api.GetUser)
	router.GET("/user_by_login", api.GetUserByLogin)

	router.POST("/event", api.CreateEvent)
	router.POST("/join_event", api.JoinEvent)
	router.POST("/update_event", api.UpdateEvent)
	router.POST("/change_event_privacy", api.ChangeEventPrivateStatus)
	router.POST("/change_user_event_role", api.ChangeUserEventRole)
	router.POST("/event_info_post", api.PutEventInfoPost)

	router.POST("/user", api.CreateUser)
	router.POST("/update_user", api.UpdateUser)

	router.POST("/login", serverAuthApi.Login)
	router.POST("/logout", serverAuthApi.Logout)
	router.POST("/refresh", serverAuthApi.Refresh)


	router.DELETE("/event", api.DeleteEvent)
	router.DELETE("/user_from_event", api.DeleteUserFromEvent)
	router.DELETE("/event_info_post", api.DeleteEventInfoPost)

}
