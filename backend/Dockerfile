FROM golang:1.17

RUN apt-get update -y && apt install -y netcat iputils-ping traceroute
# Move to working directory (/build).

# install Air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin


RUN mkdir /build
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download


# Copy the code into the container.
COPY . /build

RUN chmod +x scripts/prestart.sh
RUN chmod +x scripts/dev_start.sh
RUN chmod +x scripts/prod_start.sh

# Set necessary environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64