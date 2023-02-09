# Deploy Your Own MeZonHub Testnet

This document describes 3 ways to setup a network of `mezon` nodes, each serving a different usecase:

1. Single-node, local, manual testnet
2. Multi-node, local, automated testnet
3. Multi-node, remote, automated testnet

Supporting code can be found in the [networks directory](https://github.com/mezonhub/network) and additionally the `local` or `remote` sub-directories.

> NOTE: The `remote` network bootstrapping may be out of sync with the latest releases and is not to be relied upon.

## Available Docker images

In case you need to use or deploy mezonhub as a container you could skip the `build` steps and use the official images, \$TAG stands for the version you are interested in:

- `docker run -it -v ~/.mezonhub:/root/.mezonhub tendermint:$TAG mezon init`
- `docker run -it -p 26657:26657 -p 26656:26656 -v ~/.mezonhub:/root/.mezonhub tendermint:$TAG mezon start`
- ...
- `docker run -it -v ~/.mezonhub:/root/.mezonhub tendermint:$TAG mezon version`

The same images can be used to build your own docker-compose stack.

## Single-node, Local, Manual Testnet

This guide helps you create a single validator node that runs a network locally for testing and other development related uses.

### Requirements

- [Install mezonhub](./install.md)
- [Install `jq`](https://stedolan.github.io/jq/download/) (optional)

### Create Genesis File and Start the Network

```bash
# You can run all of these commands from your home directory
cd $HOME

# Initialize the genesis.json file that will help you to bootstrap the network
mezon init my-node --chain-id my-chain

# Create a key to hold your validator account
mezon keys add my-account

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
mezon add-genesis-account $(mezon keys show my-account -a) 1000000000ufis,1000000000validatortoken

# Generate the transaction that creates your validator
mezon gentx my-account 1000000000ufis --chain-id my-chain

# Add the generated bonding transaction to the genesis file
mezon collect-gentxs

# Now its safe to start `mezon`
mezon start
```

This setup puts all the data for `mezon` in `~/.mezonhub`. You can examine the genesis file you created at `~/.mezonhub/config/genesis.json`. With this configuration `mezon` is also ready to use and has an account with tokens (both staking and custom).

## Multi-node, Local, Automated Testnet

From the [networks/local directory](https://github.com/mezonhub/network):

### Requirements

- [Install mezonhub](./install.md)
- [Install docker](https://docs.docker.com/engine/installation/)
- [Install docker-compose](https://docs.docker.com/compose/install/)

### Build

Build the `mezon` binary (linux) and the `tendermint/mezonnode` docker image required for running the `localnet` commands. This binary will be mounted into the container and can be updated without rebuilding the image, so you only need to build the image once.

```bash
# Clone the mezonhub repo
git clone https://github.com/mezonhub/mezonhub.git

# Work from the SDK repo
cd mezonhub

# Build the linux binary in ./build
make build-linux

# Build tendermint/mezonnode image
make build-docker-mezonnode
```

### Run Your Testnet

To start a 4 node testnet run:

```bash
make localnet-start
```

This command creates a 4-node network using the mezonnode image.
The ports for each node are found in this table:

| Node ID     | P2P Port | RPC Port |
| ----------- | -------- | -------- |
| `mezonhubnode0` | `26656`  | `26657`  |
| `mezonhubnode1` | `26659`  | `26660`  |
| `mezonhubnode2` | `26661`  | `26662`  |
| `mezonhubnode3` | `26663`  | `26664`  |

To update the binary, just rebuild it and restart the nodes:

```bash
make build-linux localnet-start
```

### Configuration

The `make localnet-start` creates files for a 4-node testnet in `./build` by
calling the `mezon testnet` command. This outputs a handful of files in the
`./build` directory:

```bash
$ tree -L 2 build/
build/
├── mezon
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   └── mezon
│       ├── key_seed.json
│       ├── keys
│       ├── ${LOG:-mezon.log}
│       ├── config
│       └── data
├── node1
│       ├── key_seed.json
│       ├── ${LOG:-mezon.log}
│       ├── config
│       └── data
├── node2
│       ├── key_seed.json
│       ├── ${LOG:-mezon.log}
│       ├── config
│       └── data
└── node3
         ├── key_seed.json
         ├── ${LOG:-mezon.log}
         ├── config
         └── data
```

Each `./build/nodeN` directory is mounted to the `/mezon` directory in each container.

### Logging

Logs are saved under each `./build/nodeN/mezon/mezonhub.log`. You can also watch logs
directly via Docker, for example:

```bash
docker logs -f mezonnode0
```

### Keys & Accounts

To interact with `mezon` and start querying state or creating txs, you use the
`mezon` directory of any given node as your `home`, for example:

```bash
mezon keys list --home ./build/node0/mezon
```

Now that accounts exists, you may create new accounts and send those accounts
funds!

::: tip
**Note**: Each node's seed is located at `./build/nodeN/mezon/key_seed.json` and can be restored to the CLI using the `mezon keys add --restore` command
:::

### Special Binaries

If you have multiple binaries with different names, you can specify which one to run with the BINARY environment variable. The path of the binary is relative to the attached volume. For example:

```bash
# Run with custom binary
BINARY=mezonhubfoo make localnet-start
```

## Multi-Node, Remote, Automated Testnet

The following should be run from the [networks directory](https://github.com/mezonhub/network).

### Terraform & Ansible

Automated deployments are done using [Terraform](https://www.terraform.io/) to create servers on AWS then
[Ansible](http://www.ansible.com/) to create and manage testnets on those servers.

### Prerequisites

- Install [Terraform](https://www.terraform.io/downloads.html) and [Ansible](http://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html) on a Linux machine.
- Create an [AWS API token](https://docs.aws.amazon.com/general/latest/gr/managing-aws-access-keys.html) with EC2 create capability.
- Create SSH keys

```bash
export AWS_ACCESS_KEY_ID="2345234jk2lh4234"
export AWS_SECRET_ACCESS_KEY="234jhkg234h52kh4g5khg34"
export TESTNET_NAME="remotenet"
export CLUSTER_NAME= "remotenetvalidators"
export SSH_PRIVATE_FILE="$HOME/.ssh/id_rsa"
export SSH_PUBLIC_FILE="$HOME/.ssh/id_rsa.pub"
```

These will be used by both `terraform` and `ansible`.

### Create a Remote Network

```bash
SERVERS=1 REGION_LIMIT=1 make validators-start
```

The testnet name is what's going to be used in --chain-id, while the cluster name is the administrative tag in AWS for the servers. The code will create SERVERS amount of servers in each availability zone up to the number of REGION_LIMITs, starting at us-east-2. (us-east-1 is excluded.) The below BaSH script does the same, but sometimes it's more comfortable for input.

```bash
./new-testnet.sh "$TESTNET_NAME" "$CLUSTER_NAME" 1 1
```

### Quickly see the /status Endpoint

```bash
make validators-status
```

### Delete Servers

```bash
make validators-stop
```

### Logging

You can ship logs to Logz.io, an Elastic stack (Elastic search, Logstash and Kibana) service provider. You can set up your nodes to log there automatically. Create an account and get your API key from the notes on [this page](https://app.logz.io/#/dashboard/data-sources/Filebeat), then:

```bash
yum install systemd-devel || echo "This will only work on RHEL-based systems."
apt-get install libsystemd-dev || echo "This will only work on Debian-based systems."

go get github.com/mheese/journalbeat
ansible-playbook -i inventory/digital_ocean.py -l remotenet logzio.yml -e LOGZIO_TOKEN=ABCDEFGHIJKLMNOPQRSTUVWXYZ012345
```

### Monitoring

You can install the DataDog agent with:

```bash
make datadog-install
```
