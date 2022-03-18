package github.com/legendary-cookie/gon2u

import "net/http"
import "encoding/json"
import "io/ioutil"

func NameToUUID(name string) (string, error) {
	resp, err := http.Get(apiurl + name)
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
	return v.UUID, nil
}
