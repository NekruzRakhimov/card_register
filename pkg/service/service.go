package service

import (
	"card_register/models"
	"card_register/pkg/repository"
)

func AddNewOrderInfo(info models.Info) error {
	return repository.AddNewOrderInfo(info)
}

func GetAllInfo() (info []models.Info, err error) {
	return repository.GetAllInfo()
}
