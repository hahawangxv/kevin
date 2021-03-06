// Copyright 2020 Douyu
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcdv3

import (
	"time"

	"github.com/hahawangxv/kevin/pkg/ecode"
	"github.com/hahawangxv/kevin/pkg/registry"

	"github.com/hahawangxv/kevin/pkg/client/etcdv3"
	"github.com/hahawangxv/kevin/pkg/conf"
	"github.com/hahawangxv/kevin/pkg/xlog"
)

// StdConfig ...
func StdConfig(name string) *Config {
	return RawConfig("jupiter.registry." + name)
}

// RawConfig ...
func RawConfig(key string) *Config {
	var config = DefaultConfig()
	// 解析最外层配置
	if err := conf.UnmarshalKey(key, &config); err != nil {
		xlog.Panic("unmarshal key", xlog.FieldMod("registry.etcd"), xlog.FieldErrKind(ecode.ErrKindUnmarshalConfigErr), xlog.FieldErr(err), xlog.String("key", key), xlog.Any("config", config))
	}
	// 解析嵌套配置
	if err := conf.UnmarshalKey(key, &config.Config); err != nil {
		xlog.Panic("unmarshal key", xlog.FieldMod("registry.etcd"), xlog.FieldErrKind(ecode.ErrKindUnmarshalConfigErr), xlog.FieldErr(err), xlog.String("key", key), xlog.Any("config", config))
	}
	return config
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		Config:      etcdv3.DefaultConfig(),
		ReadTimeout: time.Second * 3,
		Prefix:      "jupiter",
		logger:      xlog.DefaultLogger,
	}
}

// Config ...
type Config struct {
	*etcdv3.Config
	ReadTimeout time.Duration
	ConfigKey   string
	Prefix      string
	logger      *xlog.Logger
}

// BuildRegistry ...
func (config Config) BuildRegistry() registry.Registry {
	if config.ConfigKey != "" {
		config.Config = etcdv3.RawConfig(config.ConfigKey)
	}
	return newETCDRegistry(&config)
}

// BuildResolver ...
func (config Config) BuildResolver() *etcdResolver {
	if config.ConfigKey != "" {
		config.Config = etcdv3.RawConfig(config.ConfigKey)
	}
	return newETCDResolver(&config)
}
