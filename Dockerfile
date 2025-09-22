# latest version compatible on windows nonserver and linux
ARG BUILDIMGTAG=1.24.2
ARG FINALIMAGE=gcr.io/distroless/base-debian11
# ARG BINARIESDIR=/app/bin #issue to recover the arg value in run command on windows
# ARG APPNAME=myapp #issue to recover the arg value in run command on windows
# ARG TESTDIR=/ashpp/test #issue to recover the arg value in run command on windows

FROM golang:$BUILDIMGTAG AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN mkdir -p "c:\app\bin"
RUN go build -o /app/bin/myapp

FROM build AS test

RUN go install github.com/jstemmer/go-junit-report@latest
RUN go install github.com/axw/gocov/gocov@latest
RUN go install github.com/AlekSi/gocov-xml@latest

RUN mkdir -p "/app/test"

RUN go test -v -coverprofile coverage.txt -covermode count ./... 2>&1 | go-junit-report > /app/test/junit.xml
RUN gocov convert coverage.txt | gocov-xml > /app/test/coverage.xml

FROM $FINALIMAGE AS final
ARG APPNAME
ARG BUILDNUMER=0.0.1-rc1

WORKDIR /

COPY --from=build /app/bin /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/app/myapp" ]

LABEL org..mycompany.schema-version=1.0 \
      org..mycompany.name=$APPNAME \
      org..mycompany.version=$BUILDNUMER