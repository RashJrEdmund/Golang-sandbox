// by convention, we name our package the same as the directory
package mystrings

// Reverse reverses a string left to right
// Export functions by capitalizing the first letter of it's name
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
