package nilvalidator_test

import (
	"github.com/rentifly/nilvalidator"
	"github.com/stretchr/testify/require"
	"testing"
)

type testDeps struct {
	A *int        `nilvalidator:"notnil"`
	B interface{} `nilvalidator:"notnil"`
	C string
}

func TestValidateStructNotNil(t *testing.T) {
	val := 42
	valid := testDeps{A: &val, B: "test"}

	err := nilvalidator.ValidateStructNotNil(valid)
	require.NoError(t, err)
}

func TestValidateStructWithNil(t *testing.T) {
	invalid := testDeps{}
	err := nilvalidator.ValidateStructNotNil(invalid)
	require.Error(t, err)
	require.Contains(t, err.Error(), "field 'A' is nil")
}
