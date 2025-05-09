//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strconv"
)

// GPS lataral offset encoding
type UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT uint64

const (
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_NO_DATA  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 0
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_2M  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 1
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_4M  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 2
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_6M  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 3
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_0M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 4
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_2M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 5
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_4M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 6
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_6M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 7
)

var value_to_label_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = map[UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT]string{
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_NO_DATA:  "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_NO_DATA",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_2M:  "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_2M",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_4M:  "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_4M",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_6M:  "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_6M",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_0M: "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_0M",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_2M: "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_2M",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_4M: "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_4M",
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_6M: "UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_6M",
}

var label_to_value_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = map[string]UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT{
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_NO_DATA":  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_NO_DATA,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_2M":  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_2M,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_4M":  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_4M,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_6M":  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_6M,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_0M": UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_0M,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_2M": UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_2M,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_4M": UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_4M,
	"UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_6M": UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_6M,
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT) MarshalText() ([]byte, error) {
	if name, ok := value_to_label_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT[e]; ok {
		return []byte(name), nil
	}
	return []byte(strconv.Itoa(int(e))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT) UnmarshalText(text []byte) error {
	if value, ok := label_to_value_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT[string(text)]; ok {
		*e = value
	} else if value, err := strconv.Atoi(string(text)); err == nil {
		*e = UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT(value)
	} else {
		return fmt.Errorf("invalid label '%s'", text)
	}
	return nil
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
