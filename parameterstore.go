package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/bitmark-inc/config-loader/external/aws/ssm"
)

type ParameterStore struct {
	ParameterStore *ssm.ParameterStore
}

func NewParameterStore(ctx context.Context) (*ParameterStore, error) {
	awsParameterStore, err := ssm.New(ctx, config.WithDefaultRegion("ap-northeast-1"))
	if err != nil {
		return nil, err
	}

	return &ParameterStore{
		ParameterStore: awsParameterStore,
	}, nil
}

func ensureAWSParameterStore(ctx context.Context, ptrP **ParameterStore) error {
	if *ptrP == nil {
		p, err := NewParameterStore(ctx)
		if err != nil {
			return fmt.Errorf("fail to initialize aws parametere store. error: %s", err.Error())
		}

		*ptrP = p
	}

	return nil
}

func (p *ParameterStore) GetString(ctx context.Context, key string) (string, error) {
	if err := ensureAWSParameterStore(ctx, &p); err != nil {
		return "", err
	}

	return p.ParameterStore.GetString(ctx, key)
}

func (p *ParameterStore) PutString(ctx context.Context, key, value string) error {
	if err := ensureAWSParameterStore(ctx, &p); err != nil {
		return err
	}

	return p.ParameterStore.PutString(ctx, key, value)
}
