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

package rabbitmq

import (
	"github.com/hahawangxv/kevin/pkg/conf"
	"github.com/hahawangxv/kevin/pkg/xlog"
	"github.com/streadway/amqp"
)

// ProducerConfig producer config
type ProducerConfig struct {
	Addr string `json:"addr" toml:"addr"`
}

// StdProducerConfig ...
func StdProducerConfig(name string) ProducerConfig {
	return RawProducerConfig("jupiter.rabbitmq." + name + ".stub")
}

// StdProducerConfig ...
func TemplateMqConfig(name string) ProducerConfig {
	return RawProducerConfig("template.rabbitmq." + name + ".stub")
}

// RawProducerConfig ...
func RawProducerConfig(key string) ProducerConfig {
	var config = DefaultProducerConfig()
	if err := conf.UnmarshalKey(key, &config); err != nil {
		xlog.Panic("unmarshal config", xlog.String("key", key))
	}
	return config
}

// DefaultProducerConfig ...
func DefaultProducerConfig() ProducerConfig {
	return ProducerConfig{}
}

// Build ...
func (config ProducerConfig) Build() *amqp.Connection {
	// 兼容配置
	client, err := amqp.Dial(config.Addr)
	if err != nil {
		xlog.Panic("amqp.Dial err", xlog.Any("err", err))
	}
	/*
		Channel, err := client.Channel()
		if err != nil {
			xlog.Panic("amqp.Dial err", xlog.Any("err", err))
		}*/
	//defer client.Close()
	return client
}
