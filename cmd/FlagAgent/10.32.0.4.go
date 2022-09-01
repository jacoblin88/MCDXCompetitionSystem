package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Annt92X0rEJ6K9C6VWbf2a7j96su+R+qq5Zq831ZkSijVNjosOXaA61gtTtuzBuzOWPMikDNhAyuh9QQGc1AVhBdtOydnbIPX/Yz6a9ioFLMA5JENzwhYyQrTVZe9Jt+RbL8Bstup/lK2VSB3rn5sKYHgHUKU0Oy51sphMSF0g5Nb8hWmASXhi1sYpDDq4Keshrgk+M/ZM/kqeHG6UPG2B7lHKyyRzZdJQxbBwNfHz6ErYj0cOPqH3hLOAETs+GutYcFlTj5IHVwGfDz+QKhyBt7yRqnCGGyNuxU4PuDatSSqRljxm6gbKKeJDs/7/v3Wjtoi1/UKRrCdQstBfgSOsE=\", \"E\": \"AQABAA==\", \"N\": \"AsO4lABgbLHp3jWPoEyUYbT+XFeft3UtJAxHyxTS3sgffFos2xjSvRsKH7pmoLvNEVMhRSGU729zRlFcnC/RRSb5YwIyrYHTN+PVXKTojC+HwJW7MAguJOdV9iDu1pteYvoREO0ZIJRlV83O6Xn4fgEawHLWZI/NTkPfy7gic7+ySNLP7ZL1f++srwMKWrm/XVN1c6CU43rB9UsA3MeAoZVTz+qiuIhHWubQ6Uxh2Xv9zmyTexwG9s0oy4OBnDkWBea2sNy5n+495X+l3N2Ukjo5nUOC+KMqbrJsnl+y6clPOltCjVMa7+sLjZ4jJ4KT7bJR7FRI5Y1PdjokpKWAnfk=\", \"Primes1\": \"Au8MZ6J0j5pFO89n0tqbV3LRewji9OsFJbX6N2AIRh9TAVYlbhVezc07Um1oZYK0DgxaBrNjbuE7UjsKVp5sby44i1lO58YVMMB4mR8Tt5gq8/6+LOutTouSSaaYTjAfuECiyCzxlC6b8msTx8xjTwdLmRmzv7C028+DwKbWzvMd\", \"Primes2\": \"AtGZn4tUrQJNWZLNNsV6XQ0vRRcfUbfONi6r822uf7uIY+jXi9gNrtqVWORJeXNpu8Pko74c5gOllcDteDWSC/LBvOjbNFWbDEu4fBAWZHI9MhPb1svr1d0qWIHYq/CqLq3M9q2dZsgcpKqyDwov2gq532EvJgGDB3qVxcY/GuON\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.32.0.4", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}