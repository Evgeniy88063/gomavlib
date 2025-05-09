//autogenerated:yes
//nolint:revive,govet,errcheck
package ardupilotmega

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnum_EKF_STATUS_FLAGS(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var e EKF_STATUS_FLAGS
		e.UnmarshalText([]byte{})
		e.MarshalText()
		e.String()
	})

	t.Run("first entry", func(t *testing.T) {
		enc, err := EKF_ATTITUDE.MarshalText()
		require.NoError(t, err)

		var dec EKF_STATUS_FLAGS
		err = dec.UnmarshalText(enc)
		require.NoError(t, err)

		require.Equal(t, EKF_ATTITUDE, dec)
	})
}
