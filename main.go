package main

import "net/http" 
import "fmt"
import "github.com/gorilla/mux"
import "strconv"


func index(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Hello Go World!\nThis is test!")
}

/* Converts string, which represents number in decimal,
 * 	in string, which represents number in hex.
 */
func ConvertStrToHex(str string) string {
	i, err := strconv.Atoi(str)
	if err != nil {
        fmt.Println(err)
    }
    res := fmt.Sprintf("%x", i)
    return res
}

func router_index(page http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	block_number_in_decimal := vars["block_number"]


	block_number_in_hex := ConvertStrToHex(block_number_in_decimal)


	fmt.Fprintf(page, " => %v\n", block_number_in_hex)
}


func UrlRegister() {
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	router.HandleFunc("/api/block/{block_number:[0-9]+}/total", router_index).Methods("GET")
	http.Handle("/", router)
}

func StartServer(port string) {
	http.ListenAndServe(":" + port, nil)
}

func main() {
 	UrlRegister()
 	StartServer("8080")
}












// resp, err := http.Get("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=0x10d4f&boolean=true&apikey=2GB6GZ1UT7TU42Y9NAGRCV2D7IFCT3BXU5") 
 //    if err != nil { 
 //        fmt.Println(err) 
 //        return
 //    } 
 //    defer resp.Body.Close()
 //    for true {
             
 //        bs := make([]byte, 1014)
 //        n, err := resp.Body.Read(bs)
 //        fmt.Println(string(bs[:n]))
         
 //        if n == 0 || err != nil{
 //            break
 //        }
 //    }