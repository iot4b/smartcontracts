# smartcontracts

### Requirements
* To compile contracts `everdev` tool should be installed: https://github.com/tonlabs/evernode-se
* Everscale client need some extra steps for installation, see here: https://github.com/markgenuine/ever-client-go

### Config
A tool uses config file `./config.yml` if exists. \
Otherwise default `./config.sample.yml` is used.

### Create contract
Each smartcontract of specific `{type}` is stored in a directory `./_{type}` which initially should contain a file
* `{type}.sol` - Source code of a contract, used to compile contract `abi` and `tvc` files.

### Compile contract
To compile `abi` and `tvc` files from a contract code `{type}.sol` in a directory `./_{type}` go to that directory
```
cd ./_{type}
```
and use a command 
```
everdev sol compile {type}.sol
```
that will create compiled contract files in its directory:
* `{type}.abi.json` - Compiled ABI spec file. Used for contract deployment and methods execution.
* `{type}.tvc` - Contract compiled binary. Used only for contract deployment.

### Compile smartcontracts tool
```
cd ./tools
go build .
```
`smartcontracts` executable will be created in `tools` directory.

## Usage

### Deploy contract
To deploy a contract `{type}` use a command
```
smartcontracts {type} new '{"initial": "data"}'
```
or
```
smartcontracts {type} new < path/to/initial.data.json
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
To execute a `{method}` of a contract `{type}` deployed to `{address}` use a command
```
smartcontracts execute {type} {address} {method}
```
