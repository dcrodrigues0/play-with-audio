package models

type Arecord struct{
	Card int
	Name string
	InputType string
}

func NewArecord(card int, name, inputType string) *Arecord {
	return &Arecord{
		Card: card,
		Name: name,
		InputType: inputType,
	}
}
