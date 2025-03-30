FROM golang:1.24

RUN apt-get update && apt-get install -y tzdata

WORKDIR /go

# Install golangci-lint & test library
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.58.1

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./be/ /go/src/legend_score
WORKDIR /go/src/legend_score

RUN go mod tidy

CMD ["air"]