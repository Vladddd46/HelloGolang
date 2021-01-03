package views

import "fmt"
import "net/http" 
import "github.com/gorilla/mux"
import "encoding/json"
import "../utils"

/* @ This file contains view functions, 
 *   which are called after user goes 
 *   through registered in UrlRegister 
 *   function url address. Also, there
 *	 are subfunctions for view functions. 
 *	 View functions have `View` postfix. 
 */



// URL: domain/
func IndexView(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "<h1>Hello Go World!</h1>\nGo to: <b>/api/block/<0-9+>/total</b>")
}



/* @ API_GetTotalTransactionsAmountOfEthBlockView subfunction
 *  Makes request in order to get json info about EthBlock
 *	Returns json string.
 */
func GetBlockByNumber(block_number string, api_key string) string {
	var url string
	url = "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag="
	url += block_number
	url += "&boolean=true&apikey="
	url += api_key

	var res string
	resp, err := http.Get(url)  
	if err != nil { 
	    fmt.Println(err) 
	    return ""
	} 
	defer resp.Body.Close()
	for true {  
	    bs := make([]byte, 1014)
	    n, err := resp.Body.Read(bs)
	    res += string(bs[:n])     
	    if n == 0 || err != nil{
	        break
	    }
	}
	return res
}

// @ used in DeserializeJson fucntion
type Transaction_s struct {
	Value  string `json:"value"`
}

// @ used in DeserializeJson fucntion
type DeserializedJsonData_s struct {
	Result struct {
		TransactionsList  []Transaction_s `json:"transactions"`
	}
}

// @ API_GetTotalTransactionsAmountOfEthBlockView subfunction
func DeserializeJson(json_str string) DeserializedJsonData_s {
	var deserialized_json_data DeserializedJsonData_s
	if err := json.Unmarshal([]byte(json_str), &deserialized_json_data); err != nil {
		panic(err)
	}
	return deserialized_json_data
}

// URL: domain/api/block/{block_number:[0-9]+}/total
func API_GetTotalTransactionsAmountOfEthBlockView(page http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	block_number_in_decimal := vars["block_number"]
	block_number_in_hex     := utils.ConvertStrToHex(block_number_in_decimal)
	var requested_json string = GetBlockByNumber(block_number_in_hex, "2GB6GZ1UT7TU42Y9NAGRCV2D7IFCT3BXU5")

	// jsonString := `{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x1d95715bd14","extraData":"0x","gasLimit":"0x2fefd8","gasUsed":"0x5208","hash":"0x7eb7c23a5ac2f2d70aa1ba4e5c56d89de5ac993590e5f6e79c394e290d998ba8","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0xf927a40c8b7f6e07c5af7fa2155b4864a4112b13","mixHash":"0x13dd2c8aec729f75aebcd79a916ecb0f7edc6493efcc6a4da8d7b0ab3ee88444","nonce":"0xc60a782e2e69ce22","number":"0x10d4f","parentHash":"0xf8d01370e6e274f8188954fbee435b40c35b2ad3d4ab671f6d086cd559e48f04","receiptsRoot":"0x0c44b7ed0fefb613ec256341aa0ffdb643e869e3a0ebc8f58e36b4e47efedd33","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x275","stateRoot":"0xd64a0f63e2c7f541e6e6f8548a10a5c4e49fda7ac1aa80f9dddef648c7b9e25f","timestamp":"0x55c9ea07","totalDifficulty":"0x120d56f6821b170","transactions":[{"blockHash":"0x7eb7c23a5ac2f2d70aa1ba4e5c56d89de5ac993590e5f6e79c394e290d998ba8","blockNumber":"0x10d4f","from":"0x4458f86353b4740fe9e09071c23a7437640063c9","gas":"0x5208","gasPrice":"0xba43b7400","hash":"0xa442249820de6be754da81eafbd44a865773e4b23d7c0522d31fd03977823008","input":"0x","nonce":"0x1","to":"0xbf3403210f9802205f426759947a80a9fda71b1e","transactionIndex":"0x0","value":"0xaa9f075c200000","v":"0x1b","r":"0x2c2789c6704ba2606e200e1ba4fd17ba4f0e0f94abe32a12733708c3d3442616","s":"0x2946f47e3ece580b5b5ecb0f8c52604fa5f60aeb4103fc73adcbf6d620f9872b"}],"transactionsRoot":"0x4a5b78c13d11559c9541576834b5172fe8b18507c0f9f76454fcdddedd8dff7a","uncles":[]}}`
	var deserialized_json_data DeserializedJsonData_s = DeserializeJson(requested_json)

	if (len(deserialized_json_data.Result.TransactionsList) > 0) {
		fmt.Println(deserialized_json_data.Result.TransactionsList[0].Value)
	}

	fmt.Fprintf(page, " => %v\n", requested_json)
}





