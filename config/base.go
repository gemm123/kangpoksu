package config

import "os"

func AdminEmail() string {
	return os.Getenv("ADMIN_EMAIL")
}

func AdminPassword() string {
	return os.Getenv("ADMIN_PASSWORD")
}
