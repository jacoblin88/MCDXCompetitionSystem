package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AlZPt+Cua03FiedNOeG6AzodU+Xajts/P716Ou9phisBlc2bKLNldpHA9mVE9oH9R22/LrIRWxxARPzaYOhqp2BX0/ioA0gOH3rbDs4h1dMEJIpEVA2kr8Y5jnfmpnjU0YcX0dJsYQ4YOVjwohlxprUUYuhlyIoYszFEsVVwNU1QDGvjgzXDC9xGNkrMl0f0gP63qJUbH3zqtsyBYvfDj68qYsV/CrN3ZaUvmrgyZ7aLag4s5hFZwP1h0PkhfkFUp8D9kPexgrkB300dcw9YjHqm6+/zZ5ODZFkpEP+yMpGZQzmIsa1KecrDf78WaCbpCMX3gyYoL9LBrwRrWeGqt0E=\", \"E\": \"AQABAA==\", \"N\": \"AqH/E98Gd1uL/AwTbCGizzemXZoj12WdZilzrchIxN5Jwqx47WlW822aGAFBnW2akoE6lOWBJR+nF9XiKMEMgYZhPy+jxQ43+smEOk76rVJvSdDwOZxEqYZWORutWQtMxxewxKzzXxPvsNQoX+knQPZNGPoJiatYaYdhxtwLC+6gOi7BE7kShFfPpNefhEbkva5qWQkRvvAc9kgWkm2lJm9XBODrIQosMyrjGTraQPuVFqm0abxBEY2EJQxiPtRRgc5wRQnlzrrk2yqJwh6TGI3xAIuTAk/zB6D63zsk2fZoatG40uuGhZm41Nh7A7/NselwjDZmsF8XwNRTYVi/E40=\", \"Primes1\": \"Asm4ce2oUw4IwR5EHr8c1pNlpwpa/Uo/bG9RHt7FsO6lrRHOfTfu9H64JxbBHEQJ5bRMX71nVJ49isqX8NHe8bEWDd1AIY7VLbOGhds55TbIDm7M1LNpWW/HSgM58BJ713m/xsGSumpHDz6MMfZ1ICUpgs2ZjjEBCGBsd9TQoBGV\", \"Primes2\": \"As2WOxiu+UBiMPq6fDldK/9MWb1Kn5pFXsEQoFgMYGAtxGMgfhFSr8z9WTnq4z9JbzcCYmGXzXNVWXOYRp1IDUUO+VTKWYAd/MWromvo1Op5VZS63M+Ap/5AjjlyKRsHlKy4ahMuCgFyy0sd64mQpAVu03IeXkC8MxyeRN7bSOwZ\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.0.0.5", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}