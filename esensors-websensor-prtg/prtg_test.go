package main

import (
    "fmt"
    "github.com/esensors/prtg-xml/esensors-websensor-prtg/my_state"
    "github.com/esensors/prtg-xml/esensors-websensor-prtg/parse_xml"
    )

func ExampleTemperature() {
    var st my_state.My_state 
    var body = []byte(`<sensorsSW><dvc>TBD</dvc><sht>inline</sht><ilum>inline</ilum><evin>inline</evin><ethm>inline</ethm>
<ecin>inline</ecin><efld>inline</efld><epir>none</epir><egas>none</egas><ght>1</ght><gsn>1</gsn><eDL>none</eDL><ecam>none</ecam><ephc>none</ephc>
<sid0>995805</sid0><stu0>Alarm</stu0><tm0>80.98</tm0><hu0>37.41</hu0><il0>73.85</il0><tun0>F</tun0><cin>1</cin><fin>1</fin><pin>0</pin>
<vin>0.00</vin><thm>77.87</thm><phcv>0.00</phcv></sensorsSW>`)
    
    st.AllSensors = new(bool)
    *st.AllSensors = false
    st.Sensor = new(string)
    *st.Sensor = "temperature"

    xml_out := parse_xml.Parse(body, st)
    fmt.Println(xml_out)

    // Output:
    // <?xml version="1.0" encoding="UTF-8" ?>
    // <prtg>
    //     <result>
    //         <Channel>Temperature</Channel>
    //         <Float>1</Float>
    //         <Value>80.98</Value>
    //     </result>
    // </prtg>
}

func ExampleAlarm() {
    var st my_state.My_state 
    var body = []byte(`<sensorsSW><dvc>TBD</dvc><sht>inline</sht><ilum>inline</ilum><evin>inline</evin><ethm>inline</ethm>
<ecin>inline</ecin><efld>inline</efld><epir>none</epir><egas>none</egas><ght>1</ght><gsn>1</gsn><eDL>none</eDL><ecam>none</ecam><ephc>none</ephc>
<sid0>995805</sid0><stu0>Alarm</stu0><tm0>80.98</tm0><hu0>37.41</hu0><il0>73.85</il0><tun0>F</tun0><cin>1</cin><fin>1</fin><pin>0</pin>
<vin>0.00</vin><thm>77.87</thm><phcv>0.00</phcv></sensorsSW>`)

    st.AllSensors = new(bool)
    *st.AllSensors = false
    st.Sensor = new(string)
    *st.Sensor = "alarm"

    xml_out := parse_xml.Parse(body, st)
    fmt.Println(xml_out)

    // Output:
    // <?xml version="1.0" encoding="UTF-8" ?>
    // <prtg>
    //     <result>
    //         <Channel>Alarm</Channel>
    //         <Value>1</Value>
    //     </result>
    // </prtg>
}
