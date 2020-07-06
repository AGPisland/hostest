package gowiki

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Resultparticledev struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Variables struct {
		Gongs string `json:"Gongs"`
	} `json:"variables"`
	Functions              []string    `json:"functions"`
	Connected              bool        `json:"connected"`
	PlatformID             int         `json:"platform_id"`
	ProductID              int         `json:"product_id"`
	SystemFirmwareVersion  string      `json:"system_firmware_version"`
	Cellular               bool        `json:"cellular"`
	SerialNumber           string      `json:"serial_number"`
	LastIPAddress          string      `json:"last_ip_address"`
	LastHandshakeAt        time.Time   `json:"last_handshake_at"`
	LastHeard              time.Time   `json:"last_heard"`
	Notes                  interface{} `json:"notes"`
	FirmwareUpdatesEnabled bool        `json:"firmware_updates_enabled"`
	FirmwareUpdatesForced  bool        `json:"firmware_updates_forced"`
}

type Loginsesion struct {
	Id       float32
	Token    string
	Iddevice string
}

var constantegrupal Loginsesion

func InitApi(tk string, id string) {
	r := rand.New(rand.NewSource(99))
	constantegrupal.Id = r.Float32()
	constantegrupal.Token = tk
	constantegrupal.Iddevice = id
	fmt.Println("Grupo iniciado")
}

func Getonlinedevices() /*Resultparticledev*/ string {
	url := "https://api.particle.io/v1/devices/DEV?access_token=TK"
	url = strings.Replace(url, "DEV", constantegrupal.Iddevice, 100)
	url = strings.Replace(url, "TK", constantegrupal.Token, 100)
	fmt.Println(url)

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//var particleapi Resultparticledev
	//json.Unmarshal([]byte(body), particleapi)
	return string(body)
}
