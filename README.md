# goacme

An ACME Client for Let's Encrypt, written in Go.  
Made for Linux

## Installation

You can install goacme in three different ways.

## Using the goinit file

Get the file by running

```
wget https://raw.githubusercontent.com/lerichardson/goacme/main/goinit.go
```

Then use `go run goinit.go`. This is the best option if you have go installed.  
This will get the rest of the go files from GitHub, then build it and add it to your `PATH`.

## Get the pre-built file

Use `wget` or `curl` to get the built file.  
Make sure to switch your directory to something like `/opt`. Do not install in the root directory, or in the `/var` directory.  
`cURL`:

```
curl https://github.com/lerichardson/goacme/releases/download/v0.1.0/goacme
```

`wget`:

```
wget -O goacme https://github.com/lerichardson/goacme/releases/download/v0.1.0/goacme
```
