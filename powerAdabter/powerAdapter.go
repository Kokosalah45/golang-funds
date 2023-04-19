package powerAdapter

import "cards/electronics"

type charger interface{
	GetPower()
	Plug()
	Unplug()
}
type batteryAdapter struct {
	currentPowerLevel float64
	drawingDevices    []electronics.Appliance
	charger
}

func newBatteryAdapter(powerLevel float64) *batteryAdapter {
	return &batteryAdapter{
		currentPowerLevel: powerLevel,
	}
}

func (ba *charger) GetPower(a electronics.Iappliance) {
	
}

func (ba *charger) Plug(a electronics.Iappliance) {
	
}

func (ba *charger) Unplug(a electronics.Iappliance) {
	
}
