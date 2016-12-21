
# Factom P2P Networking

### Diagram
![diagram.jpg](https://raw.githubusercontent.com/FactomProject/factomd/m2/p2p/diagram.jpg)

## Conceptual


## Operations

#### Command line options

  -blktime int
    	Seconds per block.  Production is 600.
  -clonedb string
    	Override the main node and use this database for the clones in a Network.
  -count int
    	The number of nodes to generate (default 1)
  -customnet string
    	This string specifies a custom blockchain network ID.
  -db string
    	Override the Database in the Config file and use this Database implementation
  -deadline int
    	Timeout Delay in milliseconds used on Reads and Writes to the network comm (default 1000)
  -drop int
    	Number of messages to drop out of every thousand
  -enablenet
    	Enable or disable networking (default true)
  -exclusive
    	If true, we only dial out to special/trusted peers.
  -fnet string
    	Read the given file to build the network connections
  -follower
    	If true, force node to be a follower.  Only used when replaying a journal.
  -journal string
    	Rerun a Journal of messages
  -journaling
    	Write a journal of all messages recieved. Default is off.
  -keepmismatch
    	If true, do not discard DBStates even when a majority of DBSignatures have a different hash
  -leader
    	If true, force node to be a leader.  Only used when replaying a journal. (default true)
  -net string
    	The default algorithm to build the network connections (default "tree")
  -netdebug int
    	0-5: 0 = quiet, >0 = increasing levels of logging
  -network string
    	Network to join: MAIN, TEST or LOCAL
  -networkPort int
    	Address for p2p network to listen on.
  -node int
    	Node Number the simulator will set as the focus
  -peers string
    	Array of peer addresses. 
  -port int
    	Address to serve WSAPI on
  -prefix string
    	Prefix the Factom Node Names with this value; used to create leaderless networks.
  -rotate
    	If true, responsiblity is owned by one leader, and rotated over the leaders.
  -rpcpass string
    	Password to protect factomd local API. Ignored if rpcuser is blank
  -rpcuser string
    	Username to protect factomd local API with simple HTTP authentication
  -runtimeLog
    	If true, maintain runtime logs of messages passed.
  -selfaddr string
    	comma seperated IPAddresses and DNS names of this factomd to use when creating a cert file
  -startdelay int
    	Delay to start processing messages, in seconds (default 10)
  -test.bench string
    	regular expression per path component to select benchmarks to run
  -test.benchmem
    	print memory allocations for benchmarks
  -test.benchtime duration
    	approximate run time for each benchmark (default 1s)
  -test.blockprofile string
    	write a goroutine blocking profile to the named file after execution
  -test.blockprofilerate int
    	if >= 0, calls runtime.SetBlockProfileRate() (default 1)
  -test.count n
    	run tests and benchmarks n times (default 1)
  -test.coverprofile string
    	write a coverage profile to the named file after execution
  -test.cpu string
    	comma-separated list of number of CPUs to use for each test
  -test.cpuprofile string
    	write a cpu profile to the named file during execution
  -test.memprofile string
    	write a memory profile to the named file after execution
  -test.memprofilerate int
    	if >=0, sets runtime.MemProfileRate
  -test.outputdir string
    	directory in which to write profiles
  -test.parallel int
    	maximum test parallelism (default 8)
  -test.run string
    	regular expression to select tests and examples to run
  -test.short
    	run smaller test suite to save time
  -test.timeout duration
    	if positive, sets an aggregate time limit for all tests
  -test.trace string
    	write an execution trace to the named file after execution
  -test.v
    	verbose: print additional output
  -timedelta int
    	Maximum timeDelta in milliseconds to offset each node.  Simulates deltas in system clocks over a network.
  -tls
    	Set to true to require encrypted connections to factomd API and Control Panel

#### Config file

An example of the config file is below.  Network determines which network we are participating in.  Main is the production blockchain.  TEST is the Testnet.

For each network type you can change the port, seedURL and peers.   The SeedURL is a static file of known trusted peers to connect to.  The special peers are peers you want to always dial out to.

````
; --------------- Network: MAIN | TEST | LOCAL
Network                               = LOCAL
PeersFile            = "peers.json"
MainNetworkPort      = 8108
MainSeedURL          = "https://raw.githubusercontent.com/FactomProject/factomproject.github.io/master/seed/mainseed.txt"
MainSpecialPeers     = ""
TestNetworkPort      = 8109
TestSeedURL          = "https://raw.githubusercontent.com/FactomProject/factomproject.github.io/master/seed/testseed.txt"
TestSpecialPeers     = ""
LocalNetworkPort     = 8110
LocalSeedURL         = "https://raw.githubusercontent.com/FactomProject/factomproject.github.io/master/seed/localseed.txt"
LocalSpecialPeers    = ""
````

Seed file example:
```
1.2.3.4:5678
2.3.4.5:6789
```

## Architecture

App <-> Controller <-> Connection <-> TCP (or UDP in future)

Controller - controller.go
This manages the peers in the network. It keeps connections alive, and routes messages 
from the application to the appropriate peer.  It talks to the application over several
channels. It uses a commandChannel for process isolation, but provides public functions
for all of the commands.  The messages for the network peers go over the ToNetwork and
come in on the FromNetwork channels.

Connection - connection.go
This struct represents an individual connection to another peer. It talks to the 
controller over channels, again providing process/memory isolation. 
