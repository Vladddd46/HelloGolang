package utils

import "fmt"
import "strconv"

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


