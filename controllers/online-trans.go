package controllers

import "fmt"

type OnlineTranslate struct {
}

func (ot *OnlineTranslate) Text(sl, tl, text string) string {
	fmt.Println("I'm a man")
	return ""
}
