
## key concepts
- wallet
- public key
- private(secure) key
- block
- hash
- unspent output
- transaction
- blockchain
- address
- base58 
- checksum

### public key
- you can generate public key from private key

### address
Addresses are the Ripemd160 of the double SHA256 of the public key. 
- public key must be in compressed format

In the block chain the address is 20+1 bytes
- the first byte is the version byte
- the next twenty bytes are RIPMD160(SHA256(SHA256(pubkey)))

//version is after Key to enable better vanity address generation
//Address stuct is a 25 byte with a 20 byte publickey hash, 1 byte address
//type and 4 byte checksum.
type Address struct {
	Version byte      //1 byte
	Key     Ripemd160 //20 byte pubkey hash
}

In base 58 format the address is 20+1+4 bytes
- the first 20 bytes are RIPMD160(SHA256(SHA256(pubkey))).
-- this is to allow for any prefix in vanity addresses
- the next byte is the version byte
- the next 4 bytes are a checksum
-- the first 4 bytes of the SHA256 of the 21 bytes that come before

### checksum
type Checksum [4]byte


### block 
a block is comprised of a blockheader and a blockbody. Blockheader contains version, time stamp, block sequence number, fee, hash of previous block, hash of transaction body

    type BlockHeader struct {
    	Version uint32    

    	Time  uint64
    	BkSeq uint64 //increment every block
    	Fee   uint64 //fee in block, used for Proof of Stake    

    	PrevHash cipher.SHA256 //hash of header of previous block
    	BodyHash cipher.SHA256 //hash of transaction block    

    	UxHash cipher.SHA256 //XOR of sha256 of elements in unspent output set    

    }

### unspent transaction output UTXO
[The Challenges of Optimizing Unspent Output Selection](https://blog.bitgo.com/challenges-optimizing-unspent-output-selection/)
An Unspent Transaction Output (UTXO) that can be spent as an input in a new transaction.

Transactions are composed of inputs and outputs; unspent outputs are the actual bitcoins.

When Bitcoin wallet software creates a new transaction, it has a fair amount of flexibility in how it can arrange the internal data of the transaction. This is because a user instructs the wallet: “send X bitcoins to address Y” and the wallet needs to comb through the user’s UTXOs in order to find enough outputs with an aggregate value that reaches the target of X bitcoins.




## folder strucuture

- cmd
- electron
- scripts
- src

where `cmd` contains ...
`electron` contains ...
`scripts` contains ... 
`src` contains all the packages that skycoin commands builds upon. 

### cmd
cmd contains three commands, address_gen, mesh, and skycoin, of which skycoin is the core. 

### src
src contains the following packages: 
- aether
- cipher
- coin
- consensus
- dameon
- gui
- mesh
- util
- visor
- wallet

#### cipher
cipher again contains the following packages:
- base58
- chacha20
- encoder
- ripemd160
- secp256k1-go

and 4 sources files: address.go, chacha20.go, crypto.go, and hash.go

These packages are all different cryptogprhy algorithms. Skycoin doesn't invent them, they are just GOlang implementions of popular algorithms. 

Hank: seems to me that in Skycoin, cryptography algorithms are pluggable.



#### coin
coin contains the following source files： block.go, outputs.go, transactions.go, unspent_pool.go. The names of these files can tell a lot about the key things about a cryptocurrency. 

The key exported functions of block.go are NewBlock, and CreateUnspents. It also exports the following key structures: Block, BlockHeader, BlockBody, where BlockBody contains transactions and BlockHeader contains Version, Time, BkSeq, Fee, PreHash, BodyHash, and UxHash

The key exported functions of transactions.go are Verify, VerifyInput, PushInput, PushOutput, SignInputs, Serialize 



#### gui
#### wallet
#### util



### scripts 

### electron

