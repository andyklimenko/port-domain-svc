FROM golang:latest

# Install Go
RUN \
    apt-get update && \
    apt-get install -y postgresql

# Set environment variables.
#ENV GOROOT /usr/local/go
#ENV GOPATH /home/gopath
#ENV PATH /usr/local/go/bin:$PATH
ENV DB_HOST localhost
ENV DB_USER postgres
ENV DB_PASSWORD postgres
ENV DB_NAME postgres
ENV SVC_PORT :50051

# Define working directory.
WORKDIR $GOPATH

COPY . $GOPATH/ports
RUN cd $GOPATH/ports
RUN ls -al
RUN make build
#RUN ./port-domain-svc migrate
