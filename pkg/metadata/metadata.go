package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOnlineJson(url string) (interface{}, error) {
	// http get request to retrieve json file
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch %q: %v", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad http response: %s", res.Status)
	}

	// read http response
	// TODO: can possibly place ioutil.ReadAll() with the following line:
	//body, err := fs.ReadFile(http.FS(&http.File{File: fs.FileBase{Dir: ".", Name: "file.json"}}), "file.json")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// unmarshal json into parsable data
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %v", err)
	}

	return data, nil
}

func GetGpuId(vendor string, name string, notebook bool) (string, error) {
	var gpuId string
	var err error

	switch vendor {
	case "NVIDIA":
		gpuId, err = NvidiaGpuId(name, notebook)
		if err != nil {
			return "", err
		}
	}

	return gpuId, nil
}
