package config

import "os"

func AdminEmail() string {
	return os.Getenv("ADMIN_EMAIL")
}

func AdminPassword() string {
	return os.Getenv("ADMIN_PASSWORD")
}

func MasterEmail() string {
	return os.Getenv("MASTER_EMAIL")
}

func MasterPassword() string {
	return os.Getenv("MASTER_PASSWORD")
}
