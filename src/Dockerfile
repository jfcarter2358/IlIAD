FROM golang:1.23.1-alpine

WORKDIR /${APP_NAME}-build

# Copy in server source
COPY src /${APP_NAME}-build

# Build binary
RUN env GOOS=linux CGO_ENABLED=0 go build -v -o ${APP_NAME}

# **************************************************************** #

FROM ubuntu:24.10

# Copy built binary
COPY --from=0 /${APP_NAME}-build/${APP_NAME} ./

user root

# Add start script and make it executable
ADD src/start.sh /home/${APP_NAME}/start.sh
RUN chmod +x /home/${APP_NAME}/start.sh

# Copy over built UI files
COPY src/page/static /home/${APP_NAME}/static

# # Own the entire ${APP_NAME} home user
RUN chown -R ${APP_NAME}:${APP_NAME} /home/${APP_NAME}

USER ${APP_NAME}
