# Introduction
The Futaba S-BUS protocol is a serial protocol to control servos. Up to 16 proportional and two digital channels are available. The protocol is derived from the very known RS232 protocol used everywhere. The signal must be first inverted. The frame is 8E2.

This board provide a complete electrical separation of RC gear and mbed controller. The S-BUS serial signal is converted and isolated by optocoppler. To keep control of the plane/car/ship an additional controller (ATTINY13) is on board. It monitors a standard servo signal (Master) and sends the serial S-BUS signal either to the mbed or directly to the servos. A special 'middle' position sends the S-BUS data to the servos and the mbed. In this mode the mbed can check stick positions and range without sending data to servos. Because the S-BUS data contains the Master signal, mbed knows about this mode.

# S-BUS protocol
The protocol is 25 Byte long and is send every 14ms (analog mode) or 7ms (highspeed mode).
One Byte = 1 startbit + 8 databit + 1 paritybit + 2 stopbit (8E2), baudrate = 100'000 bit/s
The highest bit is send first. The logic is inverted (Level High = 1)

```
[startbyte] [data1] [data2] .... [data22] [flags][endbyte]

startbyte = 11110000b (0xF0)

data 1-22 = [ch1, 11bit][ch2, 11bit] .... [ch16, 11bit] (ch# = 0 bis 2047)
channel 1 uses 8 bits from data1 and 3 bits from data2
channel 2 uses last 5 bits from data2 and 6 bits from data3
etc.

flags:
bit7 = ch17 = digital channel (0x80)
bit6 = ch18 = digital channel (0x40)
bit5 = Frame lost, equivalent red LED on receiver (0x20)
bit4 = failsafe activated (0x10)
bit3 = n/a
bit2 = n/a
bit1 = n/a
bit0 = n/a

endbyte = 00000000b
```

https://quadmeup.com/generate-s-bus-with-arduino-in-a-simple-way/ the iNav developer
https://github.com/bolderflight/SBUS/blob/c205c2941ad6f9433afc0540492f4db9cc96e53b/src/SBUS.cpp#L147

Frame represents an SBUS data frame
https://github.com/SvenKratz/TTY-SBUS

# tips
- use `godoc -http=localhost:8080` to develop with docs

# TODO
+ repo called go-sbus
+ package name sbus
- make Flags have decode and encode methods
- add decoding
- tests
- move types up like what Jaana said
- example code
- make import path without /pkg/
- add readme with rationale and history, warning
- add citations to readme

# publishing checklist
- func names don't stutter
- review what docs look like
- revewi Jaana Dogan blog for tips
