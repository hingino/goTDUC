package main

import (
	"fmt"
	"goTDUC/pkg/hwinfo"
	"goTDUC/pkg/metadata"
)

func main() {
	gpus := hwinfo.GetGpuData()
	isNotebook := hwinfo.GetChassisInfo()

	println("TEST 1: check if device is notebook")
	println("Is Notebook?: ", isNotebook, "\n")

	println("TEST 2: Get metadata for all detected GPUs\n")
	for i, gpu := range gpus {
		fmt.Printf("TEST 2.%v: Get online metadata", i+1)

		println("GPU Vendor: ", gpu.Vendor)
		println("GPU Name: ", gpu.Name)
		println("Driver: ", gpu.OfflineDriver)

		gpuId, err := metadata.GetGpuId(gpu.Vendor, gpu.Name, isNotebook)
		if err != nil {
			panic(err)
		}
		// deref_metadata := *metadata
		println("GPU ID: ", gpuId, "\n")
		// println(deref_metadata)
	}

}
