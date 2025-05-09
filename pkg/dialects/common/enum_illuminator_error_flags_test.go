//autogenerated:yes
//nolint:revive,govet,errcheck
package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnum_ILLUMINATOR_ERROR_FLAGS(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var e ILLUMINATOR_ERROR_FLAGS
		e.UnmarshalText([]byte{})
		e.MarshalText()
		e.String()
	})

	t.Run("first entry", func(t *testing.T) {
		enc, err := ILLUMINATOR_ERROR_FLAGS_THERMAL_THROTTLING.MarshalText()
		require.NoError(t, err)

		var dec ILLUMINATOR_ERROR_FLAGS
		err = dec.UnmarshalText(enc)
		require.NoError(t, err)

		require.Equal(t, ILLUMINATOR_ERROR_FLAGS_THERMAL_THROTTLING, dec)
	})
}
