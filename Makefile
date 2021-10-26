#================================
#== GOLANG ENVIRONMENT
#================================
GO := @go
GIN := @gin

goinstall:
	${GO} get -u .

gorun:
	${GIN} -a 4000 -p 3000 -b bin/main run main.go

gobuild:
	${GO} build -o main .

gotest:
	${GO} test -v

goformat:
	${GO} fmt ./...