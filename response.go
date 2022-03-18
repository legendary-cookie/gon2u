package github.com/legendary-cookie/gon2u

const apiurl string = "https://api.mojang.com/users/profiles/minecraft"

type MojangApiResponse struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}
