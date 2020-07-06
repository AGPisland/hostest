package gowiki

import (
	"database/sql"
	"encoding/json"

	"bitbucket.com/agagro/database"
)

type Result struct {
	Farm_name                 string
	Farm_id                   int
	Client_id                 int
	DeviceId                  int
	Supplier_name             string
	Client_name               string
	Device_name               string
	Device_location           string
	Device_latitud_ag         float64
	Device_longitud_ag        float64
	Device_json_conection_api sql.NullString
	Device_json_parameters    sql.NullString
}
type EstructuraApoyoCalculos struct {
	EquipID            string
	Deviceparticle     string
	Usuario            string
	Token              string
	Timeexpirationtoke string
}

type EstructuraWebMain struct {
	IdEquip           string
	Sector            string
	User              string
	Status            string
	Side              string
	IntentoDeBodyAdd2 string
}

var result Result
var apoyo EstructuraApoyoCalculos
var estrucweb EstructuraWebMain

func Recuperadordeviceindb() EstructuraWebMain {
	database.Run()
	SQLAPI := `SELECT devices.id, clients.name, clients.id, farms.id, farms.farm_name, suppliers.supplier_name, devices.device_json_conection_api, devices.device_name, devices.device_latitud_ag, 
	DEVICES.device_longitud_ag, DEVICES.device_location, DEVICES.device_json_parameters FROM devices, farms, clients, suppliers WHERE
	 DEVICES.farm_id=FARMS.ID AND FARMS.CLIENT_ID = clients.ID AND DEVICES.supplier_id = suppliers.ID  AND  DEVICES.type_device='AUTOMATAD1';`
	rows, _ := database.DB.Raw(SQLAPI).Rows()
	defer rows.Close()
	//fmt.Println(rows.Next())

	for rows.Next() {
		rows.Scan(&result.DeviceId, &result.Client_name, &result.Client_id, &result.Farm_id, &result.Farm_name, &result.Supplier_name, &result.Device_json_conection_api, &result.Device_name, &result.Device_latitud_ag, &result.Device_longitud_ag, &result.Device_location, &result.Device_json_parameters)
		json.Unmarshal([]byte(result.Device_json_parameters.String), &apoyo)
		//formatjson, _ := json.MarshalIndent(respuesta, "", "	")
		//fmt.Printf("%+v\n", string(formatjson))
		//formatjson, _ = json.MarshalIndent(result, "", "	")
		//fmt.Printf("%+v\n", string(formatjson))
	}

	estrucweb.IdEquip = apoyo.EquipID
	estrucweb.Sector = result.Device_location
	estrucweb.User = apoyo.Usuario
	if len(estrucweb.IdEquip) > 0 {
		estrucweb.Status = "OK"
	} else {
		estrucweb.Status = "FAIL"
	}
	return estrucweb
}

func Getresult() Result {
	return result
}
func GetEstrucApoyo() EstructuraApoyoCalculos {
	return apoyo
}

func Getwebmain() EstructuraWebMain {
	return estrucweb
}
