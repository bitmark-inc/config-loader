# Bitmark Config Loader

The config-loader reads configurations from both the local configuration file and environment variables and sets them using the `viper` package. A user can read them later using `viper`. For example:

```go
viper.GetString("my_key")
```

## Usage

After import the package, you simply call:

```go
config.LoadConfig("MY_CONF")
```

For more detail, please look into `example` folder.

# AWS Parameter Store

AWS Parameter Store store parameter and provide to client via SSM service

## Usage
### 1. Set up AWS Policy for IAM 
- Add this policy to your IAM, for example:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ssm:GetParameters",
                "ssm:PutParameters"
            ],
            "Resource": "*"
        }
    ]
}
```

- configure AWS credentials before use aws/ssm package