//Package morestrings implements additional functions to manipulate UTF-8
//encode strings, beyond what is provided in the standard "string" package.
package morestrings

//ReverseRunes returns its argument string reversed run-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i,j := 0, len(r)-1; i < len(r)/2; i,j = i +1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}