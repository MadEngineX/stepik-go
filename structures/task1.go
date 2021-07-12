package main

import (
	"fmt"
	"math"
)

type Human struct {
	On bool
	Ammo, Power int
}

func (input *Human) Shoot() bool {
	if input.On{
		if input.Ammo > 0 {
			input.Ammo = input.Ammo - 1
			return true
		} else {
			return false
		}
	}
	return false
}

func (input *Human) RideBike() bool {
	if input.On{
		if input.Power > 0 {
			input.Power = input.Power - 1
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	newOne := Human{}
	testStruct := &newOne
}
