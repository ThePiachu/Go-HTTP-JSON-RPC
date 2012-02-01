package GAEBitcoinJSON

import(
	"gaehttpjsonrpc"
	"os"
	"log"
	"appengine"
)
var Address string
var Username string
var Password string

func SetAddress(newAddress string, newUsername string, newPassword string){
	Address=newAddress
	Username=newUsername
	Password=newPassword
}

/*
//https://en.bitcoin.it/wiki/Api#Full_list\
*/

func BackupWallet(c appengine.Context, id interface{}, destination []interface{})(map[string]interface{}, os.Error){
	//Safely copies wallet.dat to destination, which can be a directory or a path with filename.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "backupwallet", id, destination)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func EncryptWallet(c appengine.Context, id interface{}, passphrase []interface{})(map[string]interface{}, os.Error){
	//Encrypts the wallet with <passphrase>.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "encryptwallet", id, passphrase)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetAccount(c appengine.Context, id interface{}, bitcoinaddress []interface{})(map[string]interface{}, os.Error){
	//Returns the account associated with the given address.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getaccount", id, bitcoinaddress)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetAccountAddress(c appengine.Context, id interface{}, account []interface{})(map[string]interface{}, os.Error){
	//Returns the current bitcoin address for receiving payments to this account.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getaccountaddress", id, account)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetAddressByAccount(c appengine.Context, id interface{}, account []interface{})(map[string]interface{}, os.Error){
	//Returns the list of addresses for the given account.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getaddressesbyaccount", id, account)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetBalance(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//If [account] is not specified, returns the server's total available balance.
	//If [account] is specified, returns the balance in the account.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getbalance", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

/*func GetBlockByCount(c appengine.Context, id interface{}, height []interface{})(map[string]interface{}, os.Error){
	//Dumps the block existing at specified height. Note: this is not available in the official release
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getblockbycount", id, height)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}*/

func GetBlockCount(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns the number of blocks in the longest block chain.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getblockcount", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}








func GetBlockNumber(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns the block number of the latest block in the longest block chain.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getblocknumber", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetConnectionCount(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns the number of connections to other nodes.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getconnectioncount", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetDifficulty(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns the proof-of-work difficulty as a multiple of the minimum difficulty.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getdifficulty", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetGenerate(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns true or false whether bitcoind is currently generating hashes
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getgenerate", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetHashesPerSec(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns a recent hashes per second performance measurement while generating.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "gethashespersec", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func GetInfo(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Returns an object containing various state info.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getinfo", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetNewAddress(c appengine.Context, id interface{}, account []interface{})(map[string]interface{}, os.Error){
	//Returns a new bitcoin address for receiving payments. If [account] is specified (recommended), it is added to the address book so payments received with the address will be credited to [account].
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getnewaddress", id, account)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetReceivedByAccount(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns the total amount received by addresses with [account] in transactions with at least [minconf] confirmations. If [account] not provided return will include all transactions to all accounts. (version 0.3.24-beta)
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getreceivedbyaccount", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetReceivedByAddress(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns the total amount received by <bitcoinaddress> in transactions with at least [minconf] confirmations. While some might consider this obvious, value reported by this only considers *receiving* transactions. It does not check payments that have been made *from* this address. In other words, this is not "getaddressbalance".
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getreceivedbyaddress", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetTransaction(c appengine.Context, id interface{}, txid []interface{})(map[string]interface{}, os.Error){
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
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "gettransaction", id, txid)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func GetWork(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//If [data] is not specified, returns formatted hash data to work on:
	//"midstate" : precomputed hash state after hashing the first half of the data
	//"data" : block data
	//"hash1" : formatted hash buffer for second hash
	//"target" : little endian hash target
	//If [data] is specified, tries to solve the block and returns true if it was successful.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "getwork", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	//result:=resp["result"]
	//log.Println(result)
	
	return resp, err
}

func Help(c appengine.Context, id interface{}, command string)(map[string]interface{}, os.Error){
	//List commands, or get help for a command.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "help", id, []interface{}{command})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func KeyPoolRefill(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Fills the keypool, requires wallet passphrase to be set.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "keypoolrefill", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListAccounts(c appengine.Context, id interface{}, minconf []interface{})(map[string]interface{}, os.Error){
	//Returns Object that has account names as keys, account balances as values.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listaccounts", id, minconf)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListReceivedByAccount(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns an array of objects containing:
	//"account" : the account of the receiving addresses
	//"amount" : total amount received by addresses with this account
	//"confirmations" : number of confirmations of the most recent transaction included
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listreceivedbyaccount", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListReceivedByAddress(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Returns an array of objects containing:
	//"address" : receiving address
	//"account" : the account of the receiving address
	//"amount" : total amount received by the address
	//"confirmations" : number of confirmations of the most recent transaction included
	//To get a list of accounts on the system, execute bitcoind listreceivedbyaddress 0 true
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listreceivedbyaddress", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ListTransactions(c appengine.Context, id interface{}, account, count, from interface{})(map[string]interface{}, os.Error){
	//Returns up to [count] most recent transactions skipping the first [from] transactions for account [account]. If [account] not provided will return recent transaction from all accounts.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "listtransactions", id, []interface{}{account, count, from})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func Move(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Move from one account in your wallet to another
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "move", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SendFrom(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//<amount> is a real and is rounded to 8 decimal places. Will send the given amount to the given address, ensuring the account has a valid balance using [minconf] confirmations. Returns the transaction ID if successful (not in JSON object).
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "sendfrom", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SendMany(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//amounts are double-precision floating point numbers
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "sendmany", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SendToAddress(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//<amount> is a real and is rounded to 8 decimal places. Returns the transaction ID <txid> if successful.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "sendtoaddress", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SetAccount(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Sets the account associated with the given address. Assigning address that is already assigned to the same account will create a new address associated with that account.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "setaccount", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SetGenerate(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//<generate> is true or false to turn generation on or off.
	//Generation is limited to [genproclimit] processors, -1 is unlimited.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "setgenerate", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func SetTxFee(c appengine.Context, id interface{}, amount []interface{})(map[string]interface{}, os.Error){
	//<amount> is a real and is rounded to the nearest 0.00000001
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "settxfee", id, amount)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func Stop(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Stop bitcoin server.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "stop", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func ValidateAddress(c appengine.Context, id interface{}, bitcoinaddress interface{})(map[string]interface{}, os.Error){
	//Return information about <bitcoinaddress>.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "validateaddress", id, []interface{}{bitcoinaddress})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
func WalletLock(c appengine.Context, id interface{})(map[string]interface{}, os.Error){
	//Removes the wallet encryption key from memory, locking the wallet. After calling this method, you will need to call walletpassphrase again before being able to call any methods which require the wallet to be unlocked.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "walletlock", id, nil)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func WalletPassPhrase(c appengine.Context, id interface{}, passphrase, timeout interface{})(map[string]interface{}, os.Error){
	//Stores the wallet decryption key in memory for <timeout> seconds.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "walletpassphrase", id, []interface{}{passphrase, timeout})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}

func WalletPassPhraseChange(c appengine.Context, id interface{}, data []interface{})(map[string]interface{}, os.Error){
	//Changes the wallet passphrase from <oldpassphrase> to <newpassphrase>.
	resp, err:=gaehttpjsonrpc.CallWithBasicAuth(c, Address, Username, Password, true, "walletpassphrasechange", id, data)
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}


