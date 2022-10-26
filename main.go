package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var exampleDsns []string = []string{
		"sqlserver://{username}:{password}@{host}/{instance}?database={db}",
		"sqlserver://{username}:{password}@{host}:{port}?database={db}",
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Printf("\tdbtest \"DSN\"\n")
		fmt.Printf("\tdbtest \"%s\"\n", exampleDsns[0])
		fmt.Println("")
		fmt.Println("Avaible DSN connection strings:")

		for i, eDsn := range exampleDsns {
			fmt.Printf("[%d] %s\n", i+1, eDsn)
		}
		os.Exit(0)
	}

	fmt.Print("Connecting to database...")

	db, err := sql.Open("sqlserver", os.Args[1])
	printOnError(err)
	fmt.Println("OK")

	defer db.Close()

	fmt.Print("Pinging...")
	err = db.Ping()
	printOnError(err)

	fmt.Println("OK")

	fmt.Println("=======")
	fmt.Println("Connection established")
}

func printOnError(err error) {
	if err != nil {
		fmt.Println("FAIL")
		fmt.Println("=======")
		fmt.Println("ERROR:", err.Error())
		fmt.Println("=======")
		os.Exit(1)
	}
}
