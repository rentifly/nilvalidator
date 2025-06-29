package nilvalidator_test

import (
	"github.com/rentifly/nilvalidator"
	"github.com/stretchr/testify/require"
	"testing"
)

type testDeps struct {
	A *int        `nilvalidator:"required"`
	B interface{} `nilvalidator:"required"`
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

func TestValidateStructPointer(t *testing.T) {
	val := 42
	valid := &testDeps{A: &val, B: "ok"}

	err := nilvalidator.ValidateStructNotNil(valid)
	require.NoError(t, err)
}

func TestValidateStructPointerReturnErr(t *testing.T) {
	var notValid *testDeps
	err := nilvalidator.ValidateStructNotNil(notValid)
	require.Error(t, err)
	require.Contains(t, err.Error(), "ValidateStructNotNil: nil pointer")
}
