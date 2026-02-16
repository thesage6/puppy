package  puppy

import (
	"github.com/thesage6/dog"
)

func Bark() string {
	return "Woof!"
}

func Sit() string {
	return "Okay, I'm sitting."
}

func Barks() string{
	return "Woof! Woof! Woof!"
}


func BigBarks() string{
	return dog.WhenDogGrowsUp(Bark())
}