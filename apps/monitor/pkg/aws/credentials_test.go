package aws

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestGetCredentials(t *testing.T) {
	basePath := "../../hack/testdata/aws"

	t.Run("ValidCredentialsFile", func(t *testing.T) {
		credentials := aws.Credentials{
			AccessKeyID:     "accessKey",
			SecretAccessKey: "secretKey",
		}

		provider := GetCredentials(filepath.Join(basePath, "credentials"))

		ctx := context.TODO()
		creds, err := provider.Retrieve(ctx)
		assert.NoError(t, err)
		assert.Equal(t, credentials, creds)
	})
}

func Test_getCredentialsFunc(t *testing.T) {
	basePath := "../../hack/testdata/aws"

	t.Run("ValidCredentialsFile", func(t *testing.T) {
		credentials := aws.Credentials{
			AccessKeyID:     "accessKey",
			SecretAccessKey: "secretKey",
		}

		getCreds := getCredentials(filepath.Join(basePath, "credentials"))

		ctx := context.TODO()
		creds, err := getCreds(ctx)
		assert.NoError(t, err)
		assert.Equal(t, credentials, creds)
	})

	t.Run("InvalidCredentialsFile", func(t *testing.T) {
		expectedErr := errInvalidCredentialsFile

		getCreds := getCredentials(filepath.Join(basePath, "invalid-credentials"))

		ctx := context.TODO()
		_, err := getCreds(ctx)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("FileOpenError", func(t *testing.T) {
		expectedErr := os.ErrNotExist

		getCreds := getCredentials("/invalid/path")

		ctx := context.TODO()
		_, err := getCreds(ctx)
		assert.ErrorIs(t, err, expectedErr)
	})
}
