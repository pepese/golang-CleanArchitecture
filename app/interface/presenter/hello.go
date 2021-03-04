package presenter

import "github.com/pepese/golang-CleanArchitecture/app/usecase"

func HelloRender(o *usecase.HelloOutput) string {
	return o.Reply
}
