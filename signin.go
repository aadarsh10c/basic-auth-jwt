package main

// create a scecret key , to be used as a signature
var jwtKey = []byte("my_secret_key")

// create an in-memory map to store user credentials
var users = map[string]string{
	"user1":  "password1",
	"users2": "password2",
}

// create a struct that stores the credentials form the user
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struc that will be encoded to a jwt.
// We add jwt.RegisteredClaims as an embedded type , to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
