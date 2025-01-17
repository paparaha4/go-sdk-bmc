# Tests

## Setup

To run tests you'll need to run an instance of `mockserver` running at port 1080. This can be done using Docker:
- Install docker
- Prepare and start `mockserver` by running the following: `docker run --name mockserver -d --rm -p 1080:1080 mockserver/mockserver:5.13.0`

You can then get logs from the mock server by running `docker logs -f mockserver`. The container id can be found using `docker ps`.

The logs can also be viewed through the MockServer UI portal, by going to http://localhost:1080/mockserver/dashboard while the MockServer container is running.

Once you're done, you can stop the mock server by running `docker container stop -f mockserver`.

Alternatively, you can also run an instance of `portainer` by running the following command: `docker run -d -p 9000:9000 --restart always -v /var/run/docker.sock:/var/run/docker.sock -v /opt/portainer:/data portainer/portainer`. Going to `localhost:9000` will then give you a GUI where you can stop, start, and look at the logs for `mockserver`.

## Running

Tests can be run per-module by running its respective go script:

```sh
go test -v -run <TestSuiteName>
```

To run all tests for all modules at the same time, use:

```sh
go test -v -run ./
```

## Payloads

Payloads are created according to the mockserver documentation. In addition, they are formatted using `jq`. 

To format any additional payloads, install `jq` by doing `sudo apt-get install jq`, then run `format.sh` while inside the `payloads` folder.