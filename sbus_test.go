package sbus

import (
	"fmt"
	"testing"
)

func TestFrame(t *testing.T) {
	// Basic
	blankFrame := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if result := (Frame{}).Marshal(); result != blankFrame {
		t.Fatalf("Failed to marshal blank Frame, blankFrame %b, got %b", blankFrame, result)
	}

	// Flags for RC
	failsafe := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10, 0}
	if result := (Frame{Flags: Flags{Failsafe: true}}).Marshal(); result != failsafe {
		t.Fatalf("Failed to marshal Frame with failsafe, expected %b, got %b", failsafe, result)
	}

	framelost := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x20, 0}
	if result := (Frame{Flags: Flags{Framelost: true}}).Marshal(); result != framelost {
		t.Fatalf("Failed to marshal Frame with failsafe, expected %b, got %b", failsafe, result)
	}

	// Digital Channels
	ch17 := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80, 0}
	if result := (Frame{Flags: Flags{Ch17: true}}).Marshal(); result != ch17 {
		t.Fatalf("Failed to marshal Frame with failsafe, expected %b, got %b", failsafe, result)
	}

	ch18 := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x40, 0}
	if result := (Frame{Flags: Flags{Ch18: true}}).Marshal(); result != ch18 {
		t.Fatalf("Failed to marshal Frame with failsafe, expected %b, got %b", failsafe, result)
	}
}

// TODO test String methods also

// TODO test that the bit packing happens into proper channels, middle channel values, etc.
// Channel values

// TODO check all this against SBUS meter and change other code to suit, updating tests

func ExampleFrame() {
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

func TestUnmarshalFrame(t *testing.T) {
	blankFrame := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	unmarshalled, _ := UnmarshalFrame(blankFrame)
	if unmarshalled.Ch[0] != 0 {
		t.Fatal("Channel 0 is non-zero", unmarshalled.Ch)
	}
	if unmarshalled.Ch != [16]uint16{} {
		t.Fatal("Channels are not all zero")
	}
	if unmarshalled.Flags.Failsafe != false {
		t.Fatal("Failsafe is not false")
	}

	// Flags
	failsafe := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10, 0}
	unmarshalled, _ = UnmarshalFrame(failsafe)
	if unmarshalled.Flags.Failsafe != true {
		t.Fatal("Failsafe is not true")
	}
	framelost := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x20, 0}
	unmarshalled, _ = UnmarshalFrame(framelost)
	if unmarshalled.Flags.Framelost != true {
		t.Fatal("framelost is not true")
	}

	ch18 := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x40, 0}
	unmarshalled, _ = UnmarshalFrame(ch18)
	if unmarshalled.Flags.Ch18 != true {
		t.Fatal("ch18 is not true")
	}
	ch17 := [25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80, 0}
	unmarshalled, _ = UnmarshalFrame(ch17)
	if unmarshalled.Flags.Ch17 != true {
		t.Fatal("ch17 is not true")
	}

	// Channels
	_, err := UnmarshalFrame([25]byte{0x5})
	if err == nil {
		t.Fatal("Did not catch bad start byte")
	}
	_, err = UnmarshalFrame([25]byte{0xf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x44})
	if err == nil {
		t.Fatal("Did not catch bad end byte")
	}

	// Proportional channels
	ch1mid := [25]byte{0xf, 0xff, 0x7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10, 0}
	unmarshalled, err = UnmarshalFrame(ch1mid)
	if err != nil {
		t.Fatal(err)
	}
	if unmarshalled.Ch[0] != 0x7ff {
		t.Fatalf("Channel 0 is not at center value, channels are: %b", unmarshalled.Ch)
	}
}

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
