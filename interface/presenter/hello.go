package presenter

import "github.com/pepese/golang-CleanArchitecture/usecase"

func HelloRender(o *usecase.HelloOutput) string {
	return o.Reply
}
