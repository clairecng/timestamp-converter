# Timestamp-converter

This tool is a small script converting timestamps in a epoch seconds, milliseconds or in RFC3339 format to view it in those other formats. Made solely to make my life easier when testing with timestamps.

## BUILD

```
go get
go build
```

## USAGE

```
./timestamp-converter --type SECONDS --value 1626796500
```

Results
```
Seconds: 1626796500
Milliseconds: 1626796500000
RFC3339: 2021-07-20T17:55:00+02:00
```
