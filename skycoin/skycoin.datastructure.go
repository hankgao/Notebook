// # datat structures 
// - Sig // signature 

type PubKey [33]byte
type SecKey [32]byte

// ## 
type Config struct {
	// Disable peer exchange
	DisablePEX: true,
	// Don't make any outgoing connections
	DisableOutgoingConnections: false,
	// Don't allowing incoming connections
	DisableIncomingConnections: false,
	// Disables networking altogether
	DisableNetworking: false,
	// Only run on localhost and only connect to others on localhost
	LocalhostOnly: false,
	// Which address to serve on. Leave blank to automatically assign to a
	// public interface
	Address: "",
	//gnet uses this for TCP incoming and outgoing
	Port: 6000,

	MaxConnections: 16,
	// How often to make outgoing connections, in seconds
	OutgoingConnectionsRate: time.Second * 5,
	// Wallet Address Version
	//AddressVersion: "test",
	// Remote web interface
	WebInterface:             true,
	WebInterfacePort:         6420,
	WebInterfaceAddr:         "127.0.0.1",
	WebInterfaceCert:         "",
	WebInterfaceKey:          "",
	WebInterfaceHTTPS:        false,
	PrintWebInterfaceAddress: false,
	LaunchBrowser:            true,
	// Data directory holds app data -- defaults to ~/.skycoin
	DataDirectory: ".skycoin",
	// Web GUI static resources
	GUIDirectory: "./src/gui/static/",
	// Logging
	LogLevel: logging.DEBUG,
	ColorLog: true,
	logLevel: "DEBUG",

	// Wallets
	WalletDirectory: "",
	BlockchainFile:  "",
	BlockSigsFile:   "",

	// Centralized network configuration
	RunMaster:        false,
	BlockchainPubkey: cipher.PubKey{},
	BlockchainSeckey: cipher.SecKey{},

	GenesisAddress:   cipher.Address{},
	GenesisTimestamp: GenesisTimestamp,
	GenesisSignature: cipher.Sig{},

	/* Developer options */

	// Enable cpu profiling
	ProfileCPU: false,
	// Where the file is written to
	ProfileCPUFile: "skycoin.prof",
	// HTTP profiling interface (see http://golang.org/pkg/net/http/pprof/)
	HTTPProf: false,
	// Will force it to connect to this ip:port, instead of waiting for it
	// to show up as a peer
	ConnectTo: "",
} 

//## Transaction

// pacakge coin
type Transaction struct {
	Length    uint32        //length prefix
	Type      uint8         //transaction type
	InnerHash cipher.SHA256 //inner hash SHA256 of In[],Out[]

	Sigs []cipher.Sig        //list of signatures, 64+1 bytes each
	In   []cipher.SHA256     //ouputs being spent
	Out  []TransactionOutput //ouputs being created
}

//cipher.SHA256
type SHA256 [32]byte

//cipher.Sig
type Sig [64 + 1]byte //64 byte signature with 1 byte for key recovery

//hash output/name is function of Hash
type TransactionOutput struct {
	Address cipher.Address //address to send to
	Coins   uint64         //amount to be sent in coins
	Hours   uint64         //amount to be sent in coin hours
}

//version is after Key to enable better vanity address generation
//Address stuct is a 25 byte with a 20 byte publickey hash, 1 byte address
//type and 4 byte checksum.
// cipher.Address
type Address struct {
	Version byte      //1 byte
	Key     Ripemd160 //20 byte pubkey hash
}

type Ripemd160 [20]byte