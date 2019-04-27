# iio-sensor-proxy - Go bindings

[iio-sensor-proxy](https://github.com/hadess/iio-sensor-proxy) bindings for Go, based on [godbus](https://github.com/godbus/dbus/)

### Note

This is my first Go project. It might not follow best practices, etc. If so, please **do** let me know ;)

## Usage

```
$ go get github.com/Depau/go-iio-sensor-proxy
```

Get a connection to the system bus:

```go
conn, err := dbus.SystemBus()
if err != nil {
    log.Fatal(err)
}
```

Get an instance of the sensor proxy:

```go
sensorProxy, err := NewSensorProxyFromBus(conn)
if err != nil {
    log.Fatal(err)
}
```

Then use it:

```go
err := sensorProxy.ClaimAccelerometer()
// [handle error]
orientation, err := sensorProxy.GetAccelerometerOrientation()
// ...
```