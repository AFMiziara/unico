package cpfcnpj

import (
	"strconv"
	"strings"
)

const BlacklistCnpj = `00000000000000
11111111111111
22222222222222
33333333333333
44444444444444
55555555555555
66666666666666
77777777777777
88888888888888
99999999999999`

const BlacklistCpf = `00000000000
11111111111
22222222222
33333333333
44444444444
55555555555
66666666666
77777777777
88888888888
99999999999`

func sumDigit(s string, table []int) int {

	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

// Clean can be used for cleaning formatted documents
func Clean(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}
