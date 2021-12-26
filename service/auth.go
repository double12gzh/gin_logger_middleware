package service

import (
	"net/http"
	"time"

	"main/model"
	"main/util"
)

func Auth() util.InnerResponse {
	util.LOG.Info("service Auth: start")

	time.Sleep(6 * time.Second)

	result := util.InnerResponse{
		Code: http.StatusOK,
		Data: model.Auth{
			Tel:  "1350090010",
			Addr: "beijing",
			Age:  18,
			Persons: []model.Person{
				{Name: "Auth", Sex: "boy", Company: "Auth"},
			},
		},
		Message: "success",
	}

	util.LOG.Info("service Auth: end")

	return result
}
