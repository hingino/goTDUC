package hwinfo

import (
	"fmt"
	"strings"

	"goTDUC/pkg/metadata"

	"github.com/jaypipes/ghw"
)

type GpuInfo struct {
	Vendor        string
	Name          string
	OfflineDriver string
	NewestDriver  string // need to implement
	Compatible    bool
	GpuId         string
	IsDchDriver   bool // need to implement
	IsNotebook    bool
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
	var gid string
	var err error

	cards, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v", err)
	}

	for _, card := range cards.GraphicsCards {
		// get correct gpu name as string
		var name string
		nameSlice := strings.Split(card.DeviceInfo.Product.Name, card.DeviceInfo.Vendor.Name+" ")
		for _, letter := range nameSlice {
			name = name + letter
		}

		// get gpuId
		gid, err = metadata.GetGpuId(card.DeviceInfo.Vendor.Name, name, GetChassisInfo())
		if err != nil {
			panic(err)
		}

		gpu := GpuInfo{
			Vendor:        card.DeviceInfo.Vendor.Name,
			Name:          name,
			OfflineDriver: card.DeviceInfo.Driver,
			GpuId:         gid,
			Compatible:    false,
			IsNotebook:    GetChassisInfo(),
		}

		if gid != "" {
			gpu.Compatible = true
		}

		gpus = append(gpus, gpu)

	}

	return gpus
}
