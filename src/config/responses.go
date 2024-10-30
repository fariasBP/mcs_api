package config

import "time"

type (
	Res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	ResErr struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Err  string `json:"err"`
	}
	resJson struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	resJsonCount struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Data  interface{} `json:"data"`
		Count int64       `json:"count"`
	}
	resToken struct {
		Code    int       `json:"code"`
		Msg     string    `json:"msg"`
		Token   string    `json:"token"`
		Expires time.Time `json:"expires"`
	}
)

func SetRes(code int, msg string) *Res {
	d := &Res{
		Code: code,
		Msg:  msg,
	}
	return d
}

func SetResError(code int, msg string, err string) *ResErr {
	d := &ResErr{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
	return d
}

func SetResJson(code int, msg string, json interface{}) *resJson {
	d := &resJson{
		Code: code,
		Msg:  msg,
		Data: json,
	}
	return d
}

func SetResJsonCount(code int, msg string, count int64, json interface{}) *resJsonCount {
	d := &resJsonCount{
		Code:  code,
		Msg:   msg,
		Data:  json,
		Count: count,
	}
	return d
}

func SetResToken(code int, msg string, token string, expires time.Time) *resToken {
	d := &resToken{
		Code:    code,
		Msg:     msg,
		Token:   token,
		Expires: expires,
	}
	return d
}
