package usecase

type helloInput struct {
	name string
}

/*
NewHelloInput returns *helloInput.
*/
func NewHelloInput(name string) *helloInput {
	return &helloInput{name: name}
}
