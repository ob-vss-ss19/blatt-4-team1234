FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o halladministration/main halladministration/main.go

FROM iron/go
COPY --from=builder /app/halladministration/main /app/halladministration
EXPOSE 8091
ENTRYPOINT [ "/app/halladministration" ]
