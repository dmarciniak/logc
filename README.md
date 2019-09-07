# logc (log-console)

[![Go Report Card](https://goreportcard.com/badge/github.com/dmarciniak/logc)](https://goreportcard.com/report/github.com/dmarciniak/logc)

Console application to read logs from files, sorting them and write result on the console output.

Usage:
```
./logc file1.log file2.log [...]
```

You can use grep to filtering result:
```
./logc file1.log file2.log | grep "TOKEN"
```
