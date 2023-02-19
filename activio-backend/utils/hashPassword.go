package utils

const SECRET = "activio" // Get this from an environment variable

func HashPassword(password string) string {
	hashedPassword := StringHash(password + SECRET)
	return hashedPassword
}