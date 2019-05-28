FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o hallservice/main hallservice/main.go

FROM iron/go
COPY --from=builder /app/hallservice/main /app/hallservice
EXPOSE 8091
ENTRYPOINT [ "/app/hallservice" ]
