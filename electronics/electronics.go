package electronics

import "fmt"

type powerDrawer interface {
	DrawPower(power float64) bool
	GetCurrentPowerLevel() float64
}
type openerCloser interface {
	TurnOn() bool
	TurnOff() bool
} 

type electronic struct {
	id                string
	isOn             bool
	applianceType     string
	openerCloser
}


func (a *Appliance) TurnOn() bool {
	if !a.isOn {
		a.isOn = true
		return true
	}
	return false

}
func (a *Appliance) TurnOff() bool {
	if a.isOn {
		a.isOn = false
		return true
	}
	return false

}
func (a *Appliance) Draw(power float64) bool {
	a.currentPowerLevel += power
	return true

}
func (a *Appliance) getIsOn() {
	fmt.Println(a.isOn)
}
func (a *Appliance) GetCurrentPowerLevel() {
	fmt.Println(a.currentPowerLevel)
}
func (a *Appliance) GetApplianceType() {
	fmt.Println(a.applianceType)
}


// type washingMachine struct {
// 	*appliance
// }
// func newWashingMachine() *washingMachine {
// 	return &washingMachine{
// 		appliance: &appliance{
// 			currentPowerLevel: 0,
// 			applianceType:     "Washing Machine",
// 		},
// 	}
// }
// type coffeMaker struct {
// 	*appliance
// }
// func newCoffeMaker() *coffeMaker {
// 	return &coffeMaker{
// 		appliance: &appliance{
// 			currentPowerLevel: 0,
// 			applianceType:     "Coffe Maker",
// 		},
// 	}
// }
