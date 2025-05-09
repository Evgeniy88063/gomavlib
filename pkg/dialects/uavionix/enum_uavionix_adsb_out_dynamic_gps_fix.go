//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strconv"
)

// Status for ADS-B transponder dynamic input
type UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX uint64

const (
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0 UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 0
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1 UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 1
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D     UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 2
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D     UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 3
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS   UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 4
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK    UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 5
)

var value_to_label_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = map[UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX]string{
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0: "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1: "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D:     "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D:     "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS:   "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK:    "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK",
}

var label_to_value_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = map[string]UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX{
	"UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0": UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0,
	"UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1": UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1,
	"UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D":     UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D,
	"UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D":     UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D,
	"UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS":   UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS,
	"UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK":    UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
