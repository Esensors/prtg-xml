package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
    "github.com/esensors/prtg-xml/esensors-websensor-prtg/my_state"
    "github.com/esensors/prtg-xml/esensors-websensor-prtg/parse_xml"
)

var st my_state.My_state 

func debug_print(msg string) {
    if *st.Debug == true {
        fmt.Println(msg)
    }
}

func main() {
    st.Help = flag.Bool("help", false, "specify to get help about the plugin")
    st.Host = flag.String("host", "", "hostname or ip address of websensor device")
    st.Port = flag.Int("port", 80, "tcp port of the sensor, defaults to 80")
    st.Timeout = flag.Int("timeout", 15, "timeout for http request to the sensor, defaults to 15")
    st.Url = flag.String("url", "ssetings.xml", "url to query for sensor data, defaults to ssetings.xml")
    st.Sensor = flag.String("sensor", "", "name of the sensor to query")
    st.AllSensors = flag.Bool("all-sensors", false, "output all sensors as channels")
    st.Debug = flag.Bool("debug", false, "")
    flag.Parse()

    if (*st.Help == true || *st.Host == "") {
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

    var full_url = *st.Host
    if strings.Contains(*st.Host, ":") == false && *st.Port != 80 {
        full_url = full_url + ":" + strconv.Itoa(*st.Port)
    }
    if strings.Contains(*st.Host, "/") == false {
        full_url = full_url + "/" + *st.Url
    }

    full_url = "http://" + full_url
    debug_print("initiating HTTP request to " + full_url)

    client := &http.Client {
        Timeout: time.Duration(*st.Timeout) * time.Second,
    }
    resp, err := client.Get(full_url)
    if err != nil {
        log.Fatal("Error connecting to sensor [", full_url, "]: ", err)
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error reading from sensor [", full_url, "]: ", err)
    }

    debug_print("got response from the sensor:" + string(body))

    xml_out := parse_xml.Parse(body, st)

    fmt.Println(xml_out)
    os.Exit(0)
}
