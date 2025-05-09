//autogenerated:yes
//nolint:revive,govet,errcheck
package ardupilotmega

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnum_LIMITS_STATE(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var e LIMITS_STATE
		e.UnmarshalText([]byte{})
		e.MarshalText()
		e.String()
	})

	t.Run("first entry", func(t *testing.T) {
		enc, err := LIMITS_INIT.MarshalText()
		require.NoError(t, err)

		var dec LIMITS_STATE
		err = dec.UnmarshalText(enc)
		require.NoError(t, err)

		require.Equal(t, LIMITS_INIT, dec)
	})
}
