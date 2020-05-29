############################
# STEP 1 build executable binary
############################
# golang debian buster 1.14.1 linux/amd64
# https://github.com/docker-library/golang/blob/master/1.14/buster/Dockerfile
FROM golang@sha256:eee8c0a92bc950ecb20d2dffe46546da12147e3146f1b4ed55072c10cacf4f4c as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR $GOPATH/src/betpawa/test/

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod tidy
RUN go mod download
RUN go mod verify

COPY . .

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/betpawa-test .


############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/betpawa-test /go/bin/betpawa-test

# Use an unprivileged user.
USER appuser:appuser

# Run the hello binary.
ENTRYPOINT ["/go/bin/betpawa-test"]