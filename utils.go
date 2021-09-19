package jin

import (
	"jin/log"
	"os"
)

func assertPath(result bool, msg string) {
	if !result {
		panic(msg)
	}
}

// check the server port.
// If the server port isn't assigned, using the enviroment variable port.
// If the enviroment variable port if undefined, using :8080 by default.
// If the parameters is more than 2, the function will call panic.
func resolveAddr(addr []string) string {
	switch len(addr) {
	case 0:
		log.Warn("The Server port is not assigned")
		if port := os.Getenv("PORT"); port != "" {
			log.Infof("Enviroment variable port=\"%s\"\n", port)
			return ":" + port
		}
		log.Infof("Environment variable PORT is undefined. Using port :8080 by default\n")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

// Determines whether the specified file exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return nil != err || os.IsExist(err)
}
