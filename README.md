# Magic Home CLI + Lib

This repository offers the possibility to control Magic Home (Magic Hue) LED Strip Controller. The control can be done via terminal using the CLI or programmatically using the library. It is written in [Go](https://golang.org/) and already compiled for various systems and architectures like Linux, FreeBSD, macOS and Windows for amd64, i386, ARM. Take a look on the GitHub [releases](https://github.com/moonliightz/magic-home/releases) page for a list of available binaries.


## Features

- Change the state of the LED Strip Controller to on or off
- Change the color of the LED Strip Controller with RGBW
- Request the current state ot the LED Strip Controller
- Provides a simple ready-to-use CLI
- Can be used as a library to control the controller in your own applications

## CLI

### Install

Choose the archive matching the destination platform and extract it.

```bash
$ wget -qO- https://github.com/moonliightz/magic-home/releases/download/v1.1.0/magic-home_1.1.0_linux_x86_64.tar.gz | tar -zxvf - magic-home
```

> Binaries are available on the GitHub [releases](https://github.com/moonliightz/magic-home/releases) page.

### Usage

```bash
$ ./magic-home --help

NAME:
   magic-home - A CLI for controlling Magic Home (Magic Hue) LED Strip Controller

USAGE:
   magic-home [global options] command [command options] [arguments...]

COMMANDS:
   color, c     Set the color of the LED Strip
   state, s     Switch the LED Strip state to on or off
   status       Prints the status of the LED Strip
   discover, d  Discover for Magic Home devices on the network
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

- Discover devices
```bash
$ ./magic-home discover
```

- Turn LED Strip Controller on
```bash
$ ./magic-home state 192.168.1.42 on
```

- Turn LED Strip Controller off
```bash
$ ./magic-home state 192.168.1.42 off
```

- Switch Color to red with 100% brightness
```bash
$ ./magic-home color 192.168.1.42 255 0 0 0
```

- Switch Color to cyan with 100% brightness
```bash
$ ./magic-home color 192.168.1.42 0 255 255 0
```

- Request the status of a device
```bash
$ ./magic-home status 192.168.1.42

# You can also get the response as JSON
$ ./magic-home status --json 192.168.1.42
```

> Change `192.168.1.42` to the IP of your Controller.  
> If your controller only supports RGB instead of RGBW, just set the last value to 0.


## Library

### Install

```bash
$ go get -u github.com/moonliightz/magic-home
```

### Usage

```go
// Create a new Magic Home LED Strip Controller
controller, _ := magichome.New(net.ParseIP("192.168.1.42"), 5577)

// Turn LED Strip Controller on
controller.SetState(magichome.On)

// Tun LED Strip Controller off
controller.SetState(magichome.Off)

// Set color of LED Strip to white
controller.SetColor(magichome.Color{
  R: 255,
  G: 255,
  B: 255,
  W: 0,
})

// Get the current state of the device
// The response looks like
// {
//   "DeviceType": 51,
//   "State": 1,
//   "LedVersionNumber": 6,
//   "Mode": 97,
//   "Slowness": 15,
//   "Color": {
//     "R": 21,
//     "G": 77,
//     "B": 70,
//     "W": 0
//   }
// }
deviceState, _ := controller.GetDeviceState()

// Don't forget to close the connection to the LED Strip Controller
controller.Close()
```
For a ready-to-use example take a look at [examples/basic/main.go](examples/basic/main.go) or generally in the [examples](examples/) folder.


## License

magic-home is released under the [MIT license](LICENSE).
