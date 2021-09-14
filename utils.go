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

func Exists(path string) bool {
	_, err := os.Stat(path)
	return nil != err || os.IsExist(err)
}
