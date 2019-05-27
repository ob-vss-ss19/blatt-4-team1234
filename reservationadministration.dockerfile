FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o reservationadministration/main reservationadministration/main.go

FROM iron/go
COPY --from=builder /app/reservationadministration/main /app/reservationadministration
EXPOSE 8091
ENTRYPOINT [ "/app/reservationadministration" ]
