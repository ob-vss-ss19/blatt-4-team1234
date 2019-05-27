FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o useradministration/main useradministration/main.go

FROM iron/go
COPY --from=builder /app/useradministration/main /app/useradministration
EXPOSE 8091
ENTRYPOINT [ "/app/useradministration" ]
