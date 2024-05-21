package configparser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	OpenNebula struct {
		Endpoint string
		Username string
		Password string
		Port     int
		SSL      bool
	}
	Exporter struct {
		Listen          string
		Port            int
		IntervalSeconds int
	}
}

func newConfig(endpoint, username, password string, port int, ssl bool, listen string, listenPort, interval int) *Config {
	var config Config
	config.OpenNebula.Endpoint = endpoint
	config.OpenNebula.Username = username
	config.OpenNebula.Password = password
	config.OpenNebula.Port = port
	config.OpenNebula.SSL = ssl
	config.Exporter.Listen = listen
	config.Exporter.Port = listenPort
	config.Exporter.IntervalSeconds = interval

	return &config
}

func GenarateConfig() error {
	const tmpFile = "/tmp/sample.json"
	conf := newConfig("127.0.0.1", "oneadmin", "yourpasswordhere", 2633, false, "0.0.0.0", 9200, 10)
	jsondata, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsondata)
	if err != nil {
		return err
	} else {
		fmt.Printf("Sample configuration has been created in: %v\n", tmpFile)
	}

	return nil
}

func ParseConfig(filePath string) (endpoint, username, password string, port int, ssl bool, listenOn string, listenPort int, interval int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	return config.OpenNebula.Endpoint, config.OpenNebula.Username, config.OpenNebula.Password, config.OpenNebula.Port, config.OpenNebula.SSL, config.Exporter.Listen, config.Exporter.Port, config.Exporter.IntervalSeconds

}

func CheckConfigFile(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return false
		}
	} else {
		return true
	}
}
