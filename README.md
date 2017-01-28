# timbrelite-server

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
