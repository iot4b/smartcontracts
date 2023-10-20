# smartcontracts

### Requirements
* To compile contracts `everdev` tool should be installed: https://github.com/tonlabs/evernode-se
* Everscale client need some extra steps for installation, see here: https://github.com/markgenuine/ever-client-go

### Config
A tool uses config file `./config.yml` if exists. \
Otherwise default `./config.sample.yml` is used.

### Create contract
Each smartcontract `{name}` is stored in a directory `./{name}` which initially should contain a file
* `{name}.sol` - Source code of a contract, used to compile contract `abi` and `tvc` files.

### Compile contract
To compile `abi` and `tvc` files from a contract code `{name}.sol` in a directory `./{name}` go to that directory
```
cd ./{name}
```
and use a command 
```
everdev sol compile {name}.sol
```
that will create compiled contract files in its directory:
* `{name}.abi.json` - Compiled ABI spec file. Used for contract deployment and methods execution.
* `{name}.tvc` - Contract compiled binary. Used only for contract deployment.

### Compile smartcontracts tool
```
cd ./tools
go build .
```
`smartcontracts` executable will be created in `tools` directory.

## Usage

### Deploy contract
To deploy a contract `{name}` use a command
```
smartcontracts {name} new '{"initial": "data"}'
```
or
```
smartcontracts {name} new < path/to/initial.data.json
```
#### Examples:
```
smartcontracts device new < builds/network_1/initial/device.initial.json > builds/network_1/device_1.json
smartcontracts device new < builds/network_1/initial/device_1.initial.json > builds/network_1/device_2.json
smartcontracts device new < builds/network_1/initial/device_2.initial.json > builds/network_1/device_3.json

smartcontracts node new < builds/network_1/initial/node_1.initial.json > builds/network_1/node_1.json
smartcontracts node new < builds/network_1/initial/node_2.initial.json > builds/network_1/node_2.json
smartcontracts node new < builds/network_1/initial/node_3.initial.json > builds/network_1/node_3.json

smartcontracts elector new '{"defaultNodes":[]}'

smartcontracts owner new < builds/network_1/initial/owner.initial.json > builds/network_1/owner_1.json
smartcontracts owner new < builds/network_1/initial/owner.initial.json > builds/network_1/owner_2.json

smartcontracts vendor new < builds/network_1/initial/vendor.initial.json > builds/network_1/vendor_1.json
```

### Execute contract method
To execute a `{method}` of a contract `{name}` deployed to `{address}` use a command
```
smartcontracts execute {name} {address} {method}
```
