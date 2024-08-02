package aws

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	errInvalidCredentialsFile = errors.New("invalid credentials file")
)

// GetCredentials returns a new aws.CredentialsProvider
// that reads AWS credentials from a file.
func GetCredentials(filePath string) aws.CredentialsProvider {
	return aws.NewCredentialsCache(aws.CredentialsProviderFunc(getCredentials(filePath)))
}

func getCredentials(filePath string) func(context.Context) (aws.Credentials, error) {
	return func(ctx context.Context) (aws.Credentials, error) {
		if filePath == "" {
			homedir, err := os.UserHomeDir()
			if err != nil {
				return aws.Credentials{}, fmt.Errorf("invalid credentials file: %w", err)
			}
			filePath = filepath.Join(homedir, ".aws", "credentials")
		}
		file, err := os.Open(filepath.Clean(filePath))
		if err != nil {
			return aws.Credentials{}, err
		}
		defer file.Close()

		accessKey := ""
		secretKey := ""

		// Read credentials from file
		r := bufio.NewReader(file)
		for i := 0; i < 2; i++ {
			bytes, err := r.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					return aws.Credentials{}, err
				}
			}

			// Parse credentials
			s := strings.Split(string(bytes), "=")
			if len(s) != 2 {
				return aws.Credentials{}, errInvalidCredentialsFile
			}

			key := s[0]
			switch {
			case strings.Contains(key, "access"):
				accessKey = strings.TrimSpace(s[1])
			case strings.Contains(key, "secret"):
				secretKey = strings.TrimSpace(s[1])
			default:
				return aws.Credentials{}, errInvalidCredentialsFile
			}
		}

		return aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}, nil
	}
}
