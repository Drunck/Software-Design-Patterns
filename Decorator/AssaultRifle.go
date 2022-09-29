package main

type AssaultRifle struct {
}

func (a *AssaultRifle) getDamage() float64 {
	return 20
}

func (a *AssaultRifle) getDescription() string {
	return "Assault Rifle"
}

func (a *AssaultRifle) getAccuracy() float64 {
	return 20
}
