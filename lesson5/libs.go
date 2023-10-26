package main

import (
	"fmt"
	"net/http"

	"lesson5/magic"
	anotherMagic "lesson5/useful_module/magic_two"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DefaultLibs() {
	fmt.Println("The Doors", http.StatusAccepted)
}

func InternalLibs() {
	magic.NewMagic()

	anotherMagic.PrintBaka()
}

func ExternalLibs() {
	gin.New()
	_ = uuid.NewString()
}
