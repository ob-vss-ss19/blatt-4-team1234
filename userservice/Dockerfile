FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o userservice/main userservice/main.go

FROM iron/go
COPY --from=builder /app/userservice/main /app/userservice
EXPOSE 52000-53000
ENTRYPOINT [ "/app/userservice" ]