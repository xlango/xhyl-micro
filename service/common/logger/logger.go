package logger

import (
	"github.com/cihub/seelog"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"xhyl-micro/service/common/utils"
)

func init() {
	//解决windows无法debug启动的问题
	currentDir := utils.GetCurrentExeDir()
	if len(currentDir) > 0 {
		currentDir = currentDir + string(os.PathSeparator)
	}
	InitLogger(currentDir)
}

func LogDebug(v ...interface{}) {
	seelog.Debug(v)
}
func LogInfo(v ...interface{}) {
	seelog.Info(v)
}
func LogError(v ...interface{}) {
	seelog.Error(v)
}

func InitLogger(currentDir string) {

	config, err := ioutil.ReadFile(currentDir + "seelog.xml")
	if err != nil {
		log.Fatalln(err)
	}
	configStr := string(config)
	newConfigStr := strings.Replace(configStr, "filename=\"./", "filename=\""+currentDir, -1)
	Logger, err := seelog.LoggerFromConfigAsString(newConfigStr)
	if err != nil {
		log.Fatalln(err)
	}
	seelog.ReplaceLogger(Logger)
}
