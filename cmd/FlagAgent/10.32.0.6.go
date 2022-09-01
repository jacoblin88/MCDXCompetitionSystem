package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AndkKM4L1Dy352mV8boWrHoFb8m2cJVVs7SEZxy1Tq3OyuMuQvaEWRXwUDUxIms3mG7qvwlkPFr08NdSN1DbCqE1mTU23iCuAyirHWxF9PBQtljMacUaRRMTB76KLpowawrUG+zpmiNTjxFg3xmAnao/71Cn+OHP+CZldzxAQA3Xl+BAqFPyfspThwMtgYqMTGBP3djXrqSia8QgKjOVJ0/aayB9zQk0d3xN/vmerdHlbwfm35pwEZNl5RDZQHfroA8PNlph6cWatj//MGSy4jygXj3M1p3sUz65NRlGmcnml4XhZwZxxBbK/JoFq4gNmjSRjxEE88yWcvCW4UrdCgE=\", \"E\": \"AQABAA==\", \"N\": \"Aqsro3Nvv+QchHGaip5J3vCZR1IF4ceVmbqaGGDul0USJhfL4nUQfCjpiy/ERsRhJtXQaJdzLIdmTgZKGGO+NHy7EBmJprTjy8xJ3kKvRe029wDSZ9wuSvvsb+y+7l9VbCBAu3VSvpJPq/qIpkY/kdjrQ+xq3Kk5C1Nojh9QWsI43jyJ9T5KJ9NKTOmYoWPmbaBFmjYDP95G1aO7/viH09gds/79q0SxlQYGSPunNbIOpru1wd5SZp2SG2b+1zpfZ3X7ISwrur2gQICDyJht7GbtYawLz0wBUNxo9b8NcSTVuTCPFvcs5d4ljkDOLxOOIzf92ljh6QhgajNPndSmcpc=\", \"Primes1\": \"AsKl+BZKkNl2SZjLvfnMbuXkxV6XVMVsjYLhHhAPOfnJyxcNrBfP1q8121APvSXLT9y6RaB9A+3N2ctNSucJ7pn6ZfsbeCWJXJ5WdVP7l5V3JyD0iEcfHXMn9vuJzmoYKXkcxf/B2o/eWv68iaiD7/+evsc1l5xi+fd1dIaUvp/x\", \"Primes2\": \"AuEfQioauxU2o5jW3LBq2yGxILmOEuIMMjKOOJmWPo7ZiTXWHlYJxPDUhxy4fpWteEPh9CpVTSZyCpysqeUpS02lnHdIY6HiyeJZyQKGPRO8ap28jyfhP1h9oIzgajMh0T+9Foa4z9AVCF5ipz0ZHAVM1mV4TWiYFNI4eCgs1EMH\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.32.0.6", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}