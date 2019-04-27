package main

import (
	"BerryHub/boot"

	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load()

	boot.InitServer()
}
