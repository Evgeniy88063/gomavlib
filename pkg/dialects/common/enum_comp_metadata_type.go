//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"fmt"
	"strconv"
)

// Supported component metadata types. These are used in the "general" metadata file returned by COMPONENT_METADATA to provide information about supported metadata types. The types are not used directly in MAVLink messages.
type COMP_METADATA_TYPE uint64

const (
	// General information about the component. General metadata includes information about other metadata types supported by the component. Files of this type must be supported, and must be downloadable from vehicle using a MAVLink FTP URI.
	COMP_METADATA_TYPE_GENERAL COMP_METADATA_TYPE = 0
	// Parameter meta data.
	COMP_METADATA_TYPE_PARAMETER COMP_METADATA_TYPE = 1
	// Meta data that specifies which commands and command parameters the vehicle supports. (WIP)
	COMP_METADATA_TYPE_COMMANDS COMP_METADATA_TYPE = 2
	// Meta data that specifies external non-MAVLink peripherals.
	COMP_METADATA_TYPE_PERIPHERALS COMP_METADATA_TYPE = 3
	// Meta data for the events interface.
	COMP_METADATA_TYPE_EVENTS COMP_METADATA_TYPE = 4
	// Meta data for actuator configuration (motors, servos and vehicle geometry) and testing.
	COMP_METADATA_TYPE_ACTUATORS COMP_METADATA_TYPE = 5
)

var value_to_label_COMP_METADATA_TYPE = map[COMP_METADATA_TYPE]string{
	COMP_METADATA_TYPE_GENERAL:     "COMP_METADATA_TYPE_GENERAL",
	COMP_METADATA_TYPE_PARAMETER:   "COMP_METADATA_TYPE_PARAMETER",
	COMP_METADATA_TYPE_COMMANDS:    "COMP_METADATA_TYPE_COMMANDS",
	COMP_METADATA_TYPE_PERIPHERALS: "COMP_METADATA_TYPE_PERIPHERALS",
	COMP_METADATA_TYPE_EVENTS:      "COMP_METADATA_TYPE_EVENTS",
	COMP_METADATA_TYPE_ACTUATORS:   "COMP_METADATA_TYPE_ACTUATORS",
}

var label_to_value_COMP_METADATA_TYPE = map[string]COMP_METADATA_TYPE{
	"COMP_METADATA_TYPE_GENERAL":     COMP_METADATA_TYPE_GENERAL,
	"COMP_METADATA_TYPE_PARAMETER":   COMP_METADATA_TYPE_PARAMETER,
	"COMP_METADATA_TYPE_COMMANDS":    COMP_METADATA_TYPE_COMMANDS,
	"COMP_METADATA_TYPE_PERIPHERALS": COMP_METADATA_TYPE_PERIPHERALS,
	"COMP_METADATA_TYPE_EVENTS":      COMP_METADATA_TYPE_EVENTS,
	"COMP_METADATA_TYPE_ACTUATORS":   COMP_METADATA_TYPE_ACTUATORS,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e COMP_METADATA_TYPE) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_COMP_METADATA_TYPE[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *COMP_METADATA_TYPE) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_COMP_METADATA_TYPE[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = COMP_METADATA_TYPE(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e COMP_METADATA_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
