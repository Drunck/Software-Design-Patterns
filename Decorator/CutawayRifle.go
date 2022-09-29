package main

type CutawayRifle struct {
}

func (c *CutawayRifle) getDamage() float64 {
	return 30
}

func (c *CutawayRifle) getDescription() string {
	return "CutawayRifle"
}

func (c *CutawayRifle) getAccuracy() float64 {
	return 24
}
