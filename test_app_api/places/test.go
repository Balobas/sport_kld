package places

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sport_kld/test_app_api/config"
)

func TestGetPlaceByUID() {
	resp, err := http.Get(config.BaseUrl + "/place?uid=" + "fdc346025f0111eb82530c9d92446328")
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("DONE : GetPlaceByUID")
}

func TestGetPlacesByUIDs() {
	resp, err := http.Get(config.BaseUrl + "/places?uid=fdc346025f0111eb82530c9d92446328&uid=fdc3bade5f0111eba9290c9d92446328&uid=fdc5676f5f0111ebb95f0c9d92446328")
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("DONE : GetPlacesByUIDs")
}