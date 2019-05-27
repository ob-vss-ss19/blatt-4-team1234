FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o movieadministration/main movieadministration/main.go

FROM iron/go
COPY --from=builder /app/movieadministration/main /app/movieadministration
EXPOSE 8091
ENTRYPOINT [ "/app/movieadministration" ]
