# Websites
[bitcoin.org](https://bitcoin.org/en/)
[developer documentation](https://bitcoin.org/en/developer-documentation)
[developer communities](https://bitcoin.org/en/development#devcommunities)

# Key concepts
- blockchain
- Merkle tree
- satoshis
- coinbase transactions
- proof of work
- hard fork
- soft fork


## Merkle tree
In cryptography and computer science, a hash tree or Merkle tree is a tree in which every non-leaf node is labelled with the hash of the labels or values (in case of leaves) of its child nodes. Hash trees allow efficient and secure verification of the contents of large data structures. Hash trees are a generalization of hash lists and hash chains.

Hash trees can be used to verify any kind of data stored, handled and transferred in and between computers. Currently the main use of hash trees is to make sure that data blocks received from other peers in a peer-to-peer network are received undamaged and unaltered, and even to check that the other peers do not lie and send fake blocks. Suggestions have been made to use hash trees in trusted computing systems.[4] Hash trees are used in the IPFS, Btrfs and ZFS file systems(to counter data degradation), BitTorrent protocol, Apache Wave protocol,[7] Git distributed revision control system, the Tahoe-LAFS backup system, the Bitcoin peer-to-peer network, the Ethereum peer-to-peer network,[8] the Certificate Transparency framework, and a number of NoSQL systems like Apache Cassandra, Riak and DynamoDB.

A hash tree is a tree of hashes in which the leaves are hashes of data blocks in, for instance, a file or set of files. 

Most hash tree implementations are binary (two child nodes under each node) but they can just as well use many more child nodes under each node.

Usually, a cryptographic hash function such as SHA-2 is used for the hashing [according to whom?]. If the hash tree only needs to protect against unintentional damage, much less secure checksums such as CRCs can be used.

In the top of a hash tree there is a top hash (or root hash or master hash). Before downloading a file on a p2p network, in most cases the top hash is acquired from a trusted source, for instance a friend or a web site that is known to have good recommendations of files to download. When the top hash is available, the hash tree can be received from any non-trusted source, like any peer in the p2p network. Then, the received hash tree is checked against the trusted top hash, and if the hash tree is damaged or fake, another hash tree from another source will be tried until the program finds one that matches the top hash

=====================================
The main difference from a hash list is that one branch of the hash tree can be downloaded at a time and the integrity of each branch can be checked immediately, even though the whole tree is not available yet. For example, in the picture, the integrity of data block 2 can be verified immediately if the tree already contains hash 0-0 and hash 1 by hashing the data block and iteratively combining the result with hash 0-0 and then hash 1 and finally comparing the result with the top hash. Similarly, the integrity of data block 3 can be verified if the tree already has hash 1-1 and hash 0. This can be an advantage since it is efficient to split files up in very small data blocks so that only small blocks have to be re-downloaded if they get damaged. If the hashed file is very big, such a hash tree or hash list becomes fairly big. But if it is a tree, one small branch can be downloaded quickly, the integrity of the branch can be checked, and then the downloading of data blocks can start.[
=====================================

second preimage attack
he Merkle hash root does not indicate the tree depth, enabling a second-preimage attack in which an attacker creates a document other than the original that has the same Merkle hash root. For the example above, an attacker can create a new document containing two data blocks, where the first is hash 0-0 + hash 0-1, and the second is hash 1-0 + hash 1-1.

One simple fix is defined in Certificate Transparency: when computing leaf node hashes, a 0x00 byte is prepended to the hash data, while 0x01 is prepended when computing internal node hashes. Limiting the hash tree size is a prerequisite of some formal security proofs, and helps in making some proofs tighter. Some implementations limit the tree depth using hash tree depth prefixes before hashes, so any extracted hash chain is defined to be valid only if the prefix decreases at each step and is still positive when the leaf is reached.

Tiger tree hash
The Tiger tree hash is a widely used form of hash tree. It uses a binary hash tree (two child nodes under each node), usually has a data block size of 1024 bytes and uses the cryptographically secure Tiger hash

Tiger tree hashes are used in Gnutella, Gnutella2, and Direct Connect P2P file sharing protocols and in file sharing applications such as Phex, BearShare, LimeWire, Shareaza, DC++[11] and Valknut




## blockchain

The block chain provides Bitcoin’s public ledger, an ordered and timestamped record of transactions. This system is used to protect against **double spending** and **modification of previous transaction** records.

Each full node in the Bitcoin network independently stores a block chain **containing only blocks validated by that node**. When several nodes all have the same blocks in their block chain, they are considered to be in consensus. The validation rules these nodes follow to maintain consensus are called consensus rules. 

A block of one or more new transactions is collected into the transaction data part of a block. Copies of each transaction are hashed, and the hashes are then paired, hashed, paired again, and hashed again until a single hash remains, the merkle root of a merkle tree.

The merkle root is stored in the block header. Each block also stores the hash of the previous block’s header, chaining the blocks together. This ensures a transaction cannot be modified without modifying the block that records it and all following blocks.

============================
**Transactions are also chained together**. Bitcoin wallet software gives the impression that satoshis are sent from and to wallets, but bitcoins really move from transaction to transaction. Each transaction spends the satoshis previously received in one or more earlier transactions, so the input of one transaction is the output of a previous transaction. 


A single transaction can create multiple outputs, as would be the case when sending to multiple addresses, but each output of a particular transaction can only be used as an input once in the block chain. Any subsequent reference is a forbidden double spend—an attempt to spend the same satoshis twice.

Outputs are tied to transaction identifiers (TXIDs), which are the hashes of signed transactions.

Because each output of a particular transaction can only be spent once, the outputs of all transactions included in the block chain can be categorized as either Unspent Transaction Outputs (UTXOs) or spent transaction outputs. For a payment to be valid, it must only use UTXOs as inputs.

Ignoring coinbase transactions (described later), if the value of a transaction’s outputs exceed its inputs, the transaction will be rejected—but if the inputs exceed the value of the outputs, any difference in value may be claimed as a transaction fee by the Bitcoin miner who creates the block containing that transaction. For example, in the illustration above, each transaction spends 10,000 satoshis fewer than it receives from its combined inputs, effectively paying a 10,000 satoshi transaction fee.
============================

## Proof of work
The block chain is collaboratively maintained by anonymous peers on the network, so Bitcoin requires that each block prove a significant amount of work was invested in its creation to ensure that untrustworthy peers who want to modify past blocks have to work harder than honest peers who only want to add new blocks to the block chain.

Chaining blocks together makes it impossible to modify transactions included in any block without modifying all following blocks. As a result, the cost to modify a particular block increases with every new block added to the block chain, magnifying the effect of the proof of work.

