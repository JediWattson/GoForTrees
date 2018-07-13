package main

import(
	"myitcv.io/react"
)

type LoginDef struct {
	react.ComponentDef
}

func Login() *LoginElem{
	return buildLoginElem()
}

func (f LoginDef) Render() react.Element{
	return FormWrap(
		react.Div(nil,
			Input(InputProps{Key: "username", Type: "text", Label: "User Handle:"}),
			Input(InputProps{Key: "Password", Type: "password", Label: "Password:"}),
	))
}
