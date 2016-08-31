package my_state

var State struct {
  Help       *bool
  Host       *string
  Port       *int
  Timeout    *int
  Url        *string
  Sensor     *string
  AllSensors *bool
  Debug      *bool
  Body       []byte
}
