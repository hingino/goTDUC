package metadata

/*
type nvidiaGpus struct {
	desktop  []string
	notebook []string
}
*/

func NvidiaGpuId(gpuName string, isNotebook bool) (string, error) {
	var gpuId string
	var gpus map[string]interface{}
	var ok bool
	//var queryString string
	gpuDataURL := "https://raw.githubusercontent.com/ZenitH-AT/nvidia-data/main/gpu-data.json"

	/*
		if isNotebook {
			queryString = "notebook." + gpuName
		} else {
			queryString = "desktop." + gpuName
		}
	*/

	// get latest known gpuid list
	res, err := GetOnlineJson(gpuDataURL)
	if err != nil {
		return "", err
	}

	switch isNotebook {
	case true:
		gpus, ok = res.(map[string]interface{})["notebook"].(map[string]interface{})
	default:
		gpus, ok = res.(map[string]interface{})["desktop"].(map[string]interface{})
	}

	if !ok {
		panic("gpu id lookup: type error")
	}

	gpuId, ok = gpus[gpuName].(string)
	if !ok {
		panic("gpu id lookup: assertion error")
	}

	return gpuId, nil
}
