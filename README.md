# Errlogbeat

Errlogbeat collects entries from the error log of the AIX operating system and ships them to Elasticsearch or Logstash.

## Getting Started with Errlogbeat

### Requirements

#### Build Requirements

- An AIX operating system instance. I'm using version 7.1. Support for Go should be even better on version 7.2 but I haven't try it yet.
- A recent version of `gcc-go` and its dependencies. I'm using the packages kindly provided by [BullFreeware](http://www.bullfreeware.com/search.php?package=gcc-go).
- Various open source packages, at least GNU `make`. Some tasks also require `find-utils` and Python. I'm using the packages kindly provided by [Michael Perzl](http://www.perzl.org/aix/).

#### Runtime Requirements

- `libgo` should be the same as in the build environment.
- `libgcc`, not necessarily at the same level or from the same source like `libgo`.

### Build

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/WuerthIT/errlogbeat`

To build the binary for Errlogbeat run the command below. This will generate a binary
in the same directory with the name errlogbeat.

```
gmake
```

### Run

To run Errlogbeat with debugging output enabled, run:

```
./errlogbeat -c errlogbeat.yml -e -d "*"
```

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```

### Cleanup

To clean Errlogbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

## Vendoring

Errlogbeat currently includes version 6.2.4 of beats in the `vendor` subfolder with [some minor modifications](https://github.com/WuerthIT/beats/releases/tag/v6.2.4-support_aix) on libraries inside their `vendor` directory. Later versions make use of Go modules that are not available on the AIX operation system currently.

## Packaging

The original packaging process makes use of containers and will obviously not work here. So currently the binary file has to be distributed manually. Maybe an RPM spec file will be provided later.

## Disclaimer

AIX is a registered trademark of the International Business Machines Corporation.

Elasticsearch and Logstash are trademarks of Elasticsearch BV.

Errlogbeat is not endorsed by any of these companies.
