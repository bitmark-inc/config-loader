package config

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetStringNotFound(t *testing.T) {
	ctx := context.Background()
	key := "/notfound"

	c := LoadConfig("CONFIG_TEST")

	value, err := c.ParameterStore.GetString(ctx, key)

	assert.NotNil(t, err)
	assert.Empty(t, value)
	assert.Contains(t, err.Error(), "region: ap-northeast-1")
}

func TestGetString(t *testing.T) {
	ctx := context.Background()
	key := "/test"

	c := LoadConfig("CONFIG_TEST")

	value, err := c.ParameterStore.GetString(ctx, key)

	assert.NoError(t, err)
	assert.Equal(t, "test", value)
}

func TestPutString(t *testing.T) {
	ctx := context.Background()
	key := "/testWrite"

	c := LoadConfig("CONFIG_TEST")

	now := time.Now()
	err := c.ParameterStore.PutString(ctx, key, now.String())
	assert.NoError(t, err)

	value, err := c.ParameterStore.GetString(ctx, key)

	assert.NoError(t, err)
	assert.Equal(t, now.String(), value)
}
