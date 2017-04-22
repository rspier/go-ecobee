# go-ecobee

A go library and simple tools for accessing the
[Ecobee Thermostat API](https://www.ecobee.com/home/developer/api/documentation/v1).

## Setup

### Build

go install github.com/rspier/go-ecobee/cli

### Application ID

Obtain an Ecobee API Application ID at
https://www.ecobee.com/consumerportal/index.html#/dev

We'll store it in the APPID environment variable for the examples
below.

## Thermostat ID

You'll need to know the serial number (Thermostat ID) of your
thermostat.  You can find it on the About page of the
[Ecobee portal](https://www.ecobee.com/consumerportal/index.html).

We'll store it in the THERMID environment variable for the examples
below.


## CLI Usage

### Status

```shell
$ ./cli --appid=${APPID} --thermostat=${THERMID} --command=status
Current Settings (Home): 68.0 - 75.0.  Fan: auto
Holding at 68.0 - 75.0 (Fan: auto) until 2017-04-22 00:00:00
Temperature: 75.0
  Bedroom\*: 74.7
  My ecobee3\*: 76.2
```

### Hold Temperature

```shell
$ ./cli --appid=${APPID} --thermostat=${THERMID} --command=hold --heat 70 --cool 76 --duration 10m
Successfully Held Temperature
```

### Message

```shell
$ ./cli --appid=${APPID} --thermostat=${THERMID} --command=message --message="Hi there!"
Successfully sent message: "Hi there!"
```

### Fan

```shell
$ ./cli --appid=${APPID} --thermostat=${THERMID} --command=fan --duration=5m
Running fan for 5 minutes"
```

### List

```shell
$ ./cli --command=list
${THERMID}: My ecobee3
```

## Ideas for future development

### Sensor Failure Test

Use includeSensors on the thermostat endpoint to analyze sensors where
the presence sensor may have failed -- i.e. present all the time.

## Disclaimer

This is not an official Google project.
