package iaws

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestGetCredentials(t *testing.T) {
	basePath := "../hack/testdata/aws"

	t.Run("ValidCredentialsFile", func(t *testing.T) {
		credentials := aws.Credentials{
			AccessKeyID:     "accessKey",
			SecretAccessKey: "secretKey",
		}

		provider := GetCredentials(filepath.Join(basePath, "test-credentials.json"))

		ctx := context.TODO()
		creds, err := provider.Retrieve(ctx)
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

func Test_readDefaultCredentials(t *testing.T) {
	basePath := "../hack/testdata/aws"

	t.Run("ValidCredentialsFile", func(t *testing.T) {
		credentials := aws.Credentials{
			AccessKeyID:     "accessKey",
			SecretAccessKey: "secretKey",
		}

		ctx := context.TODO()
		creds, err := readDefaultCredentials(ctx, filepath.Join(basePath, "cli-credentials"))
		assert.NoError(t, err)
		assert.Equal(t, credentials, creds)
	})

	t.Run("InvalidCredentialsFile", func(t *testing.T) {
		expectedErr := errInvalidCredentialsFile

		ctx := context.TODO()
		_, err := readDefaultCredentials(ctx, filepath.Join(basePath, "invalid-credentials"))
		assert.ErrorIs(t, err, expectedErr)
	})
}
