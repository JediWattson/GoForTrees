package main

import(
	"myitcv.io/react"
	"fmt"
)

type FormWrapDef struct {
	react.ComponentDef
}
/*
type FormWrapState struct {
	fields map[string]interface{}
}
*/

func FormWrap(c react.Element) *FormWrapElem {
	return buildFormWrapElem(c)
}

func (f FormWrapDef) Render() react.Element {
	fmt.Printf("%#v", f.Children()[0])
	return react.Form(&react.FormProps{
		ClassName: "form-inline",
	}, f.Children()...)
}
