package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkjZtO5H8vW/5dYyoATEve8k8ODWJ0ApERu6czXcurVyGXql2pZRF8pnIOjealT19NeL8i5PLPX5Y8haDHq4HOZrBACdAYMti8/BFqMF6sW9RHpkH3tj0HqM/9TKuQx5KJzBguvIv4GV7Wvzuxxfd/FCID8gHP8lJc0ajgUO2VYP3SuAbnrTJEY2kupP4e/tIszP+RS74WsX4UpGVssi0aXcziT7kcThiCZoSH93u4qatE8IX/HG3uNsYjqqGtC4TpT6shqkMtncqv8ERLKcdn+3D8F6FasIbdGZpJ3nrQ99fNnFmYSkf+ACIQQDIucORr9+Fgo2j3o5D5HpLmSXu2k=\", \"E\": \"AQABAA==\", \"N\": \"ArnCgLobdoZVVfp3d1akc6h1Ii58L6IjHN93GElOrBiXET0hOFE7rusDR7G+u2g1KoumxdWVI3NDpeV75Qj9d16kXoojwj91LtiqbMSK3VrHWQ/NlGJICByAUATi1akg0fFR2cPqhLwGccP4gCFppGmFLzRVx4mc9syWoY1s2alxjYP9pb9P4NZzqLxELsJiReWIsrseUqQuTJBe4Gmlx4b/Jc44yJmEyCFwNeHEiAODWQb5ihN4ggE+EGirj4HaYeIqNH0utNUj6vMYqet1ydjQ92m+B+QA6JDBHBC7NCZaUapCGgDsga3gmhdjHzDFQ6V2ek5zkiacyUwq6xZYWL0=\", \"Primes1\": \"AuYRR3EhHcZ1yX9cEbXnqNsVB6S0OktwUGSresrtSzoEZVQko0XzV6tqfe1NDXG65YB0m2c20lFFnbQFvUp9qcl+3XtbPyQMYrOWtbqfarvGzRa9g5gAwmHqEWY5UXqo8SlpSECxWsl3+td/FNoiYzMBIoKEaODIWd0NBoiq+6R/\", \"Primes2\": \"As6ys2WcNRmghZN78EsUJQ+jZWfxDFN4D+cdm9I4EdWtAeJgwc5sNtK64uprKJCXzHN/QcyfvmMmK0+MEEJmB13UDsJcx0BS/SEDWTx7EHFA33K0nc9GxEjWTVH2w+AvHgNjWr/4KEZHhvezsfGG7lP+8EotjysJHYOS8NOrn/TD\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.32.0.5", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}