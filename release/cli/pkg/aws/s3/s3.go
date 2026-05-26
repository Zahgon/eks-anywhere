// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
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

package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Read reads the content of an object from an S3 bucket.
// It returns an io.ReadCloser that should be closed after use.
func Read(bucket, key string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func DownloadFile(filePath, bucket, key string, s3Downloader *s3manager.Downloader, private bool) error {
	_ = "STUB: not implemented"
	return nil
}

func UploadFile(filePath string, bucket, key *string, s3Uploader *s3manager.Uploader, private bool) error {
	_ = "STUB: not implemented"
	return nil
}

func KeyExists(s3Client *s3.S3, bucket, key string, private bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
