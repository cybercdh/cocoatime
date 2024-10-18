# cocoatime

`cocoatime` is a simple Go utility that converts Apple Cocoa Core Data timestamps into human-readable dates.

Core Data timestamps represent the number of seconds since January 1, 2001 (UTC). This utility helps convert those timestamps to standard human-readable date formats.

## Features

- Converts Cocoa Core Data timestamps to human-readable dates.
- Displays both the input timestamp and the converted date.

## Install

Assuming Go is installed:

```bash
go install github.com/cybercdh/cocoatime@latest`
```

## Usage

```bash
cocoatime 684223200

684223200,2022-09-07T06:00:00Z
```

