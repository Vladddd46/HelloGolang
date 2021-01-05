package utils

import "fmt"
import "strconv"
import "io/ioutil"
import "log"
import "math/big"

/* @ This file contains general-purpose functions
 */



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



func HexToBigFloat(hex_string string) *big.Float {
    distance := new(big.Float)
    distance.SetString(hex_string)
    return distance
}



// Gets API key from config file.
func GetApiKey() string {
	content, err := ioutil.ReadFile("api_key.txt")
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}


