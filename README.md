# aws-secretsmanager-env

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sgreben/aws-secretsmanager-env.svg)](https://hub.docker.com/r/sgreben/aws-secretsmanager-env/tags)

Injects AWS Secrets Manager secrets as environment variables - or just prints them, if no command is given. (If you need secrets as **files** instead, you can use [aws-secretsmanager-files](https://github.com/sgreben/aws-secretsmanager-files))

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
  - [Examples](#examples)
- [Comments](#comments)

<!-- /TOC -->

## Get it

Using go get:

```bash
go get -u github.com/sgreben/aws-secretsmanager-env
```

Or [download the binary](https://github.com/sgreben/aws-secretsmanager-env/releases/latest) from the releases page.

```bash
# Linux
curl -L https://github.com/sgreben/aws-secretsmanager-env/releases/download/1.1.0/aws-secretsmanager-env_1.1.0_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/aws-secretsmanager-env/releases/download/1.1.0/aws-secretsmanager-env_1.1.0_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/aws-secretsmanager-env/releases/download/1.1.0/aws-secretsmanager-env_1.1.0_windows_x86_64.zip
unzip aws-secretsmanager-env_1.1.0_windows_x86_64.zip
```

## Use it

```text

aws-secretsmanager-env [OPTIONS] [COMMAND [ARGS...]]

Usage of aws-secretsmanager-env:
  -profile string
    	override the current AWS_PROFILE setting
  -secret-binary-base64 ENV_VAR=SECRET_ARN
    	a key/value pair ENV_VAR=SECRET_ARN (may be specified repeatedly)
  -secret-binary-string ENV_VAR=SECRET_ARN
    	a key/value pair ENV_VAR=SECRET_ARN (may be specified repeatedly)
  -secret-json-key ENV_VAR=SECRET_ARN#JSON_KEY
    	a key/value pair ENV_VAR=SECRET_ARN#JSON_KEY (may be specified repeatedly)
  -secret-json-key-string ENV_VAR=SECRET_ARN#JSON_KEY
    	a key/value pair ENV_VAR=SECRET_ARN#JSON_KEY (may be specified repeatedly)
  -secret-string ENV_VAR=SECRET_ARN
    	a key/value pair ENV_VAR=SECRET_ARN (may be specified repeatedly)
```

### Examples

```shell
$ aws-secretsmanager-env -secret-string MY_SECRET=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1
MY_SECRET={"hello":"world"}

$ aws-secretsmanager-env -secret-json-key MY_SECRET=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
MY_SECRET="world"

$ aws-secretsmanager-env -secret-json-key-string MY_SECRET=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
MY_SECRET=world

$ aws-secretsmanager-env -secret-json-key-string MY_SECRET=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello sh -c 'echo the secret is "$MY_SECRET"'
the secret is "world"
```
