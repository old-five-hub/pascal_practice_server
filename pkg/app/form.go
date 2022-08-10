package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"net/http"
	"pascal_practice_server/pkg/e"
)

func Valid(form interface{}) (int, int) {

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	fmt.Println(err)

	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
