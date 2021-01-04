package utils

import "fmt"
import "strconv"
import "unsafe"



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



/* Converts hex string in float64.
 *** SHOULD BE FIXED IN FUTURE
 */
func HexToFloat(hex_string string) float64 {
	n, err := strconv.ParseUint(hex_string, 16, 64)

	/* !!! All values bigger than float64 max value will be zero.
	 * In future it should be fixed by changing float64 to bigger data types
	 */
	if err != nil {
		return 0
	}
	n2 := uint64(n)
	f := *(*float64)(unsafe.Pointer(&n2))
	return f
}
