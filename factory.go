package main

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func NewAk() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name: "Musket",
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk(), nil
	}

	if gunType == "musket" {
		return NewMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}
