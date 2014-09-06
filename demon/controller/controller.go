package main


import (
	"log"
	"github.com/kmeinke/hydra/demon"
)


func main() {
	log.Print("hi")
	d := demon.NewDemon(false)
	log.Print(d)
}