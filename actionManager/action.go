package actionManager

import (
	"errors"
	"net/http"
)

type CallBack func(res http.ResponseWriter, r *http.Request) string

type Action struct {
	CallFn map[string]CallBack
}

var RemoveError = errors.New("无改名字的函数")

func NewAction() *Action {
	return &Action{CallFn: make(map[string]CallBack)}
}

func (a *Action) AddAction(s string, fn CallBack) {
	a.CallFn[s] = fn
}

func (a *Action) RemoveAction(s string) (error, CallBack) {
	removeFn, ok := a.CallFn[s]
	if ok {
		delete(a.CallFn, s)
		return nil, removeFn
	}
	return RemoveError, nil
}

func (a *Action) GetAction(s string) (error, CallBack) {
	getFn, ok := a.CallFn[s]
	if ok {
		return nil, getFn
	}
	return RemoveError, nil
}


