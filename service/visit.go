package service

import (
	"net/http"
	"time"

	"main/model"
	"main/util"
)

func Visit() util.InnerResponse {
	util.LOG.Info("service Visit: start")

	time.Sleep(3 * time.Second)

	result := util.InnerResponse{
		Code: http.StatusOK,
		Data: model.Auth{
			Tel:  "1350090010",
			Addr: "beijing",
			Age:  18,
			Persons: []model.Person{
				{Name: "Visit", Sex: "boy", Company: "Visit"},
			},
		},
		Message: "success",
	}

	util.LOG.Info("service.Visit: end")

	return result
}

func Do() util.InnerResponse {
	util.LOG.Info("service Visit: start")

	time.Sleep(3 * time.Second)

	result := util.InnerResponse{
		Code: http.StatusOK,
		Data: model.Auth{
			Tel:  "1350090010",
			Addr: "beijing",
			Age:  18,
			Persons: []model.Person{
				{Name: "Visit", Sex: "boy", Company: "Visit"},
			},
		},
		Message: "success",
	}

	util.LOG.Info("service.Visit: end")

	return result
}
