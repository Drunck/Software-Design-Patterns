package main

type Weapon interface {
	getDamage() float64
	getDescription() string
	getAccuracy() float64
}
