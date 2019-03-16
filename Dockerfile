FROM postgres

# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
		wget \
	&& rm -rf /var/lib/apt/lists/*

RUN wget --no-check-certificate -c https://golang.org/dl/go1.12.1.linux-amd64.tar.gz -O - | tar -xz -C /usr/local/

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go version

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Set environment variables.
ENV DB_HOST localhost
ENV DB_USER postgres
ENV DB_PASSWORD postgres
ENV DB_NAME postgres
ENV DB_PORT 5432
ENV SVC_PORT :50051
ENV SRC_DIR $GOPATH/src/port-domain-svc

WORKDIR $SRC_DIR

COPY . $SRC_DIR
RUN make build
