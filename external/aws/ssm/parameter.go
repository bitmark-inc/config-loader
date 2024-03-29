package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/aws-sdk-go/aws"
)

type ParameterStore struct {
	client *ssm.Client
}

// NewParameterStore create new ParameterStore
// config will load secret, region from aws configure
func NewParameterStore(ctx context.Context) (*ParameterStore, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &ParameterStore{
		client: ssm.NewFromConfig(cfg),
	}, nil
}

func (p *ParameterStore) Get(ctx context.Context, key string) (value any, err error) {
	input := &ssm.GetParameterInput{
		Name: &key,
	}

	parameter, err := p.client.GetParameter(ctx, input)
	if err != nil {
		return
	}

	value = *parameter.Parameter.Value

	return
}

func (p *ParameterStore) Put(ctx context.Context, key string, value string) (err error) {
	input := &ssm.PutParameterInput{
		Name:      &key,
		Value:     &value,
		Type:      types.ParameterTypeString,
		Overwrite: aws.Bool(true),
	}

	_, err = p.client.PutParameter(ctx, input)
	return
}

// GetString find parameter in AWS SSM parameter store
func (p *ParameterStore) GetString(ctx context.Context, key string) (s string, err error) {
	val, err := p.Get(ctx, key)

	if val != nil && err == nil {
		s, _ = val.(string)
	}

	return
}
