/*
File Name:  Main.go
Copyright:  2022
Author:     Emir Skopljak
*/

package main

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator
	Hatched       bool
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.Hatched {
		return nil
	}
	egg.Hatched = true
	return egg.CreateReptile()
}

type FireDragon struct {
	Name   string
	Height int
}

func (fireDragon *FireDragon) Lay() ReptileEgg {
	var egg = ReptileEgg{
		CreateReptile: func() Reptile {
			return &FireDragon{
				Name:   "Dragon baby",
				Height: 1,
			}
		},
	}
	return egg
}
