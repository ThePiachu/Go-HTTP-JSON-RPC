package BitcoinJSON

import(
	"httpjsonrpc"
	"os"
	"log"
)
var Address string

func SetAddress(newAddress string){
	Address=newAddress
}

/*
//https://en.bitcoin.it/wiki/Api#Full_list\
*/

func BackupWallet(id interface{}, destination []interface{})(map[string]interface{}, os.Error){
	//Safely copies wallet.dat to destination, which can be a directory or a path with filename.
	resp, err:=httpjsonrpc.Call(Address, "backupwallet", id, destination)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func EncryptWallet(id interface{}, passphrase []interface{})(map[string]interface{}, os.Error){
	//Encrypts the wallet with <passphrase>.
	resp, err:=httpjsonrpc.Call(Address, "encryptwallet", id, passphrase)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetAccount(id interface{}, bitcoinaddress []interface{})(map[string]interface{}, os.Error){
	//Returns the account associated with the given address.
	resp, err:=httpjsonrpc.Call(Address, "getaccount", id, bitcoinaddress)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetAccountAddress(id interface{}, account []interface{})(map[string]interface{}, os.Error){
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err:=httpjsonrpc.Call(Address, "getaccountaddress", id, account)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetAddressByAccount(id interface{}, account []interface{})(map[string]interface{}, os.Error){
	//Returns the list of addresses for the given account.
	resp, err:=httpjsonrpc.Call(Address, "getaddressesbyaccount", id, account)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetBalance(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//If [account] is not specified, returns the server's total available balance.
	//If [account] is specified, returns the balance in the account.
	resp, err:=httpjsonrpc.Call(Address, "getbalance", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

/*func GetBlockByCount(id interface{}, height []interface{})(map[string]interface{}, os.Error){
	//Dumps the block existing at specified height. Note: this is not available in the official release
	resp, err:=httpjsonrpc.Call(Address, "getblockbycount", id, height)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}*/

func GetBlockCount(id interface{})(map[string]interface{}, os.Error){
	//Returns the number of blocks in the longest block chain.
	resp, err:=httpjsonrpc.Call(Address, "getblockcount", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}








func GetBlockNumber(id interface{})(map[string]interface{}, os.Error){
	//Returns the block number of the latest block in the longest block chain.
	resp, err:=httpjsonrpc.Call(Address, "getblocknumber", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetConnectionCount(id interface{})(map[string]interface{}, os.Error){
	//Returns the number of connections to other nodes.
	resp, err:=httpjsonrpc.Call(Address, "getconnectioncount", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetDifficulty(id interface{})(map[string]interface{}, os.Error){
	//Returns the proof-of-work difficulty as a multiple of the minimum difficulty.
	resp, err:=httpjsonrpc.Call(Address, "getdifficulty", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetGenerate(id interface{})(map[string]interface{}, os.Error){
	//Returns true or false whether bitcoind is currently generating hashes
	resp, err:=httpjsonrpc.Call(Address, "getgenerate", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetHashesPerSec(id interface{})(map[string]interface{}, os.Error){
	//Returns a recent hashes per second performance measurement while generating.
	resp, err:=httpjsonrpc.Call(Address, "gethashespersec", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetInfo(id interface{})(map[string]interface{}, os.Error){
	//Returns an object containing various state info.
	resp, err:=httpjsonrpc.Call(Address, "getinfo", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetMemoryPool(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//If [data] is not specified, returns data needed to construct a block to work on:
  	//"version" : block version
  	//"previousblockhash" : hash of current highest block
  	//"transactions" : contents of non-coinbase transactions that should be included in the next block
  	//"coinbasevalue" : maximum allowable input to coinbase transaction, including the generation award and transaction fees
  	//"time" : timestamp appropriate for next block
  	//"bits" : compressed target of next block
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err:=httpjsonrpc.Call(Address, "getmemorypool", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetNewAddress(id interface{}, account []interface{})(map[string]interface{}, os.Error){
	//Returns a new bitcoin address for receiving payments. If [account] is specified (recommended), it is added to the address book so payments received with the address will be credited to [account].
	resp, err:=httpjsonrpc.Call(Address, "getnewaddress", id, account)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetReceivedByAccount(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns the total amount received by addresses with [account] in transactions with at least [minconf] confirmations. If [account] not provided return will include all transactions to all accounts. (version 0.3.24-beta)
	resp, err:=httpjsonrpc.Call(Address, "getreceivedbyaccount", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetReceivedByAddress(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns the total amount received by <bitcoinaddress> in transactions with at least [minconf] confirmations. While some might consider this obvious, value reported by this only considers *receiving* transactions. It does not check payments that have been made *from* this address. In other words, this is not "getaddressbalance".
	resp, err:=httpjsonrpc.Call(Address, "getreceivedbyaddress", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetTransaction(id interface{}, txid []interface{})(map[string]interface{}, os.Error){
	//Returns an object about the given transaction containing:
	//"amount" : total amount of the transaction
	//"confirmations" : number of confirmations of the transaction
	//"txid" : the transaction ID
	//"time" : time the transaction occurred
	//"details" - An array of objects containing:
	//"account"
	//"address"
	//"category"
	//"amount"
	resp, err:=httpjsonrpc.Call(Address, "gettransaction", id, txid)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetWork(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//If [data] is not specified, returns formatted hash data to work on:
	//"midstate" : precomputed hash state after hashing the first half of the data
	//"data" : block data
	//"hash1" : formatted hash buffer for second hash
	//"target" : little endian hash target
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err:=httpjsonrpc.Call(Address, "getwork", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	//result:=resp["result"]
	//log.Println(result)
	
	return resp, err
}

func Help(id interface{}, command string)(map[string]interface{}, os.Error){
	//List commands, or get help for a command.
	resp, err:=httpjsonrpc.Call(Address, "help", id, []interface{}{command})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func KeyPoolRefill(id interface{})(map[string]interface{}, os.Error){
	//Fills the keypool, requires wallet passphrase to be set.
	resp, err:=httpjsonrpc.Call(Address, "keypoolrefill", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListAccounts(id interface{}, minconf interface{})(map[string]interface{}, os.Error){
	//Returns Object that has account names as keys, account balances as values.
	resp, err:=httpjsonrpc.Call(Address, "listaccounts", id, []interface{}{minconf})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListReceivedByAccount(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns an array of objects containing:
	//"account" : the account of the receiving addresses
	//"amount" : total amount received by addresses with this account
	//"confirmations" : number of confirmations of the most recent transaction included
	resp, err:=httpjsonrpc.Call(Address, "listreceivedbyaccount", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListReceivedByAddress(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns an array of objects containing:
	//"address" : receiving address
	//"account" : the account of the receiving address
	//"amount" : total amount received by the address
	//"confirmations" : number of confirmations of the most recent transaction included
	//To get a list of accounts on the system, execute bitcoind listreceivedbyaddress 0 true
	resp, err:=httpjsonrpc.Call(Address, "listreceivedbyaddress", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListSinceBlock(id interface{}, blockid, targetconfirmations interface{})(map[string]interface{}, os.Error){
	//Get all transactions in blocks since block [blockid], or all transactions if omitted
	resp, err:=httpjsonrpc.Call(Address, "listsinceblock", id, []interface{}{blockid, targetconfirmations})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListTransactions(id interface{}, account, count, from interface{})(map[string]interface{}, os.Error){
	//Returns up to [count] most recent transactions skipping the first [from] transactions for account [account]. If [account] not provided will return recent transaction from all accounts.
	resp, err:=httpjsonrpc.Call(Address, "listtransactions", id, []interface{}{account, count, from})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func Move(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Move from one account in your wallet to another
	resp, err:=httpjsonrpc.Call(Address, "move", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SendFrom(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//<amount> is a real and is rounded to 8 decimal places. Will send the given amount to the given address, ensuring the account has a valid balance using [minconf] confirmations. Returns the transaction ID if successful (not in JSON object).
	resp, err:=httpjsonrpc.Call(Address, "sendfrom", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SendMany(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//amounts are double-precision floating point numbers
	resp, err:=httpjsonrpc.Call(Address, "sendmany", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SendToAddress(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//<amount> is a real and is rounded to 8 decimal places. Returns the transaction ID <txid> if successful.
	resp, err:=httpjsonrpc.Call(Address, "sendtoaddress", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SetAccount(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Sets the account associated with the given address. Assigning address that is already assigned to the same account will create a new address associated with that account.
	resp, err:=httpjsonrpc.Call(Address, "setaccount", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SetGenerate(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//<generate> is true or false to turn generation on or off.
	//Generation is limited to [genproclimit] processors, -1 is unlimited.
	resp, err:=httpjsonrpc.Call(Address, "setgenerate", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SetTxFee(id interface{}, amount []interface{})(map[string]interface{}, os.Error){
	//<amount> is a real and is rounded to the nearest 0.00000001
	resp, err:=httpjsonrpc.Call(Address, "settxfee", id, amount)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SignMessage(id interface{}, bitcoinaddress, message interface{})(map[string]interface{}, os.Error){
	//Sign a message with the private key of an address
	resp, err:=httpjsonrpc.Call(Address, "signmessage", id, []interface{}{bitcoinaddress, message})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func Stop(id interface{})(map[string]interface{}, os.Error){
	//Stop bitcoin server.
	resp, err:=httpjsonrpc.Call(Address, "stop", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ValidateAddress(id interface{}, bitcoinaddress interface{})(map[string]interface{}, os.Error){
	//Return information about <bitcoinaddress>.
	resp, err:=httpjsonrpc.Call(Address, "validateaddress", id, []interface{}{bitcoinaddress})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func VerifyMessage(id interface{}, bitcoinaddress, signature, message interface{})(map[string]interface{}, os.Error){
	//Verify a signed message
	resp, err:=httpjsonrpc.Call(Address, "validateaddress", id, []interface{}{bitcoinaddress, signature, message})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func WalletLock(id interface{})(map[string]interface{}, os.Error){
	//Removes the wallet encryption key from memory, locking the wallet. After calling this method, you will need to call walletpassphrase again before being able to call any methods which require the wallet to be unlocked.
	resp, err:=httpjsonrpc.Call(Address, "walletlock", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func WalletPassPhrase(id interface{}, passphrase, timeout interface{})(map[string]interface{}, os.Error){
	//Stores the wallet decryption key in memory for <timeout> seconds.
	resp, err:=httpjsonrpc.Call(Address, "walletpassphrase", id, []interface{}{passphrase, timeout})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func WalletPassPhraseChange(id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Changes the wallet passphrase from <oldpassphrase> to <newpassphrase>.
	resp, err:=httpjsonrpc.Call(Address, "walletpassphrasechange", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}


