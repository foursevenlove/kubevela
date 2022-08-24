/*
Copyright 2022 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package compression

// Type the compression type
type Type string

const (
	// Uncompressed does not compress or encode data
	Uncompressed Type = ""
	// Gzip compresses data using gzip and encodes it using base64
	Gzip Type = "gzip"
	// Zstd compresses data using zstd and encodes it using base64
	Zstd Type = "zstd"
)
