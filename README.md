# Statist

The simple way to share MTProto proxy stats from the original container.

### Problem

You can not expose the port with stats outside from official docker-container.

> The MTProto proxy server exports internal statistics as tab-separated values over the http://localhost:2398/stats endpoint.
> Please note that this endpoint is available only from localhost: depending on your configuration,
> you may need to collect the statistics with docker exec MTProto proxy curl http://localhost:2398/stats.

This project solves this problem.

### Solution

Run separate container that can run `docker exec ...` command on MTProto proxy container.
It's not the best solution, but it's the easiest way, if You do not want to rebuild original MTProto proxy container.

### Usage

The simple docker-compose file must look like this:

```yaml
version: '3'

services:
  proxy:
    image: telegrammessenger/proxy:latest
    container_name: "${CONTAINER_NAME}"
    ports:
      - "443:443"
    volumes:
      - "proxy-config:/data"
    restart: unless-stopped
    environment:
    - TAG="${TAG}"
    - SECRET="${SECRET}"
  statist:
    image: kanaglic/statist:latest
    container_name: statist
    restart: unless-stopped
    ports:
      - "8080:8080" #if You want to share stats on this port
    command: "${CONTAINER_NAME}" #important to set correct container name of  MTProto proxy container.
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock #important to share origin docker socks
volumes:
  proxy-config:
```

Now you can see some stats on some port like this: `curl localhost:8080`.
Just for now output format is influxDB line format, but it's easy to make another formats - feel free to add issue or PR.

### Why influxDB

I use this output with [telegraf](https://docs.influxdata.com/telegraf/v1.7/) and
[influxDB](https://docs.influxdata.com/influxdb/v1.6/), so, just because.

### Develop

1.  `git clone`
2.  `dep ensure`
3.  `env GOOS=linux GOARCH=amd64 go build .`
4.  `docker build -t YOUR_REGISTRY/statist .`
5.  `docker push YOUR_REGISTRY/statist`

or use Makefile `make REGISTRY=YOUR_REGISTRY`

### Docker image

Yes, of course, You can use [the image from docker hub](https://hub.docker.com/r/kanaglic/statist/)
`docker pull kanaglic/statist`
