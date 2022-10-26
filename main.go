package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var (
	exampleDsns map[string][]string = map[string][]string{
		"sqlserver": {
			"sqlserver://{username}:{password}@{host}/{instance}?database={db}",
			"sqlserver://{username}:{password}@{host}:{port}?database={db}",
		},
		"mysql": {
			"{username}:{password}@{protocol}({host})/{dbname}",
		},
	}
	avaibleDrivers []string = []string{
		"sqlserver",
		"mysql",
	}
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("tool for testing database connection")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Printf("\tdbtest driver \"DSN\"\n")
		fmt.Printf("\tdbtest %s \"%s\"\n", "sqlserver", exampleDsns["sqlserver"][0])
		fmt.Println("")
		fmt.Println("Avaible drivers with avaible DSN formats:")

		for eDriver, eDsns := range exampleDsns {
			fmt.Printf("- %s\n", eDriver)
			for i, eDsn := range eDsns {
				fmt.Printf("\t[%d] %s\n", i+1, eDsn)
			}
		}
		os.Exit(0)
	}

	driver := os.Args[1]
	dsn := os.Args[2]

	fmt.Print("Checking driver...")
	if !listHasItem(avaibleDrivers, driver) {
		printOnError(fmt.Errorf("unknown driver: %s\navaible drivers: %s", driver, strings.Join(avaibleDrivers, ",")))
	}
	fmt.Println("OK")

	fmt.Print("Connecting to database...")

	db, err := sql.Open(driver, dsn)
	printOnError(err)
	fmt.Println("OK")

	defer func() {
		fmt.Println("Closing database connection")
		db.Close()
	}()

	fmt.Print("Pinging...")
	err = db.Ping()
	printOnError(err)

	fmt.Println("OK")

	fmt.Println("=======")
	fmt.Println("Connection successful")
}

func printOnError(err error) {
	if err != nil {
		fmt.Println("FAIL")
		fmt.Println("=======")
		fmt.Println("ERROR:", err.Error())
		fmt.Println("=======")
		fmt.Println("Connection failed")
		os.Exit(1)
	}
}

func listHasItem(list []string, item string) bool {
	for _, l := range list {
		if l == item {
			return true
		}
	}
	return false
}
