package aws

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	errInvalidCredentialsFile = errors.New("invalid credentials file")
)

func GetCredentials(filePath string) aws.CredentialsProvider {
	return aws.NewCredentialsCache(aws.CredentialsProviderFunc(getCredentials(filePath)))
}

func getCredentials(filePath string) func(context.Context) (aws.Credentials, error) {
	return func(ctx context.Context) (aws.Credentials, error) {
		file, err := os.Open(filePath)
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
				return aws.Credentials{}, err
			}

			// Parse credentials
			s := strings.Split(string(bytes), "=")
			if len(s) != 2 {
				return aws.Credentials{}, errInvalidCredentialsFile
			}

			if strings.Contains(s[0], "access") {
				accessKey = strings.TrimSpace(s[1])
			} else if strings.Contains(s[1], "secret") {
				secretKey = strings.TrimSpace(s[1])
			} else {
				return aws.Credentials{}, errInvalidCredentialsFile
			}
		}

		return aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}, nil
	}
}
