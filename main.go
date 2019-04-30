package main

import (
	"github.com/BerryHub/boot"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	boot.InitServer()
}
