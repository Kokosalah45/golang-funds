package main

import (
	"fmt"
)

type Attacker interface {
	Attack(name string, action string, WielderName string)
	Stop()
}
type WeaponAttacker struct {
	isAttacking bool
}

func (wa *WeaponAttacker) Attack(weilderName string, action string, weaponName string) {
	if !wa.isAttacking {
		wa.isAttacking = true
		fmt.Printf("%s is %s the %s \n", weilderName, action, weaponName)
		return
	}
	fmt.Printf("%s is Already Attacking With his %s \n", weilderName, weaponName)

}	
func (wa *WeaponAttacker) Stop() {
	if wa.isAttacking {
		wa.isAttacking = false
	}
}

type WeaponBase struct {
	name        string
	action      string
	WielderName string
}
type NinjaSword struct {
	WeaponBase
	WeaponAttacker
}

func NewNinjaSword(name string) *NinjaSword {
	return &NinjaSword{
		WeaponBase: WeaponBase{name: "Sword", action: "Swinging", WielderName: name},
	}
}
func (ns NinjaSword) somethingOnlySpecificToNinjaSword() {
	fmt.Printf("I'm Specific to NinjaSword")
}

type NinjaStar struct {
	WeaponBase
	WeaponAttacker
}

func NewNinjaStar(name string) *NinjaStar {
	return &NinjaStar{
		WeaponBase: WeaponBase{name: "Star", action: "Throwing", WielderName: name},
	}
}
func (ns NinjaStar) somethingOnlySpecificToNinjaStar() {
	fmt.Printf("I'm Specific to NinjaStar")
}

type UnknownMystical struct {
	WeaponBase
}

func NewUnkownMysticalWeapon(name string) *UnknownMystical {
	return &UnknownMystical{
		WeaponBase: WeaponBase{name: "Unkown", action: "Xaction", WielderName: name},
	}
}

func UseWeapons(att []Attacker) {
	for _, weapon := range att {
		switch v := weapon.(type) {
		case *NinjaStar :
			v.Attack(v.WielderName, v.action, v.name)
		case *NinjaSword:
			v.Attack(v.WielderName, v.action, v.name)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
}
func StopWeapons(att []Attacker) {
	for _, weapon := range att {
		switch v := weapon.(type) {
		case *NinjaStar:
			v.Stop()
		case *NinjaSword:
			v.Stop()
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
}
func UseWeaponsError(att []Attacker) {
	for _, weapon := range att {
		weapon.(*NinjaStar).Attack("ss", "sasd", "as")
	}
}
func OnlyCallAttackOnNinjaStar(att []Attacker) {
	for _, weapon := range att {

		if val, ok := weapon.(*NinjaStar); ok {
			val.Attack("Ss", "sss", "asdasd")
		}

	}
}
func callConcreteTypesSpecificFunctions(att []Attacker) {
	for _, weapon := range att {
		val, ok := weapon.(*NinjaStar)
		if ok {
			val.somethingOnlySpecificToNinjaStar()
		}
	}
}

func callConcreteTypesSpecificFunctionsV2(att []Attacker) {
	for _, weapon := range att {

		switch v := weapon.(type) {
		case *NinjaStar:
			v.somethingOnlySpecificToNinjaStar()
		case *NinjaSword:
			v.somethingOnlySpecificToNinjaSword()
		default:
			fmt.Printf("I don't know about type %T!\n", v)

		}
	}
}
func main() {

	attackingWeapons := []Attacker{
		NewNinjaStar("Destro45"),
		NewNinjaSword("Destro45"),
		NewNinjaStar("Destro455"),
	}
	UseWeapons(attackingWeapons)
	

}
