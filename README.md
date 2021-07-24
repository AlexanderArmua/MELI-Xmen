# Test Mercadolibre

### Steps to init
1. Clone the repo
2. Import dependencies
3. Run main.go file ```go run main.go```
4. It will be running at port 8080

### Steps to test endpoints in localhost
```
Message: POST 
Url: http://localhost:8080/mutant/
Body: {
      	"dna":[
      		"AAAAGA",
      		"CAGTGC",
      		"TTATGT",
      		"AGTAGG",
      		"CTCCTA",
      		"TTACTG"
      		]
      }
Response: 200 o 403
```
```
Message: GET 
Url: http://localhost:8080/stats
Respuesta: {
               "count_human_dna": 7,
               "count_mutant_dna": 4,
               "ratio": 0.5714286
           }
```

### Notes
1. The database selected for this project was [bitcask](https://git.mills.io/prologic/bitcask), a DB built in Golang, works as KVS, has a very short response time, and is easy to implement. A relational or documental database has a lot of characteristics that we won't use here. In this case, we only need to save and read.
2. The `mutantes.conf` file and databases will be generated inside project's folder to facilitate the system execution. If it were a software running on production servers that kind of files it will have to be moved to another folder.


