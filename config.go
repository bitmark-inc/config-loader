// SPDX-License-Identifier: ISC
// Copyright (c) 2019-2021 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig first reads `config.yaml` from a list of configuration path and merges
// the configurations with environment variables if there is any.
func LoadConfig(envPrefix string, extraConfigPaths ...string) {
	// Config from file
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/.config/")

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Warn("No config file. Read config from env.")
	}

	// merge config with extra configurations
	for _, p := range extraConfigPaths {
		viper.SetConfigFile(p)
		if err := viper.MergeInConfig(); err != nil {
			log.WithError(err).WithField("path", p).Warn("Fail to read config from path.")
		}
	}

	// Config from env if possible
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// allow log.level to be adjusted
	switch strings.ToUpper(viper.GetString("log.level")) {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}
}
