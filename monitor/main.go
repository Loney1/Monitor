package main

import (
	"os"

	"monitor/server"

	logger "github.com/sirupsen/logrus"

)


func main() {
	confPath := os.Getenv("API_SRV_CONF_PATH")

	if confPath == "" {
		// dev mode, configure path load from file
		logger.Info("load configure from file")
		confPath = "./monitor.yaml"
	}
	logger.Infof("load configure from %s", confPath)

	server.InputArg()

	server.GetJenkinsData("GET", "http://jenkins.zawx.com/job/" + server.URL + "/lastBuild/api/json", "")

}


