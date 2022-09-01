package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"ApidxFhNcyUrPL2npfNrWA9sTgrb06Evlb2IRA5J+FJ6OPQs2Xye5Ml5V6CKJydXINcJL/yHpyW+RpuvxfRiXN49UA1grHIPEnafwqGurlhwgflQeIqS7xcKqDzVP7UkVeK4S1/+x/+q0uEDHCbxwx4jLbZsKI3Q0TT9CYKpAzrqBudx/ST73WBVyGhsBTts48im9HA9rhh/4zVwZs1E0XIIPQBtdtG3AiEHWcg0Rl2BRepsblIhmjQEBfhHSDNX6ppBuKzPJZW9rCD8imxlzc5Cd3K8WWPsaFsbo9/MkXYcsrokfstoqw0+bOEFIwc5i+K230MyY5AvZ3meoYj3pTE=\", \"E\": \"AQABAA==\", \"N\": \"Ar61AZ2xesmcj7lXw/UrlLn97YQY2Ob9faF7aEStbdoMDbfHLCCqkpnue27iHwWcGppXUYSnAsFAGv4N5lD6s0lR9dTEP9a+EEMsz1ZgtvCkrmHH3c7YNl2jXNTc8QqRXV3FpE4Cwz81qSoDALBBuryHncK7SJPElQ8CuKG2yF6b5AMUaFt2tj5ahVT/0SZoiAm6jjkCrC53ZZ2f/4QGKMZCI9tz5mXLz1SqmMOFfJkVvwK20odzAc71MY3thitPMsokRy/zuKBrZPAjIEtr3WdNKFYym7DSbGKXs93MvlO0OpPyp2TLYI5zQ6yB7NmRbEetTOrHo39ZYrojeWIouss=\", \"Primes1\": \"At+JqqK5TPKlT+5KY+AD/FMIRx9H7hv26da7S+1Db6SFntDSXSTUK8PHeKSxvKrCcywrhE0EIXI/fL93KFNv+LW0jtU3jDLBtirkhdYhF0ZgcBTL1pF8nrfQRbYeyN19UVM3fgmkrIH/b7ahdZXVkdQDh7FSMuVWGGy7OGjXExQp\", \"Primes2\": \"Atpmz88UG6Bna00Erm7sb9ormIPIe8Xfz2BmTRSSizVdVetfOAUN53dmUI8WEJ4+/sBfh7kCyhQuqYP/DT27FB2M23DV15ErmBKMjP6Z32SwAfRg4DrMiut4zAOssHQs8UMpPyOU2OLzu8EFH1jEJA0jb1FTGFqD5DJqBOoNgdXT\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.32.0.2", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}