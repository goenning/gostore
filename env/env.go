package env

import "os"

//GetEnvOrDefault returns environment value if it's available, otherwise returns given default value
func GetEnvOrDefault(env string, def string) string {
	v := os.Getenv(env)
	if v != "" {
		return v
	}
	return def
}
