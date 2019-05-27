FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o showadministration/main showadministration/main.go

FROM iron/go
COPY --from=builder /app/showadministration/main /app/showadministration
EXPOSE 8091
ENTRYPOINT [ "/app/showadministration" ]
