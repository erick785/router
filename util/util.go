package util

import (
	"net"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// GetLocalIP gets local IP
func GetLocalIP() string {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// IsDirExist determines whether the directory exists
func IsDirExist(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		return fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// MkDir make directory
func MkDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// IsFileExist determines whether the file exists
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// CallerInfo gets call location
func CallerInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	return "[" + filename + ":" + strconv.FormatInt(int64(line), 10) + "] "
}

// SysSignal checks exit signal
func SysSignal(function func()) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTSTP)
	for {
		select {
		case <-c:
			function()
		}
	}
}

// IsStrExist determines whether the string exists
func IsStrExist(str string, strs []string) bool {
	for _, v := range strs {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}
