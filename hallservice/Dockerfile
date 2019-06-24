FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go get -d -v ./...
RUN go build -o hallservice/main hallservice/main.go

FROM iron/go
COPY --from=builder /app/hallservice/main /app/hallservice
EXPOSE 52000-53000
ENTRYPOINT [ "/app/hallservice" ]
