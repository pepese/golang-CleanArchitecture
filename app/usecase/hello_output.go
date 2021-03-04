package usecase

type HelloOutput struct {
	Reply string
}

/*
NewHelloOutput returns *helloOutput.
*/
func NewHelloOutput(reply string) *HelloOutput {
	return &HelloOutput{Reply: reply}
}
