package service

import (
	"aioc/internal/respository/mysql"
	"aioc/pkg/setting"
)

func ServiceInit() error {
	// init service
	_, err := mysql.DBConnect(setting.DatabaseSetting.Host, setting.DatabaseSetting.Name, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.ServerSetting.RunMode == "debug")
	if err != nil {
		return err
	}
	return err
}
