package methods

import "github.com/bmizerany/assert"
import "testing"

func TestString(t *testing.T) {
	assert.Equal(t, "OPTIONS", OPTIONS.String())
	assert.Equal(t, "PUT", PUT.String())
}
