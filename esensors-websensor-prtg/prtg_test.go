package main

import (
  "fmt"
  "github.com/esensors/prtg-xml/esensors-websensor-prtg/my_state"
  "github.com/esensors/prtg-xml/esensors-websensor-prtg/parse_xml"
)

func ExampleTemperature() {
  my_state.State.Body = []byte(`<sensorsSW><dvc>TBD</dvc><sht>inline</sht><ilum>inline</ilum><evin>inline</evin><ethm>inline</ethm>
<ecin>inline</ecin><efld>inline</efld><epir>none</epir><egas>none</egas><ght>1</ght><gsn>1</gsn><eDL>none</eDL><ecam>none</ecam><ephc>none</ephc>
<sid0>995805</sid0><stu0>Alarm</stu0><tm0>80.98</tm0><hu0>37.41</hu0><il0>73.85</il0><tun0>F</tun0><cin>1</cin><fin>1</fin><pin>0</pin>
<vin>0.00</vin><thm>77.87</thm><phcv>0.00</phcv></sensorsSW>`)

  my_state.State.AllSensors = new(bool)
  *my_state.State.AllSensors = false
  my_state.State.Sensor = new(string)
  *my_state.State.Sensor = "temperature"

  xml_out := parse_xml.Parse()
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

func ExampleHumidity() {
  my_state.State.Body = []byte(`<sensorsSW><dvc>TBD</dvc><sht>inline</sht><ilum>inline</ilum><evin>inline</evin><ethm>inline</ethm>
<ecin>inline</ecin><efld>inline</efld><epir>none</epir><egas>none</egas><ght>1</ght><gsn>1</gsn><eDL>none</eDL><ecam>none</ecam><ephc>none</ephc>
<sid0>995805</sid0><stu0>Alarm</stu0><tm0>80.98</tm0><hu0>37.41</hu0><il0>73.85</il0><tun0>F</tun0><cin>1</cin><fin>1</fin><pin>0</pin>
<vin>0.00</vin><thm>77.87</thm><phcv>0.00</phcv></sensorsSW>`)

  my_state.State.AllSensors = new(bool)
  *my_state.State.AllSensors = false
  my_state.State.Sensor = new(string)
  *my_state.State.Sensor = "humidity"

  xml_out := parse_xml.Parse()
  fmt.Println(xml_out)

  // Output:
  // <?xml version="1.0" encoding="UTF-8" ?>
  // <prtg>
  //     <result>
  //         <Channel>Humidity</Channel>
  //         <Float>1</Float>
  //         <Value>37.41</Value>
  //     </result>
  // </prtg>
}

func ExampleAlarm() {
  my_state.State.Body = []byte(`<sensorsSW><dvc>TBD</dvc><sht>inline</sht><ilum>inline</ilum><evin>inline</evin><ethm>inline</ethm>
<ecin>inline</ecin><efld>inline</efld><epir>none</epir><egas>none</egas><ght>1</ght><gsn>1</gsn><eDL>none</eDL><ecam>none</ecam><ephc>none</ephc>
<sid0>995805</sid0><stu0>Alarm</stu0><tm0>80.98</tm0><hu0>37.41</hu0><il0>73.85</il0><tun0>F</tun0><cin>1</cin><fin>1</fin><pin>0</pin>
<vin>0.00</vin><thm>77.87</thm><phcv>0.00</phcv></sensorsSW>`)

  my_state.State.AllSensors = new(bool)
  *my_state.State.AllSensors = false
  my_state.State.Sensor = new(string)
  *my_state.State.Sensor = "alarm"

  xml_out := parse_xml.Parse()
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

func ExampleAll() {
  my_state.State.Body = []byte(`<sensorsSW><dvc>TBD</dvc><sht>inline</sht><ilum>inline</ilum><evin>inline</evin><ethm>inline</ethm>
<ecin>inline</ecin><efld>inline</efld><epir>none</epir><egas>none</egas><ght>1</ght><gsn>1</gsn><eDL>none</eDL><ecam>none</ecam><ephc>none</ephc>
<sid0>995805</sid0><stu0>Alarm</stu0><tm0>80.98</tm0><hu0>37.41</hu0><il0>73.85</il0><tun0>F</tun0><cin>1</cin><fin>1</fin><pin>0</pin>
<vin>0.00</vin><thm>77.87</thm><phcv>0.00</phcv></sensorsSW>`)

  my_state.State.AllSensors = new(bool)
  *my_state.State.AllSensors = true

  xml_out := parse_xml.Parse()
  fmt.Println(xml_out)

  // Output:
  // <?xml version="1.0" encoding="UTF-8" ?>
  // <prtg>
  //     <result>
  //         <Channel>Temperature</Channel>
  //         <Float>1</Float>
  //         <Value>80.98</Value>
  //     </result>
  //     <result>
  //         <Channel>Humidity</Channel>
  //         <Float>1</Float>
  //         <Value>37.41</Value>
  //     </result>
  //     <result>
  //         <Channel>Illumination</Channel>
  //         <Float>1</Float>
  //         <Value>73.85</Value>
  //     </result>
  //     <result>
  //         <Channel>Voltage</Channel>
  //         <Float>1</Float>
  //         <Value>0.00</Value>
  //     </result>
  //     <result>
  //         <Channel>Thermistor</Channel>
  //         <Float>1</Float>
  //         <Value>77.87</Value>
  //     </result>
  //     <result>
  //         <Channel>Contact</Channel>
  //         <Value>1</Value>
  //     </result>
  //     <result>
  //         <Channel>Flood</Channel>
  //         <Value>1</Value>
  //     </result>
  //     <result>
  //         <Channel>Alarm</Channel>
  //         <Value>1</Value>
  //     </result>
  // </prtg>
}
