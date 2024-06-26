// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package fingerprint

import "encoding/json"

// Config for fingerprint processor.
type Config struct {
	Method        namedHashMethod     `config:"method"`                     // Hash function to use for fingerprinting
	Fields        []string            `config:"fields" validate:"required"` // Source fields to compute fingerprint from
	TargetField   string              `config:"target_field"`               // Target field for the fingerprint
	Encoding      namedEncodingMethod `config:"encoding"`                   // Encoding to use for target field value
	IgnoreMissing bool                `config:"ignore_missing"`             // Ignore missing fields?
}

func defaultConfig() Config {
	return Config{
		Method:        hashes["sha256"],
		TargetField:   "fingerprint",
		Encoding:      encodings["hex"],
		IgnoreMissing: false,
	}
}

func (c *Config) MarshalJSON() ([]byte, error) {
	type Alias Config
	return json.Marshal(&struct {
		Method   string
		Encoding string
		*Alias
	}{
		Method:   c.Method.Name,
		Encoding: c.Encoding.Name,
		Alias:    (*Alias)(c),
	})
}
