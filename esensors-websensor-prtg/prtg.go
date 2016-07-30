package main

import (
    "encoding/xml"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

var help = flag.Bool("help", false, "specify to get help about the plugin")
var host = flag.String("host", "", "hostname or ip address of websensor device")
var port = flag.Int("port", 80, "tcp port of the sensor, defaults to 80")
var timeout = flag.Int("timeout", 15, "timeout for http request to the sensor, defaults to 15")
var url = flag.String("url", "ssetings.xml", "url to query for sensor data, defaults to ssetings.xml")
var sensor = flag.String("sensor", "", "name of the sensor to query")
var allSensors = flag.Bool("all-sensors", false, "output all sensors as channels")
var debug = flag.Bool("debug", false, "")

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
}

func debug_print(msg string) {
    if *debug == true {
        fmt.Println(msg)
    }
}

func main() {
    flag.Parse()

    if (*help == true || *host == "") {
        fmt.Println(`
PRTG plugin for Esensors websensor device (XML based only)

Syntax:
    esensors-websensor-prtg.exe --host <NAME> --sensor <NAME> [options]

Mandatory parameters:
  --host <NAME>
    address of device on network (name or IP).
  --sensor <NAME>
    one of: temperature, humidity, illumination,
            voltage, thermistor, contact, flood

Optional parameters:
    --port <N>, defaults to 80
    --timeout <M>, defaults to 15
    --url <URL>, defaults to ssetings.xml
    --debug, output a bit more information

Special modes:
    --all-sensors
      outputs information for all of the sensors as channels,
      could be used instead of --sensor parameter
    --host <HOST:PORT/URL>
      host parameter allows specification of almost full device URL
      as one parameter (very useful to minimize configuration efforts)

Example:
    esensors-websensor-prtg.exe --host 24.39.65.206:4248/status.xml --sensor temperature

`)
        os.Exit(0)
    }

    var full_url = *host
    if strings.Contains(*host, ":") == false {
        full_url = full_url + ":" + strconv.Itoa(*port)
    }
    if strings.Contains(*host, "/") == false {
        full_url = full_url + "/" + *url
    }

    full_url = "http://" + full_url
    debug_print("initiating HTTP request to " + full_url)

    client := &http.Client {
        Timeout: time.Duration(*timeout) * time.Second,
    }
    resp, err := client.Get(full_url)
    if err != nil {
        log.Fatal("Error connecting to sensor:", err)
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading from sensor:", err)
    }

    debug_print("got response from the sensor:" + string(body))

    var s SWSensors
    xml.Unmarshal(body, &s)

    var xml_out string
    xml_out = `<?xml version="1.0" encoding="UTF-8" ?>
<prtg>`

    if *allSensors || *sensor == "temperature" {
        xml_out += `
    <result>
        <Channel>Temperature</Channel>
        <Float>1</Float>
        <Value>` + s.Tm0 + `</Value>
    </result>`
    }
    if *allSensors || *sensor == "humidity" {
        xml_out = xml_out + `
    <result>
        <Channel>Humidity</Channel>
        <Float>1</Float>
        <Value>` + s.Hu0 + `</Value>
    </result>`
    }
    if *allSensors || *sensor == "illumination" {
        xml_out = xml_out + `
    <result>
        <Channel>Illumination</Channel>
        <Float>1</Float>
        <Value>` + s.Il0 + `</Value>
    </result>`
    }
    if *allSensors || *sensor == "voltage" {
        xml_out = xml_out + `
    <result>
        <Channel>Voltage</Channel>
        <Float>1</Float>
        <Value>` + s.Vin + `</Value>
    </result>`
    }
    if *allSensors || *sensor == "thermistor" {
        xml_out = xml_out + `
    <result>
        <Channel>Thermistor</Channel>
        <Float>1</Float>
        <Value>` + s.Thm + `</Value>
    </result>`
    }
    if *allSensors || *sensor == "contact" {
        xml_out = xml_out + `
    <result>
        <Channel>Contact</Channel>
        <Value>` + s.Cin + `</Value>
    </result>`
    }
    if *allSensors || *sensor == "flood" {
        xml_out = xml_out + `
    <result>
        <Channel>Flood</Channel>
        <Value>` + s.Fin + `</Value>
    </result>`
    }
    xml_out = xml_out + `
</prtg>
`

    fmt.Println(xml_out)
    os.Exit(0)
}
