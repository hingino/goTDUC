package hwinfo

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type GpuInfo struct {
	Vendor        string
	Name          string
	OfflineDriver string
	NewestDriver  string
	Compatible    bool
}

type DeviceInfo struct {
	Gpus       []GpuInfo
	IsNotebook bool
}

func GetChassisInfo() bool {
	chassis, err := ghw.Chassis()

	if err != nil {
		fmt.Printf("Error getting chassis info: %v", err)
	}

	switch chassis.Type {
	case "Laptop":
		return true
	default:
		return false
	}
}

func GetGpuData() []GpuInfo {
	gpus := []GpuInfo{}

	cards, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v", err)
	}

	for _, card := range cards.GraphicsCards {
		gpu := GpuInfo{
			Name:          card.DeviceInfo.Product.Name,
			Vendor:        card.DeviceInfo.Vendor.Name,
			OfflineDriver: card.DeviceInfo.Driver,
		}

		gpus = append(gpus, gpu)
	}

	return gpus
}

func GetHwInfo() DeviceInfo {
	hwinfo := DeviceInfo{
		Gpus:       GetGpuData(),
		IsNotebook: GetChassisInfo(),
	}

	return hwinfo
}
