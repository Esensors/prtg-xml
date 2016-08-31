package parse_xml

import (
  "encoding/xml"
  "github.com/esensors/prtg-xml/esensors-websensor-prtg/my_state"
  "strings"
  //    "fmt"
)

type SWSensors struct {
  Sht   string `xml:"sht"`
  Tm0   string `xml:"tm0"`
  Tun0  string `xml:"tun0"`
  Hu0   string `xml:"hu0"`
  Ilum  string `xml:"ilum"`
  Il0   string `xml:"il0"`
  Evin  string `xml:"evin"`
  Vin   string `xml:"vin"`
  Ethm  string `xml:"ethm"`
  Thm   string `xml:"thm"`
  Ecin  string `xml:"ecin"`
  Cin   string `xml:"cin"`
  Efld  string `xml:"efld"`
  Fin   string `xml:"fin"`
  Alarm string `xml:"stu0"`
}

// Parses output of Websensor and produces PRTG XML
func Parse() string {
  var s SWSensors
  xml.Unmarshal(my_state.State.Body, &s)

  var xml_out string
  xml_out = `<?xml version="1.0" encoding="UTF-8" ?>
<prtg>`

  if *my_state.State.AllSensors || *my_state.State.Sensor == "temperature" {
    xml_out += `
    <result>
        <Channel>Temperature</Channel>
        <Float>1</Float>
        <Value>` + s.Tm0 + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "humidity" {
    xml_out = xml_out + `
    <result>
        <Channel>Humidity</Channel>
        <Float>1</Float>
        <Value>` + s.Hu0 + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "illumination" {
    xml_out = xml_out + `
    <result>
        <Channel>Illumination</Channel>
        <Float>1</Float>
        <Value>` + s.Il0 + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "voltage" {
    xml_out = xml_out + `
    <result>
        <Channel>Voltage</Channel>
        <Float>1</Float>
        <Value>` + s.Vin + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "thermistor" {
    xml_out = xml_out + `
    <result>
        <Channel>Thermistor</Channel>
        <Float>1</Float>
        <Value>` + s.Thm + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "contact" {
    xml_out = xml_out + `
    <result>
        <Channel>Contact</Channel>
        <Value>` + s.Cin + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "flood" {
    xml_out = xml_out + `
    <result>
        <Channel>Flood</Channel>
        <Value>` + s.Fin + `</Value>
    </result>`
  }
  if *my_state.State.AllSensors || *my_state.State.Sensor == "alarm" {
    var v string
    if strings.ToLower(s.Alarm) == "ok" {
      v = "0"
    } else {
      v = "1"
    }

    xml_out = xml_out + `
    <result>
        <Channel>Alarm</Channel>
        <Value>` + v + `</Value>
    </result>`
  }
  xml_out = xml_out + `
</prtg>
`

  return xml_out
}
