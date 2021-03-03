/*
Данный пакет содержит методы обработчики запросов

*/
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/controllers/place_controller"
	"sport_kld/app/models/place_model"
	"sport_kld/app/utils"
)

func GetPlaceByUID(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("uid") == "" {
		if utils.WriteToResponseWriter(ctx.Writer, []byte("not found uid key in get query")) != nil {
			return
		}
		return
	}

	uid := ctx.Request.URL.Query()["uid"][0]

	place, err := place_controller.GetPlaceByUID(uid)
	if err != nil {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte(err.Error()))
		return
	}

	utils.WriteResult(ctx.Writer, place)
}

func GetPlacesByUIDs(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uids := ctx.Request.URL.Query()["uid"]
	if len(uids) == 0 {
		if _, err := ctx.Writer.Write([]byte("not found uids in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	resultParams := struct {
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	var errs []error

	resultParams.Places, errs = place_controller.GetPlacesByUIDs(uids)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func GetPlacesByFields(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		if utils.WriteToResponseWriter(ctx.Writer, []byte("wrong http method. access denied")) != nil {
			return
		}
		return
	}

	if ctx.Request.URL.Query().Get("search") == "" {
		if _, err := ctx.Writer.Write([]byte("not found 'search' key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	searchString := ctx.Request.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByFields(ctx.Request.URL.Query()["fields"], searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func GetOrganizationPlaces(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("organization_uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found 'organization_uid' key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	orgUid := ctx.Request.URL.Query().Get("organization_uid")

	var errs []error

	resultParams := struct {
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByOrganizationUID(orgUid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func GetPlacesByTag(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("search") == "" {
		if _, err := ctx.Writer.Write([]byte("not found 'search' key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	searchString := ctx.Request.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByTags(searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}
