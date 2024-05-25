package connector

import (
	"fmt"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
)

type Connector struct {
	User     string
	Password string
	Endpoint string
	Port     int
	SSL      bool
}

func NewConnector(username, password, endpoint string, port int, ssl bool) *Connector {
	var con Connector
	con.User = username
	con.Password = password
	con.Endpoint = endpoint
	con.Port = port
	con.SSL = ssl

	return &con

}

func Connect(connector *Connector) *goca.Controller {
	var endpoint string
	var schema string
	if connector.SSL {
		schema = "https"
	} else {
		schema = "http"
	}

	endpoint = fmt.Sprintf("%v://%v:%v/RPC2", schema, connector.Endpoint, connector.Port)
	client := goca.NewDefaultClient(
		goca.NewConfig(connector.User, connector.Password, endpoint),
	)
	controler := goca.NewController(client)
	return controler
}
