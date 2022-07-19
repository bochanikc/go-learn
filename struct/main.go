package main

import "fmt"

type Weapon struct {
	On    bool
	Ammo  int
	Power int
}

func main() {
	testStruct := new(Weapon)
	testStruct.On = true
	testStruct.Ammo = 10
	fmt.Print(testStruct.Shoot())
}

func (w *Weapon) Shoot() bool {
	if w.On != false {
		if w.Ammo > 0 {
			w.Ammo--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (w *Weapon) RideBike() bool {
	if w.On != false {
		if w.Power > 0 {
			w.Power--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
