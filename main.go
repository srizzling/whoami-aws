package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetCallerIdentity() (*sts.GetCallerIdentityOutput, error) {
	svc := sts.New(session.New())
	input := &sts.GetCallerIdentityInput{}

	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return nil, aerr
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return nil, err
		}

	}
	return result, nil
}

func main() {
	http.HandleFunc("/iam", func(w http.ResponseWriter, r *http.Request) {
		result, err := GetCallerIdentity()
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("all good")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
