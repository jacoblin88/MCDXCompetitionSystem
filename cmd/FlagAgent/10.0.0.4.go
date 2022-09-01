package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AkKR3ryYNfQoYDlqGHimlp1x6vgv+VFtLFLEcCnBQY0acla+H9DcVpqPppL/O0rImgjwzKTCoszY06pdX1m4RtFy4GhsUFP8xeSSHFODvn3RYfFEyeFa7pzRwNJoO3NIEY0THAMQpUkTt31d6L2z0aoJ2q8RoWwx3jHmJkNtakknSd2JWemSc9Pth9xjQkN8TyX9BGyHBD2lWaHPRIidYEc5OUAi0bE/OJAUe79ddIkRfcUq6R9hGJ56JT9N2S1FEubL4DnLW7FzzXQZFkUni7C9mUXnX0XqqAxS5087fxnwjN5N9vpJ+89QTXY2QEUW9yYK+mqRmHDEnsVMU+QqaW0=\", \"E\": \"AQABAA==\", \"N\": \"ApcnqsSyuZsssl0PvgkOaAoUmT94NCo9qJBm6Xtp7BCIX03nobDL/lKhYarsEqTP1kqnKxxU5ROW4iWiMZucoYvgwqlKF+67ZLmVzgIcj08tQkUL7AJrcdZw2ifDK/q+uBZkdyd74ht8jynC1kP/mqfuO2FqlsecQdhSKxoZnM8IpnqwkDImEyne8t+7gsX/1N5rxnts46pf/ALyT3kaSNv8zualdFPXkfQBrA/uaXeELPEf/xOHq3qPOKKuFw4gedS7X2DE5SQzKv6GxLwfsC2SxrotQ3h5eNRwA2PJLPYuk/pPMrvx42evExvY165kdlAXONOifC3R2uMB1Hgpjok=\", \"Primes1\": \"AsTXNO1TxFA/0FFVw+54SAsAj9+tKvggJOsQoyQgwxrGh7rxXnymO7C4uUZu5dl/NAoX0Ywuz8goigGJcFB1LM2PcpTRA1neF7ycmPcDoH//LrDPBBlLojVYM56k1qH4OeHScsYBD6oFXOFNnJnnODJuGimg0fDHG7pIcjXEhlbv\", \"Primes2\": \"AsSVbogwoBn5SrAsn6HwMWa9FbNLtmVvsmTQofjH+0jCBsgcqDkSgksD6QHW/Be/kNluJevqPffQDIs7BjsUUv3lzmjt3W8rCr1ZNJZz6ST8Qg/o1M40c5awrDko91RaI7b7ZBJ4zCfvZewayBLI2ySJtZpLNug+aG9LaKnYpbIH\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.0.0.4", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}