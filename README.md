# Esensors websensor XML plugin for PRTG Network Monitor

* [Configuring PRTG Network Monitor](#configuring-prtg-network-monitor)
  * [Download](#download)
  * [Copy files](#copy-files)
  * [Adding device](#adding-device)
* [Investigating issues](#investigating-issues)
  * [XML: The returned xml does not match the expected schema. (code: PE233)](#xml-the-returned-xml-does-not-match-the-expected-schema-code-pe233)

## Configuring PRTG Network Monitor

### Download
To install Esensors websensor XML plugin you need to download
from the [releases](https://github.com/Esensors/prtg-xml/releases) section:
* PRTG device template file (Esensors-Websensor-XML.odt, comes in the Source code)
* plugin binary (choose correct binary according to your architecture)

**Make sure to download files from the latest release.**

### Copy files
* copy `Esensors-Websensor-XML.odt` to `C:\Program Files (x86)\PRTG Network Monitor\devicetemplates`
* copy `esensors-websensor-prtg.exe` to `C:\Program Files (x86)\PRTG Network Monitor\Custom Sensors\EXEXML`

### Adding device
* Start PRTG Enterprise Console
* Right-click Local Probe and choose "Add Device"
* Choose suitable device name (e.g. Websensor)
* In the "IPv4 Address/DNS Name enter HOST:PORT/URL (without http),
  e.g. `24.39.65.206:4248/status.xml`
* Choose "Automatic sensor creation using specific device template in "Sensor Management"
* Choose "Esensors-Websensor-XML" device template


## Investigating issues

### XML: The returned xml does not match the expected schema. (code: PE233)

To investigate this issue please:
* click the failing sensor
* choose Settings button
* in the "SENSOR SETTINGS" section, item "EXE Result" mark "Write EXE result to disk"
* let the sensor run (wait for the period of execution, "Scanning Interval" on the same screen)
* collect all the files from `%programdata%\Paessler\PRTG Network Monitor\Logs (Sensors)`
* send to support for review
