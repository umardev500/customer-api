FROM golang as dev

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 6012

CMD air

# Production for

FROM golang as prod

WORKDIR /app

COPY . .

RUN go build -o main
RUN cp main /bin/main

EXPOSE 6012

CMD /bin/main
