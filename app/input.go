package main

import(
	"myitcv.io/react"
)

type InputDef struct {
	react.ComponentDef
}

type InputProps struct {
	Type string
	Value string
	Key string
	Label string
}

func Input(props InputProps) *InputElem {
	return buildInputElem(props)
}

func (i InputDef) Render() react.Element {
	return react.Div(nil,
		react.Label(&react.LabelProps{}, react.S(i.Props().Label)),
		react.Input(&react.InputProps{
			Type:		i.Props().Type,
			ClassName:	"form-control",
			Key:		i.Props().Key,
			Value:		i.Props().Value,
	}))
}

