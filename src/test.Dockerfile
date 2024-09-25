FROM golang:1.23.1-alpine

WORKDIR /iad-build

# Copy in server source
COPY src /iad-build

# Build binary
RUN env GOOS=linux CGO_ENABLED=0 go test -c -cover -covermode=count -coverpkg=./... -o iad

# **************************************************************** #

FROM ubuntu:24.10

# Copy built binary
COPY --from=0 /iad-build/iad ./

user root

# Add start script and make it executable
ADD src/start.sh /home/iad/start.sh
RUN chmod +x /home/iad/start.sh

# Copy over built UI files
COPY src/page/static /home/iad/static

# # Own the entire iad home user
RUN chown -R iad:iad /home/iad

USER iad
