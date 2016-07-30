# Esensors websensor XML plugin for PRTG Network Monitor

## Configuring PRTG Network Monitor

### Download
To install Esensors websensor XML plugin you need to download
from the [releases](https://github.com/Esensors/prtg-xml/releases) section:
* PRTG device template file (Esensors-Websensor-XML.odt, comes in the Source code)
* plugin binary (choose correct binary according to your architecture)

### Copy files
Copy `Esensors-Websensor-XML.odt` to `C:\Program Files (x86)\PRTG Network Monitor\devicetemplates`
and `esensors-websensor-prtg.exe` to `C:\Program Files (x86)\PRTG Network Monitor\Custom Sensors\EXEXML`

### Adding device
* Start PRTG Enterprise Console
* Right-click Local Probe and choose "Add Device"
* Choose suitable device name (e.g. Websensor)
* In the "IPv4 Address/DNS Name enter HOST:PORT/URL (without http),
  e.g. `24.39.65.206:4248/status.xml`
* Choose "Automatic sensor creation using specific device template in "Sensor Management"
* Choose "Esensors-Websensor-XML" device template
