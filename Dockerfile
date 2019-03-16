FROM golang:latest

# Install Go
RUN \
    apt-get update && \
    apt-get install -y postgresql

# Set environment variables.
ENV DB_HOST localhost
ENV DB_USER postgres
ENV DB_PASSWORD postgres
ENV DB_NAME postgres
ENV DB_PORT 5432
ENV SVC_PORT :50051
ENV SRC_DIR $GOPATH/src/port-domain-svc

# Define working directory.
WORKDIR $GOPATH

COPY . $SRC_DIR
RUN cd $SRC_DIR; make build
