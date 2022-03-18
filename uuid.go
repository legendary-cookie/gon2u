package gon2u

import "net/http"
import "encoding/json"
import "io/ioutil"

func UUIDToName(uuid string) (string, error) {
	resp, err := http.Get(uuidurl + uuid)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	v := &MojangApiResponse{}
	err = json.Unmarshal(data, v)
	if err != nil {
		return "", err
	}
	return v.Name, nil
}
