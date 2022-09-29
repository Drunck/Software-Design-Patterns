package main

type Accessories interface {
	getDamage() float64
	getDescription() string
	getAccuracy() float64
}

type Silencer struct {
	weapon Weapon
}

func (s *Silencer) getDamage() float64 {
	weaponDamage := s.weapon.getDamage()
	return weaponDamage - 5
}

func (s *Silencer) getDescription() string {
	weaponDescription := s.weapon.getDescription()
	return weaponDescription + "\ninstalled Silencer"
}

func (s *Silencer) getAccuracy() float64 {
	weaponAccuracy := s.weapon.getAccuracy()
	return weaponAccuracy + 2
}

type FlameArrestor struct {
	weapon Weapon
}

func (f *FlameArrestor) getDamage() float64 {
	weaponDamage := f.weapon.getDamage()
	return weaponDamage + 2
}

func (f *FlameArrestor) getDescription() string {
	weaponDescription := f.weapon.getDescription()
	return weaponDescription + "\ninstalled Flame Arrestor"
}

func (f *FlameArrestor) getAccuracy() float64 {
	weaponAccuracy := f.weapon.getAccuracy()
	return weaponAccuracy - 5
}

type OpticalSight struct {
	weapon Weapon
}

func (f *OpticalSight) getDamage() float64 {
	weaponDamage := f.weapon.getDamage()
	return weaponDamage
}

func (o *OpticalSight) getDescription() string {
	weaponDescription := o.weapon.getDescription()
	return weaponDescription + "\ninstalled Optical Sight"
}

func (o *OpticalSight) getAccuracy() float64 {
	weaponAccuracy := o.weapon.getAccuracy()
	return weaponAccuracy + 20
}

type RedDotSight struct {
	weapon Weapon
}

func (f *RedDotSight) getDamage() float64 {
	weaponDamage := f.weapon.getDamage()
	return weaponDamage
}

func (r *RedDotSight) getDescription() string {
	weaponDescription := r.weapon.getDescription()
	return weaponDescription + "\ninstalled Red Dot Sight"
}

func (r *RedDotSight) getAccuracy() float64 {
	weaponAccuracy := r.weapon.getAccuracy()
	return weaponAccuracy + 5
}

type StraightLineStock struct {
	weapon Weapon
}

func (s *StraightLineStock) getDamage() float64 {
	weaponDamage := s.weapon.getDamage()
	return weaponDamage
}

func (s *StraightLineStock) getDescription() string {
	weaponDescription := s.weapon.getDescription()
	return weaponDescription + "\ninstalled Straight Line Butt Stock"
}

func (s *StraightLineStock) getAccuracy() float64 {
	weaponAccuracy := s.weapon.getAccuracy()
	return weaponAccuracy + 2
}
