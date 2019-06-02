package zhanst

import "fmt"

type Error struct {
	Code     int
	HttpCode int
	Msg      string
}

func (err Error) Error() string {
	return fmt.Sprintf("code: %v http code: %v msg: %v", err.Code, err.HttpCode, err.Msg)
}
