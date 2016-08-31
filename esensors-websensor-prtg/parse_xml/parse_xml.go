package parse_xml

import (
    "encoding/xml"
    "strings"
    "github.com/esensors/prtg-xml/esensors-websensor-prtg/my_state"
//    "fmt"
)

type SWSensors struct {
    Sht string `xml:"sht"`
    Tm0 string `xml:"tm0"`
    Tun0 string `xml:"tun0"`
    Hu0 string `xml:"hu0"`
    Ilum string `xml:"ilum"`
    Il0 string `xml:"il0"`
    Evin string `xml:"evin"`
    Vin string `xml:"vin"`
    Ethm string `xml:"ethm"`
    Thm string `xml:"thm"`
    Ecin string `xml:"ecin"`
    Cin string `xml:"cin"`
    Efld string `xml:"efld"`
    Fin string `xml:"fin"`
    Alarm string `xml:"stu0"`
}

// Parses output of Websensor and produces PRTG XML
func Parse(body []byte, st my_state.My_state) string {
    var s SWSensors
    xml.Unmarshal(body, &s)

    var xml_out string
    xml_out = `<?xml version="1.0" encoding="UTF-8" ?>
<prtg>`

    if *st.AllSensors || *st.Sensor == "temperature" {
        xml_out += `
    <result>
        <Channel>Temperature</Channel>
        <Float>1</Float>
        <Value>` + s.Tm0 + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "humidity" {
        xml_out = xml_out + `
    <result>
        <Channel>Humidity</Channel>
        <Float>1</Float>
        <Value>` + s.Hu0 + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "illumination" {
        xml_out = xml_out + `
    <result>
        <Channel>Illumination</Channel>
        <Float>1</Float>
        <Value>` + s.Il0 + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "voltage" {
        xml_out = xml_out + `
    <result>
        <Channel>Voltage</Channel>
        <Float>1</Float>
        <Value>` + s.Vin + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "thermistor" {
        xml_out = xml_out + `
    <result>
        <Channel>Thermistor</Channel>
        <Float>1</Float>
        <Value>` + s.Thm + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "contact" {
        xml_out = xml_out + `
    <result>
        <Channel>Contact</Channel>
        <Value>` + s.Cin + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "flood" {
        xml_out = xml_out + `
    <result>
        <Channel>Flood</Channel>
        <Value>` + s.Fin + `</Value>
    </result>`
    }
    if *st.AllSensors || *st.Sensor == "alarm" {
        var v string
        if (strings.ToLower(s.Alarm) == "ok") {
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
