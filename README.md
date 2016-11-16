Fusis Balancer  [![Build Status](https://travis-ci.org/luizbafilho/fusis.svg?branch=master)](https://travis-ci.org/luizbafilho/fusis)
======

Fusis Balancer is a software [Layer 4](https://en.wikipedia.org/wiki/Transport_layer) Load Balancer powered by Linux's [IPVS](http://www.linuxvirtualserver.org/). It exposes a HTTP API to manage your services dynamically.

A layer 3 balancer take decisions based only on IP address (source and destination), a layer 4 balancer can also see transport information like TCP and UDP ports. Being a software balancer, it's tailored to be easy to deploy and scale.

## Why?
The goal of this project is to provide a friendly way to use IPVS.

It will be responsible for detecting new/failed nodes and add/remove routes to them automatically configuring the network to do so.

## State
This project it's under heavy development, it's not usable yet, but you can **Star** :star: the project and follow the updates.

# Installation

There is compilation and runtime dependency on [libnl](https://www.infradead.org/~tgr/libnl/).
On a Debian based system, you should be able to build it by running:

``` bash
sudo apt-get install libnl-3-dev libnl-genl-3-dev
```

Get this project into GOPATH:

``` bash
go get -v github.com/luizbafilho/fusis
```

And it's dependencies:

``` bash
make deps
```
You'll need **Go 1.5** or later;

## Installing IPVS

IPVS is a Kernel module. So, all you need to do is enable it via modprobe:
``` bash
sudo modprobe ip_vs
```

While you're at it, install the IPVS command line tool too:
``` bash
sudo apt-get install ipvsadm
```

And enable ipv4 forwarding with:
``` bash
sudo sysctl -w net.ipv4.ip_forward=1
```

## Running the project

Now that you have IPVS and fusis installed, run the project:

``` bash
# Remenber, fusis binary is at $GOPATH/bin/fusis. So, add it to your system PATH
sudo fusis balancer --bootstrap
```
You should see something like:
`[GIN-debug] Listening and serving HTTP on :8000`

From another host, send a HTTP request to the API querying for available services available:
``` bash
curl -i {IP OF FUSIS HOST}:8000/services
```
And you should get:
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 07 Apr 2016 21:23:18 GMT
Content-Length: 3

[]
```

Just for testing purposes, lets add a route to a fake IPv4 by runnging this on the fusis host:

``` bash
sudo ipvsadm -A -t 10.0.0.1:80 -s rr
```

Then, make another request:

``` bash
curl -i {FUSIS_HOST_IPV4}:8000/services
```

You will get that same route you just created as a response:
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 07 Apr 2016 22:08:42 GMT
Content-Length: 94

[{"Name":"","Host":"10.0.0.1","Port":80,"Protocol":"tcp","Scheduler":"rr","Destinations":[]}]
```

## Logging

Fusis uses [Logrus](https://github.com/Sirupsen/logrus) as its logging system.
By default, Fusis logs to stdout every minute.
You can change its log collection interval by passing the following command line argument:

```bash
# The argument --log-interval or -i. The value is in seconds
$> sudo fusis balancer --bootstrap --log-interval 10
 ```
