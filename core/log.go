package core

import (
	"fmt"
	rotates "github.com/lestrrat/go-file-rotatelogs"
	obliging "github.com/op/go-logging"
	"io"
	"my/config"
	"my/global"
	"my/utils"
	"os"
	"strings"
	"time"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "my"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

func init() {
	c := global.GVA_CONFIG.Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := obliging.MustGetLogger(module)
	var backends []obliging.Backend
	backends = registerStdout(c, backends)
	backends = registerFile(c, backends)

	obliging.SetBackend(backends...)
	global.GVA_LOG = logger
}

func registerStdout(c config.Log, backends []obliging.Backend) []obliging.Backend {
	if c.Stdout != "" {
		level, err := obliging.LogLevel(c.Stdout)
		if err != nil {
			fmt.Println(err)
		}
		backends = append(backends, createBackend(os.Stdout, c, level))
	}

	return backends
}

func registerFile(c config.Log, backends []obliging.Backend) []obliging.Backend {
	if c.File != "" {
		if ok, _ := utils.PathExists(logDir); !ok {
			// directory not exist
			fmt.Println("create log directory")
			_ = os.Mkdir(logDir, os.ModePerm)
		}
		fileWriter, err := rotates.New(
			logDir+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
			// generate soft link, point to latest log file
			rotates.WithLinkName(logSoftLink),
			// maximum time to save log files
			rotates.WithMaxAge(7*24*time.Hour),
			// time period of log file switching
			rotates.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Println(err)
			return backends
		}
		level, err := obliging.LogLevel(c.File)
		if err != nil {
			fmt.Println(err)
		}
		backends = append(backends, createBackend(fileWriter, c, level))
	}

	return backends
}

func createBackend(w io.Writer, c config.Log, level obliging.Level) obliging.Backend {
	backend := obliging.NewLogBackend(w, c.Prefix, 0)
	stdoutWriter := false
	if w == os.Stdout {
		stdoutWriter = true
	}
	format := getLogFormatter(c, stdoutWriter)
	backendLeveled := obliging.AddModuleLevel(obliging.NewBackendFormatter(backend, format))
	backendLeveled.SetLevel(level, module)
	return backendLeveled
}

func getLogFormatter(c config.Log, stdoutWriter bool) obliging.Formatter {
	pattern := defaultFormatter
	if !stdoutWriter {
		// Color is only required for console output
		// Other writers don't need %{color} tag
		pattern = strings.Replace(pattern, "%{color:bold}", "", -1)
		pattern = strings.Replace(pattern, "%{color:reset}", "", -1)
	}
	if !c.LogFile {
		// Remove %{logfile} tag
		pattern = strings.Replace(pattern, "%{longfile}", "", -1)
	}
	return obliging.MustStringFormatter(pattern)
}
