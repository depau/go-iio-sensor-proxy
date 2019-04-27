package sensorproxy

import "github.com/godbus/dbus"

type sensorProxy struct {
	busObject dbus.BusObject
}

type SensorProxy interface {
	HasAccelerometer() (bool, error)
	HasAmbientLight() (bool, error)

	ClaimAccelerometer() error
	ReleaseAccelerometer() error
	ClaimAmbientLight() error
	ReleaseAmbientLight() error

	GetAccelerometerOrientation() (string, error)
	GetLightLevel() (float64, error)
	GetLightLevelUnit() (string, error)
}

func NewSensorProxyFromBus(systemBus *dbus.Conn) (SensorProxy, error) {
	s := new(sensorProxy)
	s.busObject = systemBus.Object("net.hadess.sensorProxy", "/net/hadess/sensorProxy")
	return s, nil
}

func (s *sensorProxy) getStringProperty(property string) (value string, err error) {
	var variant dbus.Variant
	variant, err = s.busObject.GetProperty(property)
	if err == nil {
		value = variant.Value().(string)
	}
	return
}

func (s *sensorProxy) getBoolProperty(property string) (value bool, err error) {
	var variant dbus.Variant
	variant, err = s.busObject.GetProperty(property)
	if err == nil {
		value = variant.Value().(bool)
	}
	return
}

func (s *sensorProxy) getFloat64Property(property string) (value float64, err error) {
	var variant dbus.Variant
	variant, err = s.busObject.GetProperty(property)
	if err == nil {
		value = variant.Value().(float64)
	}
	return
}

func (s *sensorProxy) HasAccelerometer() (hasAccelerometer bool, err error) {
	return s.getBoolProperty("net.hadess.sensorProxy.HasAccelerometer")
}

func (s *sensorProxy) HasAmbientLight() (hasAmbientLight bool, err error) {
	return s.getBoolProperty("net.hadess.sensorProxy.HasAmbientLight")
}

func (s *sensorProxy) ClaimAccelerometer() error {
	return s.busObject.Call("net.hadess.sensorProxy.ClaimAccelerometer", 0).Err
}

func (s *sensorProxy) ClaimAmbientLight() error {
	return s.busObject.Call("net.hadess.sensorProxy.ClaimLight", 0).Err
}

func (s *sensorProxy) ReleaseAccelerometer() error {
	return s.busObject.Call("net.hadess.sensorProxy.ReleaseAccelerometer", 0).Err
}

func (s *sensorProxy) ReleaseAmbientLight() error {
	return s.busObject.Call("net.hadess.sensorProxy.ReleaseLight", 0).Err
}

func (s *sensorProxy) GetAccelerometerOrientation() (accelOrientation string, err error) {
	return s.getStringProperty("net.hadess.sensorProxy.AccelerometerOrientation")
}

func (s *sensorProxy) GetLightLevel() (lightLevel float64, err error) {
	return s.getFloat64Property("net.hadess.sensorProxy,LightLevel")
}

func (s *sensorProxy) GetLightLevelUnit() (lightLevelUnit string, err error) {
	return s.getStringProperty("net.hadess.sensorProxy.LightLevelUnit")
}
