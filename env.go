package goconfig

import "os"

func getFromEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
