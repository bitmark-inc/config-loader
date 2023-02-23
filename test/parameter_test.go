package test

import (
	"context"
	"github.com/bitmark-inc/config-loader/external/aws/ssm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetString(t *testing.T) {
	ctx := context.Background()
	key := "/jwtPubKey"

	parameterStore, _ := ssm.NewParameterStore(ctx)
	s, err := parameterStore.GetString(ctx, key)
	assert.NoError(t, err)
	assert.NotNil(t, s)
}
