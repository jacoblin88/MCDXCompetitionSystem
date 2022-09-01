package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"Ao5H6HO3kZl6Njn1O9JQcBbgf/0X7PLREqYCnzZSFqlrgU72KkEN3a/oC5M62fSUd+8MDKBpoB1r28e3ns9DmrtfwzaTZijGsgPOfEwnJJxN04jGUzI+5I+YdKnar9GedE13EHXHTplaYBJEA80O3/pVMimuAALnTluZd7C10u+Gw0Tbp59qsdu+t6Vh5ThLkFHlnNMSv+aCZKn49RcRS3Z4lRcuIlSKhwj2OJzrfQaxwcUk3NuFyPf4tvkwJJxGdo9/DPXAh/V0IbYkrsbFAvLAGvSRnxaPwOVZZnZriILVkAAf6NuYoAr2Qx5KGbeWCCiqu8+4NlJJ7hc+SG9vXPk=\", \"E\": \"AQABAA==\", \"N\": \"AspuzO3X6izwEJAHJVbLaR5JYy8wbHSLCYa07dk569hrbrzmbNMffOrDTXDzGKUTCaPG6RPTTCw9B3j9xjphH+w6kb6Z+61aHvhC00JfEpkl10FF9uwLCrksiNdR/VKrwHcKV6A64/cuEFTxY7FXZiwUDfoIM4j/KvWzZrBZ64tflXfxk1LwKJk2eDnzZx7vCXTf2XadxlurKI6EW2H6OhIQQMt1lEmGYXt2Cr4JoJ3g+1ZQ7g18uCrx83spavhrLufalcLPr6QXm5vAJvbwNNHIlicZ5w9KwNeCzP/tf4Cqs2C9wpMTYQQAriCfy+baE4UD06iUbXTaY7Pazsgfjkc=\", \"Primes1\": \"AtLUIVc8P/Je3wdRnwV4MZx8ILEZrvotRMI2u8bIDx77WOSfwELXuzyME5XSP74/dl+4n+k8qjB5j1+OgEswhHZZbsNpH7I+5bXqMT5wwJjSUeewtt337jt9BzwsM6+HIa3ofMvfBtIFgcaa4raVJ7OQLS7mf3tUQexySYldPov1\", \"Primes2\": \"AvXOKa0YueiHthts6hbBh1/hhZfx35ptmXZL+Fveev1BvjQyOeGBSIr7K70odHKN6Ephibn91HHXfWHpf2pZuhPcyogoS1ZvDEyTD7vqkv32kMmGNMah5CQAE1E0X2P6MM0Ebwq5fL2od8SZ1WzxLHXOFfsgWYx1kelMF2wk+2fL\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	// f, err := FlagAgent.New("0.0.0.0", 8443, 3, privateKey, publlicKey)
	f, err := FlagAgent.New("10.0.0.1", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}
