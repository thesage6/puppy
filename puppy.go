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
func BigDogBark() string {
	return dog.WhenDogGrowsUp(Bark())
}

func BigDogSit() string {
	return dog.WhenDogGrowsUp(Sit())
}
