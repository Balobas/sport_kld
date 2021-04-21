package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/controllers/organization_controller"
	"sport_kld/app/models/organization_model"
	"sport_kld/app/utils"
)

func GetOrganizationByUID(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := ctx.Request.URL.Query()["uid"][0]

	org, err := organization_controller.GetOrganizationByUID(uid)
	if err != nil {
		res := []byte(err.Error())
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}
	utils.WriteResult(ctx.Writer, org)
}

func GetOrganizationsByUIDs(ctx *gin.Context) {
	if utils.CheckHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	uids := ctx.Request.URL.Query()["uid"]
	if len(uids) == 0 {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("not found uids in get query"))
		return
	}

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	var errs []error

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByUIDs(uids)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func GetOrganizationsByFields(ctx *gin.Context) {
	if utils.CheckHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("search") == "" {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("not found 'search' key in get query"))
		return
	}

	searchString := ctx.Request.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByFields(ctx.Request.URL.Query()["fields"], searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func GetPlaceOrganization(ctx *gin.Context) {
	if utils.CheckHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("place_uid") == "" {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("not found 'organization_uid' key in get query"))
		return
	}

	placeUid := ctx.Request.URL.Query().Get("place_uid")

	var errs []error

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	resultParams.Organizations, errs = organization_controller.GetOrganizationByPlaceUID(placeUid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func GetOrganizationsByTag(ctx *gin.Context) {
	if utils.CheckHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("search") == "" {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("not found 'search' key in get query"))
		return
	}

	searchString := ctx.Request.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByTags(searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}
