package main

import (
	"fmt"
	"goTDUC/pkg/hwinfo"
	"goTDUC/pkg/metadata"
)

func main() {
	gpus := hwinfo.GetGpuData()
	isNotebook := hwinfo.GetChassisInfo()

	println("TEST 1: Get info for installed gpus")
	println("Is Notebook?: ", isNotebook)
	for i, gpu := range gpus {
		println("GPU Vendor: ", gpu.Vendor)
		println("GPU Name: ", gpu.Name)
		println("Driver: ", gpu.OfflineDriver)

		fmt.Printf("TEST 2.%v: Get online metadata", i+1)
		gpuId, err := metadata.GetGpuId(gpu.Vendor, gpu.Name, isNotebook)
		if err != nil {
			panic(err)
		}
		// deref_metadata := *metadata
		println("GPU ID: ", gpuId)
		// println(deref_metadata)
	}

}
