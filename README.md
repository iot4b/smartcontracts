# smartcontracts

### Requirements
* To compile contracts `everdev` tool should be installed: https://github.com/tonlabs/evernode-se
* Everscale client need some extra steps for installation, see here: https://github.com/markgenuine/ever-client-go

### Config
A tool uses config file `./config.sample.yml` by default. \
Copy it to `./config.yml` with your own values if necessary.

### Compile smartcontracts tool
```
go build -o smartcontracts ./main.go
```

### Create contract
Each smartcontract `{name}` is stored in a directory `./contract/{name}` which contain files:
* `{name}.sol` - Source code of a contract. Used only to compile contract `abi` and `tvc` files
* `{name}.abi.json` - Compiled ABI spec file. Used for contract deployment and methods execution.
* `{name}.tvc` - Contract compiled binary. Used only for contract deployment.

### Compile contract
To compile `abi` and `tvc` files from a contract code `{name}.sol` in a directory `./contract/{name}` go to that directory
```
cd ./contract/{name}
```
and use a command 
```
everdev sol compile {name}.sol
```

### Deploy contract
To deploy a contract `{name}` with initial `{balance}` use a command
```
  911  go run . elector new '{"defaultNodes":[]}'
  
  829  go run . owner new < builds/network_1/initial/owner.initial.json > builds/network_1/owner_1.json
  830  go run . owner new < builds/network_1/initial/owner.initial.json > builds/network_1/owner_2.json
  
  951  go run . vendor new < builds/network_1/initial/vendor.initial.json > builds/network_1/vendor_1.json
  
  918  go run . device new < builds/network_1/initial/device.initial.json > builds/network_1/device_1.json
  867  go run . device new < builds/network_1/initial/device_1.initial.json > builds/network_1/device_2.json
  919  go run . device new < builds/network_1/initial/device_2.initial.json > builds/network_1/device_3.json
  
  913  go run . node new < builds/network_1/initial/node-1.initial.json > builds/network_1/node_1.json
  914  go run . node new < builds/network_1/initial/node_2.initial.json > builds/network_1/node_2.json
  916  go run . node new < builds/network_1/initial/node_3.initial.json > builds/network_1/node_3.json
  
  
```
Transfers `{balance}` nanotokens from `giver` account in config to new address. \
Returns a newly deployed contract address.

### Execute contract method
To execute a `{method}` of a contract `{name}` deployed to `{address}` use a command
```
smartcontracts execute {name} {address} {method}
```
