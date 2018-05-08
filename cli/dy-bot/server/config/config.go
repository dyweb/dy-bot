// Copyright 2018 The Dongyue members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

// Config is the type for server config.
type Config struct {
	Owner       string `yaml:"owner"`
	Repo        string `yaml:"repo"`
	HTTPListen  string `yaml:"port"`
	AccessToken string `yaml:"token"`
	WeeklyDir   string `yaml:"weekly_dir"`
}
