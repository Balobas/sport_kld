package places

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"sport_kld/test_app_api/config"
)

func TestGetPlaceByUID() {

	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("uid", "fdc31f025f0111ebafd20c9d92446328")
	buffer.WriteString(params.Encode())
	req, err := http.NewRequest(http.MethodGet, config.BaseUrl, buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.Client{}.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
