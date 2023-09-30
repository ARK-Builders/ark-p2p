package ark_p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		title  string
		assert func(t *testing.T, res any)
	}{
		{
			title: "Check Method not empty",
			assert: func(t *testing.T, res any) {
				assert.Greater(t, len(res.(string)), 0)
			},
		},
		{
			title: "Check result equal Hello world",
			assert: func(t *testing.T, res any) {
				assert.Equal(t, "Hello World", res)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.assert(t, HelloWorld())
		})
	}
}
