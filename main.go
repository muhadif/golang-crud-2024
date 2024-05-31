package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang-crud-2024/cli"
	validatorPkg "golang-crud-2024/pkg/validator"
	"log"
)

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("enum", validatorPkg.ValidateEnum)
		if err != nil {
			log.Fatal("Error Register Validator", err)
			return
		}
	}

	cli.Run()
}
