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
	var block_number_in_decimal string = vars["block_number"]
	var block_number_in_hex     string = utils.ConvertStrToHex(block_number_in_decimal)
	var requested_json_data     string = GetBlockByNumber(block_number_in_hex, "2GB6GZ1UT7TU42Y9NAGRCV2D7IFCT3BXU5")
	var deserialized_json_data DeserializedJsonData_s = DeserializeJson(requested_json_data)

	for i := 0; i < len(deserialized_json_data.Result.TransactionsList); i++ {
        fmt.Println(deserialized_json_data.Result.TransactionsList[i].Value)
    }

	fmt.Fprintf(page, " => %v\n", requested_json_data)
}





