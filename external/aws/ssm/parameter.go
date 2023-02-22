package ssm

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
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

// GetKey find parameter in AWS SSM parameter store
func (s *ParameterStore) GetKey(ctx context.Context, parameterName string) (*ssm.GetParameterOutput, error) {
	input := &ssm.GetParameterInput{
		Name: &parameterName,
	}

	parameter, err := s.client.GetParameter(ctx, input)
	if err != nil {
		return nil, err
	}

	return parameter, nil
}
