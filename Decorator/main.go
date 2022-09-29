package main

import "fmt"

func main() {
	weapon1 := &AssaultRifle{}

	weapon1FlameArrestor := &FlameArrestor{
		weapon: weapon1,
	}

	weapon1FlameArrestorSilencer := &Silencer{
		weapon: weapon1FlameArrestor,
	}

	weapon1FlameArrestorSilencerRedDotSight := &RedDotSight{
		weapon: weapon1FlameArrestorSilencer,
	}

	fmt.Printf("%v \n damage - %v \n accuracy - %v",
		weapon1FlameArrestorSilencerRedDotSight.getDescription(),
		weapon1FlameArrestorSilencerRedDotSight.getDamage(),
		weapon1FlameArrestorSilencerRedDotSight.getAccuracy())

	weapon2 := &CutawayRifle{}

	weapon2straightLineStock := &StraightLineStock{
		weapon: weapon2,
	}

	weapon2straightLineStockSilencer := &Silencer{
		weapon: weapon2straightLineStock,
	}

	fmt.Println("\n-----------------------")
	fmt.Printf("%v \n damage - %v \n accuracy - %v",
		weapon2straightLineStockSilencer.getDescription(),
		weapon2straightLineStockSilencer.getDamage(),
		weapon2straightLineStockSilencer.getAccuracy())
}
