package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func awsSession() (*session.Session, error) {
	return session.NewSession()
}

func awsSecretsEnv(s *secretsmanager.SecretsManager) ([]string, error) {
	var errors []error
	var env []string

	for envVarName, secretID := range config.SecretStringsAssignments.Values {
		key, value := envVarName, ""
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if result.SecretString != nil {
			value = *result.SecretString
		}
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	for envVarName, secretID := range config.SecretBinariesAssignments.Values {
		key, value := envVarName, ""
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		value = base64.StdEncoding.EncodeToString(result.SecretBinary)
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	for envVarName, secretID := range config.SecretBinaryStringsAssignments.Values {
		key, value := envVarName, ""
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		value = string(result.SecretBinary)
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	for envVarName, secret := range config.SecretJSONKeyStrings {
		key, value := envVarName, ""
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secret.SecretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		var jsonObject map[string]interface{}
		switch {
		case result.SecretString != nil:
			if err := json.Unmarshal([]byte(*result.SecretString), &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value = fmt.Sprint(jsonObject[secret.JSONKey])
		case result.SecretString != nil:
			if err := json.Unmarshal(result.SecretBinary, &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value = fmt.Sprint(jsonObject[secret.JSONKey])
		}
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}

	for envVarName, secret := range config.SecretJSONKeys {
		key, value := envVarName, ""
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secret.SecretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		var jsonObject map[string]interface{}
		switch {
		case result.SecretString != nil:
			if err := json.Unmarshal([]byte(*result.SecretString), &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value = fmt.Sprint(jsonObject[secret.JSONKey])
		case result.SecretString != nil:
			if err := json.Unmarshal(result.SecretBinary, &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value = fmt.Sprint(jsonObject[secret.JSONKey])
		}
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	if len(errors) == 1 {
		return env, errors[0]
	}
	if len(errors) > 0 {
		return env, fmt.Errorf("%d error(s): [%q, ...]", len(errors), errors[0])
	}
	return env, nil
}
