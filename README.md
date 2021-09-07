# Bitmark Config Loader

The config-loader reads configurations from both the local configuration file and environment variables and sets them using the `viper` package. A user can read them later using `viper`. For example:

```go
viper.GetString("my_key")
```

## Usage

After import the package, you simply call:

```go
config.LoadConfig("my_config.yaml", "MY_CONF")
```

For more detail, please look into `example` folder.
