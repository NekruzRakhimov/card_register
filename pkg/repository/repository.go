package repository

import (
	"card_register/db"
	"card_register/models"
)

func AddNewOrderInfo(info models.Info) error {
	if err := db.GetDBConn().Table("info").Create(&info).Error; err != nil {
		return err
	}

	return nil
}

func GetAllInfo() (info []models.Info, err error) {
	sqlQuery := "SELECT * FROM info"
	if err = db.GetDBConn().Raw(sqlQuery).Scan(&info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func GetUser(login, password string) (user models.User, err error) {
	sqlQuery := "SELECT * FROM \"user\" WHERE \"login\" = ? AND \"password\" = ?"
	if err = db.GetDBConn().Raw(sqlQuery, login, password).Scan(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
