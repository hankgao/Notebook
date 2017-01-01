## 3rd party packages 

### [boltdbweb](https://github.com/evnix/boltdbweb)
A simple web based boltdb GUI Admin panel.

Install

	$ go get github.com/gin-gonic/gin
	$ go get github.com/boltdb/bolt/...
	$ go get github.com/evnix/boltdbweb
	$ cd $GOPATH/src/github.com/evnix/boltdbweb
	$ go build boltdbweb.go
	$ sudo mv boltdbweb /usr/bin (OPTIONAL)


### [boltdb/bolt](https://github.com/boltdb/bolt/blob/master/README.md)

Installing
`go get github.com/boltdb/bolt`

#### important concepts

- page 
	type page struct {
		id       pgid
		flags    uint16
		count    uint16
		overflow uint32
		ptr      uintptr
	}
every page has a pageid, and a pointer pointing to page data. 


- meta page

- memory mapping

- free page

- page pool

- node
represents an in-memory, deserialized page
	type node struct {
		bucket     *Bucket
		isLeaf     bool
		unbalanced bool
		spilled    bool
		key        []byte
		pgid       pgid
		parent     *node
		children   nodes
		inodes     inodes
	}

- Cursor
ursor represents an iterator that can traverse over all key/value pairs in a bucket in sorted order.
Cursors see nested buckets with value == nil.
Cursors can be obtained from a transaction and are valid as long as the transaction is open.

Keys and values returned from the cursor are only valid for the life of the transaction.

Changing data while traversing with a cursor may cause it to be invalidated
and return unexpected keys and/or values. You must reposition your cursor
after mutating data.

- Bucket
Bucket represents a collection of key/value pairs inside the database.

- bucket
bucket represents the on-file representation of a bucket. This is stored as the "value" of a bucket key. If the bucket is small enough, then its root page can be stored inline in the "value", after the bucket header. In the case of inline buckets, the "root" will be 0.

- transaction
Tx represents a read-only or read/write transaction on the database. Read-only transactions can be used for retrieving values for keys and creating cursors. Read/write transactions can create and remove buckets and create and remove keys.

IMPORTANT: You must commit or rollback transactions when you are done with them. Pages can not be reclaimed by the writer until no more transactions are using them. A long running read transaction can cause the database to quickly grow.

a transaction has one or more buckets 


- DB 
DB represents a collection of buckets persisted to a file on disk. All data access is performed through transactions which can be obtained through the DB. All the functions on DB will return a ErrDatabaseNotOpen if accessed before Open() is called.



Opening a database 

	func main() {
	    // Open the my.db data file in your current directory.
	    // It will be created if it doesn't exist.
	    db, err := bolt.Open("my.db", 0600, nil)
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer db.Close()

	    ...
	}

#### Transactions

##### Read-write transactions
	err := db.Update(func(tx *bolt.Tx) error {
	    ...
	    return nil
	})

##### Read-only transactions
To start a read-only transaction, you can use the DB.View() function:

	err := db.View(func(tx *bolt.Tx) error {
	    ...
	    return nil
	})


##### Batch read-write transactions

Each DB.Update() waits for disk to commit the writes. This overhead can be minimized by combining multiple updates with the DB.Batch() function:

	err := db.Batch(func(tx *bolt.Tx) error {
	    ...
	    return nil
	})


##### using buckets
Buckets are collections of key/value pairs within the database. All keys in a bucket must be unique. You can create a bucket using the DB.CreateBucket() function:

	db.Update(func(tx *bolt.Tx) error {
	    b, err := tx.CreateBucket([]byte("MyBucket"))
	    if err != nil {
	        return fmt.Errorf("create bucket: %s", err)
	    }
	    return nil
	})

### Using key/value pairs

To save a key/value pair to a bucket, use the Bucket.Put() function:

	db.Update(func(tx *bolt.Tx) error {
	    b := tx.Bucket([]byte("MyBucket"))
	    err := b.Put([]byte("answer"), []byte("42"))
	    return err
	})

## Cipher
- Cipher.PubKey
- Cipher.SecKey
- Cipher.AddressFromPubKey
- Cipher.BitcoinAddressFromPubKey
- Cipher.BitcoinWalletImportFormatFromSeckey
- Cipher.SumSHA256
- Cipher.GenerateDeterministicKeyPairs
- Cipher.PubKeyFromSecKey
- Cipher.Sig
- Cipher.Address
- Cipher.PubKey
- Cipher.SecKey
- Cipher.SigFromHex
- Cipher.DecodeBase58Address
- Cipher.PubKeyFromHex
- Cipher.SecKeyFromHex
- Cipher.MustSHA256FromHex
- Cipher.MustDecodeBase58Address
- Cipher.MustSecKeyFromHex
- Cipher.MustDecodeBase58Address

## Wallet
- wallet.Load
- wallet.NewWallet
- wallet.NewWalletEntry
- wallet.NewReadableWalletEntry
- wallet.LoadWallets
- wallet.NewWalletFileName
- wallet.GetAddresses
- wallet.GetAddressSet
- wallet.NewBalance

## CLI
- CLI.DaemonArgs
- CLI.ClientArgs
- CLI.DevArgs

Command
StringFlag
Context
Flag
BoolFlag
IntFlag

## webrpc
- webrpc.Start
- webrpc.NewRequest
- webrpc.Do
- webrpc.InjectResult{}
- webrpc.OutputResult

## util
MustGetLogger
LogConfig{}
DevLogConfig{}
InitDataDir
OpenBrowser
ResolveResourceDirectory
CreateCertIfNotExists
SaveBinary
Now
ZeroTime
UserHome
DisableLogging
UnixNow
SaveJSON
LoadJSON
DetermineResourcePath
DataDir{}
SaveJSONSafe
SaveJSON
CopyFile

## coin
Transaction{}
MustSigFromHex
BlockChain{}
GenerateDetermisticKeyPair
NewBlockChain
Block{}
SorTransactions
NewUnspentPool
NewBlock
CreateUnspent
VerifyTransactionSpending
SortTransactions
BlockHeader{}
TransactionOutput
TransactionDescrialize
SignedBlock
TransactionHeader
HashPair
UxOut
UxHead
UxBody
Balance
Meta
Wallets

## encoder
Serialize
Deserialize
DescrializeRawToValue
SerializeAtomic

## gui
LaunchWebInterfaceHTTPS
LaunchWebInterface
InitWalletRPC

## daemon
Config{}
NewConfig
DefaultConnections{}
NewDaemon
ServiceManager{}
Init
NewBlobReplicator
Start
Run
Gateway{}


