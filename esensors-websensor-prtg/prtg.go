package main

import (
  "flag"
  "fmt"
  "github.com/esensors/prtg-xml/esensors-websensor-prtg/my_state"
  "github.com/esensors/prtg-xml/esensors-websensor-prtg/parse_xml"
  "io/ioutil"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"
)

func prtg_error(msg string) {
  var xml_out string
  xml_out = `<?xml version="1.0" encoding="UTF-8" ?>
<prtg>
  <error>1</error>
  <text>
`
  xml_out = xml_out + msg

  xml_out = xml_out + `
  </text>  
</prtg>
`
  fmt.Println(xml_out)
  debug_print(msg)
  os.Exit(2)
}

func debug_print(msg string) {
  if *my_state.State.Debug == true {
    fmt.Println(msg)
  }
}

func main() {
  my_state.State.Help = flag.Bool("help", false, "specify to get help about the plugin")
  my_state.State.Host = flag.String("host", "", "hostname or ip address of websensor device")
  my_state.State.Port = flag.Int("port", 80, "tcp port of the sensor, defaults to 80")
  my_state.State.Timeout = flag.Int("timeout", 15, "timeout for http request to the sensor, defaults to 15")
  my_state.State.Url = flag.String("url", "ssetings.xml", "url to query for sensor data, defaults to ssetings.xml")
  my_state.State.Sensor = flag.String("sensor", "", "name of the sensor to query")
  my_state.State.AllSensors = flag.Bool("all-sensors", false, "output all sensors as channels")
  my_state.State.Debug = flag.Bool("debug", false, "")
  flag.Parse()

  if *my_state.State.Help == true || *my_state.State.Host == "" {
    fmt.Println(`
PRTG plugin for Esensors websensor device (XML based only)

Syntax:
    esensors-websensor-prtg.exe --host <NAME> --sensor <NAME> [options]

Mandatory parameters:
  --host <NAME>
    address of device on network (name or IP).
  --sensor <NAME>
    one of: temperature, humidity, illumination,
            voltage, thermistor, contact, flood, alarm

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

  var full_url = *my_state.State.Host
  if strings.Contains(*my_state.State.Host, ":") == false && *my_state.State.Port != 80 {
    full_url = full_url + ":" + strconv.Itoa(*my_state.State.Port)
  }
  if strings.Contains(*my_state.State.Host, "/") == false {
    full_url = full_url + "/" + *my_state.State.Url
  }

  full_url = "http://" + full_url
  debug_print("initiating HTTP request to " + full_url)

  client := &http.Client{
    Timeout: time.Duration(*my_state.State.Timeout) * time.Second,
  }
  resp, err := client.Get(full_url)
  if err != nil {
    prtg_error("Error connecting to sensor [" + full_url + "]: " + err.Error())
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    prtg_error("Error reading from sensor [" + full_url + "]: " + err.Error())
  }
  my_state.State.Body = body

  debug_print("got response from the sensor:" + string(my_state.State.Body))

  xml_out := parse_xml.Parse()

  fmt.Println(xml_out)
  os.Exit(0)
}
