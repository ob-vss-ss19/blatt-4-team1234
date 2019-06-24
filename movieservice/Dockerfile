FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o movieservice/main movieservice/main.go

FROM iron/go
COPY --from=builder /app/movieservice/main /app/movieservice
EXPOSE 52000-53000
ENTRYPOINT [ "/app/movieservice" ]