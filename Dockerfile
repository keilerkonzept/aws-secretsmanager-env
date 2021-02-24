FROM golang:1.13.11-alpine3.10 AS build
RUN apk add --no-cache make jq
WORKDIR /go/src/github.com/keilerkonzept/aws-secretsmanager-env/
COPY . .
ENV CGO_ENABLED=0
RUN make binaries/linux_x86_64/aws-secretsmanager-env && mv binaries/linux_x86_64/aws-secretsmanager-env /app

FROM alpine:3.14
RUN apk add --no-cache ca-certificates
COPY --from=build /app /bin/aws-secretsmanager-env
CMD [ "aws-secretsmanager-env" ]
