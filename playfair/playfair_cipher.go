package main

import "fmt"

const maxTrip int = 26

type table = map[string][2]int

func incrementChar(c *string) {
	*c = string(int((*c)[0]) + 1)
	if *c > "z" {
		*c = "a"
	}
	if *c == "j" {
		*c = "k"
	}
}

func tableArray(table table) [5][5]string {
	var t [5][5]string
	for key, val := range table {
		t[val[0]][val[1]] = key
	}
	return t
}

func useChar(c string, alreadyUsed map[string]bool) string {
	tripStart := 0
	for {
		if c == "j" {
			c = "i"
		}
		_, ok := alreadyUsed[c]
		if !ok {
			alreadyUsed[c] = true
			return c
		}
		incrementChar(&c)

		tripStart++
		if tripStart == maxTrip {
			panic("Unexpected error")
		}
	}
}

func removeDuplicatesFromString(str *string) {
	place := 0
	uniques := make(map[string]int)
	for _, char := range *str {
		_, ok := uniques[string(char)]
		if !ok {
			uniques[string(char)] = place
			place++
		}
	}

	newString := make([]byte, len(uniques))
	for key, val := range uniques {
		newString[val] = byte(key[0])
	}

	*str = string(newString)
}

func generateTable(key string) table {
	removeDuplicatesFromString(&key)
	table := make(table)
	alreadyUsed := make(map[string]bool)

	tabI, tabJ := 0, 0
	for _, char := range key {
		finalChar := useChar(string(char), alreadyUsed)
		table[finalChar] = [2]int{tabI, tabJ}
		tabJ++
		if tabJ == 5 {
			tabI++
			tabJ = 0
			if tabI == 5 {
				return table
			}
		}
	}

	start := "a"
	for ; tabI < 5; tabI++ {
		for ; tabJ < 5; tabJ++ {
			newChar := useChar(start, alreadyUsed)
			table[newChar] = [2]int{tabI, tabJ}
			incrementChar(&start)
		}
		tabJ = 0
	}

	return table
}

func normalizePlainText(plainText *string) {
	newString := ""
	for ind, char := range *plainText {
		if char < 'a' || char > 'z' {
			continue
		}

		if ind > 0 && string((*plainText)[ind-1]) == string(char) {
			newString += "x"
			continue
		}

		newString += string(char)
	}

	if len(*plainText)%2 != 0 {
		newString += "x"
	}

	*plainText = newString
}

func generateCipher(plainText string, tableMap table, tableArray [5][5]string) string {
	cipher := ""
	limit := len(plainText)

	for i := 0; i < limit; i += 2 {
		char1, char2 := plainText[i], plainText[i+1]

		char1Pos, char2Pos := tableMap[string(char1)], tableMap[string(char2)]

		char1PosRow, char1PosCol := char1Pos[0], char1Pos[1]
		char2PosRow, char2PosCol := char2Pos[0], char2Pos[1]

		if char1PosCol == char2PosCol {
			nxtRow := (char1PosRow + 1) % 5
			cipher += tableArray[nxtRow][char1PosCol]

			nxtRow = (char2PosRow + 1) % 5
			cipher += tableArray[nxtRow][char2PosCol]
		} else if char1PosRow == char2PosRow {
			nxtCol := (char1PosCol + 1) % 5
			cipher += tableArray[char1PosRow][nxtCol]

			nxtCol = (char2PosCol + 1) % 5
			cipher += tableArray[char2PosRow][nxtCol]
		} else {
			cipher += tableArray[char1PosRow][char2PosCol]
			cipher += tableArray[char2PosRow][char1PosCol]
		}
	}

	return cipher
}

// Encrypt - playfair cipher encryption
func Encrypt(plainText string, key string) string {
	normalizePlainText(&plainText)

	tableMap := generateTable(key)
	tableArray := tableArray(tableMap)
	fmt.Println(tableArray)

	return generateCipher(plainText, tableMap, tableArray)
}

func main() {
	var plainText string
	var key string
	fmt.Printf("Enter plain text: ")
	fmt.Scanln(&plainText)
	fmt.Printf("Enter key: ")
	fmt.Scan(&key)

	fmt.Println("Encrypted text:", Encrypt(plainText, key))
}
