package main

import (
	"errors"
)

var ErrConnFailed = errors.New("[db]数据库连接失败")

type DBError struct {
	msg string
}

func (e *DBError) Error() string {
	return e.msg
}

func ConnectDB(username, password string) error {
	if password != "123" {
		return ErrConnFailed
	}

	if password == "123" {
		return &DBError{
			msg: "弱密码",
		}
	}

	return nil
}
