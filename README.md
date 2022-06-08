# 21_school_Go_intensive
Solutions of tasks from Golang intensive in scool 21 by sber.
## gRPC project
### ex00
Create a gRPC server and client that connects to the server and receive infinite number of samples of normal distribution from it.
### ex01
By 50-100 of these samples (use sync.Pool) approximate mean and STD of the distribution. After that evaluate all other samples for anomalies
### ex02
Add all incoming anomalies to DB using Postgres
## Day04
### ex00
Create a server that imitates candy vending machine and receive POST requests with order marshaled to json.
### ex01
Implement TLS to this server
## Day03
### ex00
Load data about Moscow restaurants from csv to postgres.
### ex01
Server sends html responses with pagination for GET request with page property.
### ex02 
Server sends json responses with 10 serialized restaurants, current, previous, next pages and total number of restaurants for GET request with page property.
## Day02
### ex00
Recursively find all files/directories/symblinks. File extension might be specified.
### ex01
Count utf-8 characters/words/lines in files like unix wc does
### ex02
Run programs and specify its parameters from standard input like unix xargs does it. 
For example combine three programs from day3:
```
./myFind -f -ext 'log' /path/to/some/logs | ./myXargs ./myWc -l
```
### ex03
Make "Log rotation": compress log files to .tag.gz and add unix mark [MTIME] to its name.
## Day01
### ex00
Read json/xml files and parse them to struct and xml/json.
### ex01
Compare two files [.json, .xml] with cake recipes and print differences between them.
### ex02
Compare two huge file bases and print which files were added or removed.
## Day00
### ex00
Count Mean, Median, Mode and Standard Deviation by sample of ints.
