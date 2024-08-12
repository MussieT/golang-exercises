package gomistakes

import (
	"database/sql"
	"fmt"
)

// Code must be organized for Reading Not Writing. I.e, It should be writtn to be read easily .. not write easily.

// Init pitfalls - https://www.reddit.com/r/golang/comments/prtpqy/best_practices_regarding_init_function_and_small/
// 1. It makes binaries bigger
// 2. Imagine someone is running your code .. and from somewhere .. an error might be thrown from some package. (imagine having lots of code?)
// 3. for configurations .. you can use LoadConfiguration explicit function call.

// Each source file can define its own niladic init function to set up whatever state is required.
// Besides initializations that cannot be expressed as declarations,
// a common use of init functions is to verify or repair correctness of the program state before real execution begins.

func init() {
    // if user == "" {
    //     log.Fatal("$USER not set")
    // }
    // if home == "" {
    //     home = "/home/" + user
    // }
    // if gopath == "" {
    //     gopath = home + "/go"
    // }
    // // gopath may be overridden by --gopath flag on command line.
    // flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func LearnInit() {
	fmt.Println("Should come after the init log")
}

func init() {
	// fmt.Println("------ init -------")
	// init() - it CAN Not be called like this
}


// Init function mistakes exmaple
// opening a database pool.

/* 
1. error management is poor .. init does not return error .. 
2. testing is hard - it will be executed before running tests .. (we might have utility function which does not need db connection)
3. Global variable is assigned here.
*/

var db *sql.DB

func init() {
    // dataSourceName :=
    // os.Getenv("MYSQL_DATA_SOURCE_NAME")
    // d, err := sql.Open("mysql", dataSourceName)

    // if err != nil {
    // // log.Panic(err)
    // // fmt.Errorf()
    // }

    // err = d.Ping()
    // if err != nil {
    // // log.Panic(err)
    // // Assigns the DB connection
    // }
    // // to the global db variable
    // db = d
}
