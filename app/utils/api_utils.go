package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func HandleHTTPMethod(w http.ResponseWriter, requiredMethod, inputMethod string) error {
	if inputMethod != requiredMethod {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return errors.New("")
	}
	return nil
}

func WriteToResponseWriter(w http.ResponseWriter, message []byte) error {
	if _, err := w.Write(message); err != nil {
		fmt.Println("cant write bytes :", err)
		return err
	}
	return nil
}

func WriteResult(w http.ResponseWriter, result interface{}) {
	b, err := json.Marshal(result)
	if err != nil {
		if _, err := w.Write([]byte("cant marshal json")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("cant write bytes")
	}
}