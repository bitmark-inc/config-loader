package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type ParameterStore struct {
	region string
	client *ssm.Client
}

// New creates a new AWS ParameterStore instance
// It loads aws settings(secret, region) from the environment by default.
func New(ctx context.Context, optFns ...func(*config.LoadOptions) error) (*ParameterStore, error) {
	cfg, err := config.LoadDefaultConfig(ctx, optFns...)
	if err != nil {
		return nil, err
	}

	return &ParameterStore{
		region: cfg.Region,
		client: ssm.NewFromConfig(cfg),
	}, nil
}

func (p *ParameterStore) errorWithRegion(err error) error {
	if err != nil {
		return fmt.Errorf("%s, region: %s", err, p.region)
	}
	return nil
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

func (p *ParameterStore) GetString(ctx context.Context, key string) (string, error) {
	input := &ssm.GetParameterInput{
		Name: &key,
	}

	parameter, err := p.client.GetParameter(ctx, input)
	if err != nil {
		return "", p.errorWithRegion(err)
	}

	if parameter.Parameter.Type != types.ParameterTypeString {
		return "", fmt.Errorf("parameter type is not string")
	}

	return *parameter.Parameter.Value, nil
}

func (p *ParameterStore) GetSecretString(ctx context.Context, key string) (string, error) {
	input := &ssm.GetParameterInput{
		Name:           &key,
		WithDecryption: aws.Bool(true),
	}

	parameter, err := p.client.GetParameter(ctx, input)
	if err != nil {
		return "", p.errorWithRegion(err)
	}

	if parameter.Parameter.Type != types.ParameterTypeSecureString {
		return "", fmt.Errorf("parameter type is not string")
	}

	return *parameter.Parameter.Value, nil
}

func (p *ParameterStore) PutString(ctx context.Context, key string, value string) error {
	input := &ssm.PutParameterInput{
		Name:      &key,
		Value:     &value,
		Type:      types.ParameterTypeString,
		Overwrite: aws.Bool(true),
	}

	_, err := p.client.PutParameter(ctx, input)
	return p.errorWithRegion(err)
}
