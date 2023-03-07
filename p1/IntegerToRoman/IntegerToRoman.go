func intToRoman(n int) string {
	var output string
	 for n >= 1000 {
		 output += "M"
		 n -= 1000
	 }
 
	 for n>= 900 {
		 output += "CM"
		 n -= 900
	 }
 
	 for n>= 500 {
		 output += "D"
		 n -= 500
	 }
 
	 for n>= 400 {
		 output += "CD"
		 n -= 400
	 }
 
	 for n>= 100 {
		 output += "C"
		 n -= 100
	 }
 
	 for n>= 90 {
		 output += "XC"
		 n -= 90
	 }
 
	 for n>= 50 {
		 output += "L"
		 n -= 50
	 }
 
	 for n>= 40 {
		 output += "XL"
		 n -= 40
	 }
 
	 for n>= 10 {
		 output += "X"
		 n -= 10
	 }
 
	 for n>= 9 {
		 output += "IX"
		 n -= 9
	 }
 
	 for n>= 5 {
		 output += "V"
		 n -= 5
	 }
 
	 for n>= 4 {
		 output += "IV"
		 n -= 4
	 }
 
	 for n>= 1 {
		 output += "I"
		 n -= 1
	 }
 
	 return output
 }