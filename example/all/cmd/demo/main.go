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

package main

import (
	"fmt"
	"log"

	"github.com/hahawangxv/kevin/example/all/internal/app/demo"
)

func main() {
	eng := demo.NewEngine()

	eng.Defer(func() error {
		fmt.Println("exit...")
		return nil
	})

	// eng.SetRegistry( // 多注册中心
	// 	compound.New(
	// 		etcdv3.StdConfig("wh01").BuildRegistry(),
	// 	),
	// )

	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}
}
