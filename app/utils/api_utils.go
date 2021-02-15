package utils

import (
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
