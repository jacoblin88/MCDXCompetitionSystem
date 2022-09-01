package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagAgent"
)

func main() {
	privateKey := "{\"D\": \"AoULd6RSJYxCmBaA5PX9sBc5yjOrhi+1D7IwLymmEUx3BHWy67xMsNR1QlSw3gKGc9qEh9y4vx6pRecLoVuv4Nb1yKVnxxHK7OhjNRDlUA6LML8Bc3pCV7OuBQJR5ify0ryZgkCLgvPauozlg5xJ+1Y+iBSa6bHCANHpnofYEdldaRQRgPM3PB5vgk6oAtd17oEkv9la62qiVcRQ5gqAIdNKmTv/l8JWuBJBMDsz3qBSwIQccWtUJ+7Xfqs8oFmlhDRtWlLCluKpwNMxCMk0LI3e6q2tYvls6RhGJUxW+cQA5D/h+WAJQELgjQyq1PgxPsIn/MFOwD0ByV8Ls1dYOJE=\", \"E\": \"AQABAA==\", \"N\": \"ArOQ2QnY2V5cHXHkVFjLk0DIQstBy1weqLj8Iy9L9vIH3uqaT0788AghRVEMVxyxYp++uTlfpQW/pE9tGhJqn+fTYBWDHkfGCvZU6J/RnGu8TZ24NTQ2oPSnMNTbvcr3599AWDOcsW7A4pyUgXLi4sMuVRt7rwvGsicMDeI8uMhBi1oM4ql4RC6PuCg3Z0ziqVROxR6i6laQpNDgY5jLAe5Nu5UBKO1WqWUYJYq9TdfQyj3b7EGEBkgOaysfUiNJ0m3BbsSywE3KGxIIw9hSjw3CYby0BAmFNCKd9rNMYM7WXKnovp9Aql7KjCwNLBrL7PybYAi/rOvAT7hrYIZPqis=\", \"Primes1\": \"As3EaooqEnZK3RGLd3HttuR6Kw8nETHn3M4E85L1skq89oa43l0X2ukRTRHuPR5CaeSfTiWZswgJmO91NzFm/cORdj/DWGwpDJ3vVjBLycVCLzuxOPduLws8WlX+dE1DIn4aG5un4X/1LkKllqpAAz4zwnEQJhzYm1JZMf0QBLuP\", \"Primes2\": \"At9m89CEampgRIjzbl72jGUrWqE3OvsNZyZWeI3P0tv5O7nOpUqK6b69lgbXIsmGvQb3PDAz/eTOblOOCF+2xc+7vGGK4po2g6AjDorDtN0ktPCFGIr75mPAFohubd17ImbZvvAHGahO7f5nGEmDG9FiBHnMXrf24CqcvkjCSUml\"}"
	publlicKey := "{\"N\": \"AqRzRA81LnCcUJdJCASXw4XhAmQczECCcY4zDiYKggmsTGAebcOXaq9HelkpYp1cN7fNhYHO+AoIjAoiBrkv9JwugE2ro9Htl6bmuWJN7LMVIEyrLxGM0XLoC1mEjNNqidg54tVwhkW3wWGFky+WOeCElKo+NBnHdAe1ekSSeQLLsKPbg1WRPZtky7aseexr+NZqKbAHPLQdsLhIddlOMRNnNllWM0q1BFUl1IeRQCAivPtrGrAy1XqVRKeT8xrLAQxfXGokmK/Goo8HZW9RZipqfJ9xbO3sgfEmabOxBnxF1PlA+YvURwQrvKA41af09/aJ9WlMXEtgGzfadEFn6N0=\", \"E\": \"AQABAA==\"}"
	f, err := FlagAgent.New("10.64.0.3", 8443, 3, privateKey, publlicKey)
	if err != nil {
		log.Panic(err)
		return
	}
	f.StartService()
}