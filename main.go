package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var config struct {
	SecretStringsAssignments       AssignmentsMap
	SecretBinariesAssignments      AssignmentsMap
	SecretBinaryStringsAssignments AssignmentsMap
	SecretJSONKeyStringAssignments AssignmentsMap
	SecretJSONKeyAssignments       AssignmentsMap

	SecretJSONKeyStrings map[string]secretJSONKey
	SecretJSONKeys       map[string]secretJSONKey
	PrintEnvAndExit      bool
	Profile              string
}

type secretJSONKey struct {
	SecretID string
	JSONKey  string
}

func init() {
	config.SecretJSONKeyStrings = make(map[string]secretJSONKey)
	config.SecretJSONKeys = make(map[string]secretJSONKey)
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Ldate)
	log.SetPrefix("[aws-secretsmanager-env] ")
	flag.Var(&config.SecretStringsAssignments, "secret-string", "a key/value pair `ENV_VAR=SECRET_ARN` (may be specified repeatedly)")
	flag.Var(&config.SecretBinariesAssignments, "secret-binary-base64", "a key/value pair `ENV_VAR=SECRET_ARN` (may be specified repeatedly)")
	flag.Var(&config.SecretBinaryStringsAssignments, "secret-binary-string", "a key/value pair `ENV_VAR=SECRET_ARN` (may be specified repeatedly)")
	flag.Var(&config.SecretJSONKeyStringAssignments, "secret-json-key-string", "a key/value pair `ENV_VAR=SECRET_ARN#JSON_KEY` (may be specified repeatedly)")
	flag.Var(&config.SecretJSONKeyAssignments, "secret-json-key", "a key/value pair `ENV_VAR=SECRET_ARN#JSON_KEY` (may be specified repeatedly)")
	flag.StringVar(&config.Profile, "profile", "", "override the current AWS_PROFILE setting")
	flag.Parse()

	config.PrintEnvAndExit = flag.NArg() > 0

	for key, value := range config.SecretJSONKeyStringAssignments.Values {
		i := strings.IndexRune(value, '#')
		if i < 0 {
			log.Fatalf(`"%s" must have the form SECRET_ID#JSON_KEY`, value)
		}
		secretID, jsonKey := value[:i], value[i+1:]
		config.SecretJSONKeyStrings[key] = secretJSONKey{
			SecretID: secretID,
			JSONKey:  jsonKey,
		}
	}

	for key, value := range config.SecretJSONKeyAssignments.Values {
		i := strings.IndexRune(value, '#')
		if i < 0 {
			log.Fatalf(`"%s" must have the form SECRET_ID#JSON_KEY`, value)
		}
		secretID, jsonKey := value[:i], value[i+1:]
		config.SecretJSONKeys[key] = secretJSONKey{
			SecretID: secretID,
			JSONKey:  jsonKey,
		}
	}
}

func main() {
	awsSession, err := awsSession()
	if err != nil {
		log.Fatalf("aws: %v", err)
	}
	secretsEnv, err := awsSecretsEnv(secretsmanager.New(awsSession))
	if err != nil {
		log.Fatalf("error(s) while reading secrets: %v", err)
	}
	if !config.PrintEnvAndExit {
		for _, assignment := range secretsEnv {
			fmt.Println(assignment)
		}
		return
	}

	var osEnvAndSecretsEnv = append(os.Environ(), secretsEnv...)
	nameAndArgs := flag.Args()
	var args []string
	name := nameAndArgs[0]
	if len(nameAndArgs) > 1 {
		args = nameAndArgs[1:]
	}
	cmd := exec.Command(name, args...)
	cmd.Env = osEnvAndSecretsEnv
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("run %v: %v", cmd.Args, err)
	}
}
