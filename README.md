# Introduction

A Go implementation of the Futaba S.Bus protocol, AKA SBUS.

Package sbus offers marshalling (serialization) and unmarshalling of the Futaba S.Bus digital servo serial protocol

# SBUS Protocol
The protocol is used for controlling digital [servos](https://en.wikipedia.org/wiki/Servo_\(radio_control\)) for building hobby projects, radio-controlled airplanes, [cameras](https://www.blackmagicdesign.com/products/blackmagicmicrostudiocamera4k/customization), robots, quadcopters, and drones.

# Examples

```go
ExampleFrame() {
	// Create Frame with some data
	Frame := Frame{
		Ch: Channels{
			0x000, // Minumum channel value?
			0x000,
			0x7ff, // Maxiumum channel value
			0x000,
			0x000,
			0x000,
			0x000,
			0x400, // 0% channel value
			0x000,
			0x000,
			0x720, // 100% positive channel value
			0x000,
			0x000,
			0x0ff, // -100% channel value?
			0x000,
			0x000,
		},
	}

	// Update Frame data
	Frame.Flags = Flags{Ch17: true, Failsafe: true}

	// Marshal a Frame
	fmt.Printf("%b\n", Frame.Marshal())
	// Output: [1111 0 0 11000000 11111111 1 0 0 0 0 0 10000000 0 0 0 11001000 1 0 10000000 1111111 0 0 0 10010000 0]
}
```

```go
func ExampleUnmarshalFrame() {
	// Create data
	data := [25]byte{0xf, 0xff, 0x7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10, 0}

	frame, err := UnmarshalFrame(data)
	if err != nil {
		panic(err)
	}
	// Marshal a Frame
	fmt.Println(frame)
	// Output: {[2047 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] {false false false true}}
}
```

# Suggested reading
See [this link dump](https://gist.github.com/johnelliott/3eca91e13afa354f6687de698e06ccc6) for more.

If Go isn't the right language, there are many open-source implementations of the protocol in C++ and for Arduino: <https://duckduckgo.com/?q=(parse+OR+decode)+SBUS]>
