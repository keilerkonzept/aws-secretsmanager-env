# ${APP}

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sgreben/aws-secretsmanager-env.svg)](https://hub.docker.com/r/sgreben/aws-secretsmanager-env/tags)

Injects AWS Secrets Manager secrets as environment variables - or just prints them, if no command is given.

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
  - [Examples](#examples)
- [Comments](#comments)

<!-- /TOC -->

## Get it

Using go get:

```bash
go get -u github.com/sgreben/${APP}
```

Or [download the binary](https://github.com/sgreben/${APP}/releases/latest) from the releases page.

```bash
# Linux
curl -L https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_windows_x86_64.zip
unzip ${APP}_${VERSION}_windows_x86_64.zip
```

## Use it

```text

${APP} [OPTIONS] [COMMAND [ARGS...]]

${USAGE}
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
