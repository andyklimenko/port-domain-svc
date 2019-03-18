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

ENV POSTGRES_HOST localhost
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_NAME postgres
ENV POSTGRES_PORT 5432
ENV SVC_PORT :50051
ENV SRC_DIR $GOPATH/src/port-domain-svc

EXPOSE 5432 50051

WORKDIR $SRC_DIR

COPY . $SRC_DIR
RUN make build
