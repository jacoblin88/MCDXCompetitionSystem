package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AtJixZwlHy6dIMtqT0vh2xyzzNiXzXjhVlZxCL2miiIEtSUyPMlf15wtOsRXWRt6SxOuS9Wsqf/G768Mw456sFFBWSZZZTCJfvR+I5mZNwZKwyuqgdCzS17hasBlDfpPOwPBHdXzGDi3nTSsMczxy55JZ4A4f4g6rwvwpUfroBaPoITMvSY8kUy+FyYAQU9Bwr+/VkmC5RWnq76imDQt41nE8dVsYvWGKOI7/XoqAR0NGQ1n3MctbaKlIKnjG2oCbRX3XYf8jah/KcApbujjZSC8tGPb4NtiJmb96WUqmdW5QwUq++kqwxhoTqS5pFidLafEkXWZGNrfUW26PrfU1CE=\", \"E\": \"AQABAA==\", \"N\": \"AumvGXNo8FBxJwDk7MRgnHeE/n8EB+8pElkArXmx4BPrJygrgBdzC+stdhhMLXcTIctk5VRaH5RVVRrR4JNNUxUtEYgYqNabWf1/N54RqVXeFf2e9zXnvyLh+n9k1rs0msSXA2A3UekiTmcWHeG0e2Cdm+T6ajmIbeLeQ0MPO8I0dB8fBsgkfQc3FZmdukyJO8ljd1s+/zXCAGkNXy12YG3XIPML5RZKGQkFc7sTGklSKKiTRg6qWD20nbW9HDH2TB/eJ8bEvCYJ5Avt23JhLRd2QEyAWj7ZfVirA0nmzuFK1fJV/g/JOsRrHw2N5S+SeFQNCLnDYFU/cIxs/W4iAas=\", \"Primes1\": \"AvFaV4blTsT+zTYYe66sw8nCzDhLnzwZoSvAk5YYTBwHS9CrbjkvKNoUyMCV3pvyobPuQOarCYiFf/PFuOFT9H8BC/tbueH/EaLvspRdkV64B2DCdewhOGeXjAVsFlQB/23yUie1UzrecTFyeZj/AiemJW2gMTQP+ih8Tk5zFhtb\", \"Primes2\": \"AvfdnP4wA5PH+hb0EuAw3Q+vfEWoF3u+V7XAdGAwkpp7rDUc+CDyGQWuQN/7P0P0vIjdvUz1O1dxxFYVDZ4jOVh+JV8X5fhJ/1I0HrcccB53PHyR0aieVHWCn8lBGSm0oqsmM06wM10eyUpXpAKYMVsXsztO1wOct7cNep2ZZ5Px\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.96.0.6", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}