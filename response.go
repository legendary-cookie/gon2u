package gon2u

const nameurl string = "https://api.mojang.com/users/profiles/minecraft/"
const uuidurl string = "https://api.mojang.com/user/profile/"

type MojangApiResponse struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}
