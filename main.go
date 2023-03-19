package main

import (
	"goTDUC/hwinfo"
)

func main() {
	hw := hwinfo.GetHwInfo()

	println("Info for installed gpus")
	for _, gpu := range hw.Gpus {
		println("GPU Vendor: ", gpu.Vendor)
		println("GPU Name: ", gpu.Name)
		println("Driver: ", gpu.OfflineDriver)
	}

	println("Is Notebook?: ", hw.IsNotebook)
}
