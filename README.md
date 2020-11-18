# Go Ethereum Mod For Experiment v1.9.20

## Feature
- Add log feature. The following event can be recorded separately: when a block is created, when a block is sealed, when a transaction is being executed. It is useful to track the correct time cost for executing a transaction, without consensus time or network latency.
- The gas limit will never change. It will always keep the current block gas limit, so just specify a large gas limit like `0x1fffffffffffff` in the genesis block file and don't need to worry about the gas. 
- Add full implementation of bn256 curve as precompiled contracts. The original ethereum client only support a subset of bn256 curve which is specially designed for zk-SNARK only, which is not suitable for other kind of experiments. 

## Compile
Steps are the same as compiling the official Go Ethereum.

1. Install `git`, `gcc` and `go`.
2. Run `make geth` to compile.

## Run

Currently, the log feature is mandatory, so you must specify a path to the log file.

Let's say the log file is `/path/to/file.txt`. Add `experiment.output` flag in the front of the argument.

Example:

```shell
geth --experiment.output=/path/to/file.txt
```

```shell
geth --experiment.output=/path/to/file.txt --otherarguments...
```

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
