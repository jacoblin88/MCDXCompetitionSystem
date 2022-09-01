package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AgadF+Cb5VojcfabPaNbYm1dB/TOjsxjJMDkJUKqSHlJZcCeSeLBOqDpLPfiIOt5D4IdtV1PTnMIlNCtuUof+NgZ80z6e7DnR/Fu8Flmt0cfNivx8aSP2x/MKPuj1D+mTzOUBu19kIO1oT4a5D0PJk7aYgmjKSKW53nh+3gvU2bShc7vf/6JyJC9/k5worfQcZbVHL06D3uyrm5ztTnYxIoK3WATa0eD9G7DAXkIF2W+v0eTH91MJu8T4ZLWOmm9VdCurI/ArgW+IEqRvaPO+w5SY57RlPLCxzUwFSVoliHe2aggopNGV+9CNYm9ezO0k7L3bGSXX6X142ZkvBAB/PU=\", \"E\": \"AQABAA==\", \"N\": \"AuHd6hUMImwpop/6KbOQLf2Gx4yKr68mcFRuNgC9US9XQCs6XUEdqqQPMRmgUuxjHaOUrMJS8cVwtIdXpDLcgr6a7y6HXBl56OIaHjHEl8pv9tDsmgOrHStCrBAd3gvQWvc51LaPisaFy/ACcnP6lFgjVlB35hYmbcMYl9WvGkzWVw0y6JVLzQL0S5HSiBRBdbPbhUlRNS99roAhglSX4FhnOH4qiyBchAH6mwZdH9xWN30sA1piN2mkwHYUoEgJOuRD/NAbCje+H4Z5q1BAQvVHf1mtnwEBJf38lWIIeMjXj0gP7fC1iSqrPAMF/JObSWFckorcHH7pK+P44cbhlcU=\", \"Primes1\": \"Au9ZGStBfA9jC0PxmMJ2ulMIM5+i5Nlo742Mak6eJu+N39FkI+9HVkqSyLMKORyoJ7w4z9LVLD9Jid+sVt5hueHZAud7C/KEgi0iOAFG+S+kv5YGKCTQQaN7GqKEvtGSPWsEQG/OkoLMkA3Bx1OW4DtmveN0jaUd07v0jJGD5cNb\", \"Primes2\": \"AvGUtbJiHYAwccLdmURr4WUzohWhP06Z1+IZ+9BjELSm20Hj5Ebg/0tZH/+5w6NxQsFk20OKaUxnxFifFCtKy70X/uIMdKRiPSSXeplrx4945IAtV1nceVgX/PwjF9Jf9ZXtU5nkpiHErHvExwFojLW6K4JlLcAsL0TxI656ePVf\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.64.0.2", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}