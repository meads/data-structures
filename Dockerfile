############################
# STEP 1 build executable binary
############################
# golang alpine 1.12
FROM golang@sha256:8cc1c0f534c0fef088f8fe09edc404f6ff4f729745b85deae5510bfd4c157fb2 as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/github.com/meads/datastructures
COPY . .

# Fetch dependencies.
RUN go get -d -v

# create a test binary
RUN go test -c $GOPATH/src/github.com/meads/datastructures/pkg/trie/ -o /go/bin/trie.test

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/datastructures .

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /go/bin/datastructures /go/bin/datastructures
COPY --from=builder /go/bin/trie.test /go/bin/trie.test

# Use an unprivileged user.
USER appuser

# Run the datastructures binary.
ENTRYPOINT ["/go/bin/trie.test"]