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

func GetInfo()(map[string]interface{}, os.Error){
	resp, err:=httpjsonrpc.Call(Address, "getinfo", 1, []interface{}{})
	if err!=nil{
		log.Println(err)
		return resp, err
	}
	result:=resp["result"]
	log.Println(result)
	
	return resp, err
}
