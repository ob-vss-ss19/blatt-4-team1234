# Reservationservice Service

This is the Reservationservice service

Generated with

```
micro new github.com/ob-vss-ss19/blatt-4-team1234/reservationservice --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.reservationservice
- Type: srv
- Alias: reservationservice

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./reservationservice-srv
```

Build a docker image
```
make docker
```