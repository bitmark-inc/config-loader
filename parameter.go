package config

import (
	"context"
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"

	"github.com/bitmark-inc/config-loader/external/aws/ssm"
)

// GetRSAPublishKeyFromParameterStore get RSA Publish Key from Parameter Store
func GetRSAPublishKeyFromParameterStore(ctx context.Context, parameterName string, parameterStore *ssm.ParameterStore) (*rsa.PublicKey, error) {
	parameter, err := parameterStore.FindParameterByName(ctx, parameterName)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(*parameter.Parameter.Value))
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
