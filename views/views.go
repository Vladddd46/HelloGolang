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
		// panic(err)
		/* must be handled in future */
	}
	return deserialized_json_data
}

/* @ API_GetTotalTransactionsAmountOfEthBlockView subfunction
 * Note: This function has defect in utils.HexToFloat (values > float64 will be respresented as 0)
 */
func CountTotalValueOfTransactions(deserialized_json_data DeserializedJsonData_s, num_of_transactions int) float64 {
	var EthInWei = 1.6861753e-10 // Wei is 1✕10​**-1 Ether

	var total int64 = 0
	for i := 0; i < num_of_transactions; i++ {
        value_field := deserialized_json_data.Result.TransactionsList[i].Value
        value_field_without_0x := value_field[2: len(value_field)]

        // Value field is represented in Wei but total must be represented in Ether => multiply by EthInWei
        total += utils.HexToInt(value_field_without_0x)
    }
    var res float64 = float64(total) * EthInWei
    return res
}

// URL: domain/api/block/{block_number:[0-9]+}/total
func API_GetTotalTransactionsAmountOfEthBlockView(page http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var block_number_in_decimal string = vars["block_number"]
	var block_number_in_hex     string = utils.ConvertStrToHex(block_number_in_decimal)
	var api_key string = utils.GetApiKey()
	var requested_json_data     string = GetBlockByNumber(block_number_in_hex, api_key)
	var deserialized_json_data DeserializedJsonData_s = DeserializeJson(requested_json_data)

	var num_of_transactions = len(deserialized_json_data.Result.TransactionsList)
	var total               = CountTotalValueOfTransactions(deserialized_json_data, num_of_transactions)
	
	// fmt.Println(num_of_transactions, total) // Debug log
	var sendback_data string = fmt.Sprintf(`{"transactions": %d, "amount": %e}`, num_of_transactions, total);
	fmt.Fprintf(page, "%v\n", sendback_data)
}





