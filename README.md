# Timbrelite Server

[![CircleCI](https://circleci.com/gh/adriangonz/timbrelite-server.svg?style=shield&circle-token=0553dee2b0d94b2637fd99cb98b9450128cb6c33)](https://circleci.com/gh/adriangonz/timbrelite-server)

API built in GO to interact with `timbrelite-client`.

It's responsible for receiving requests coming from IoT clients and pushing the events and notifications accordingly.

## Installing deps
We'll try out `dep` for this project. So, the first time, in order to fetch all deps we'll run:

```console
$ dep ensure
```

## Adding new deps

To install new deps, we'll run:

```console
$ dep ensure github.com/pkg/foo@^1.0.1
```
