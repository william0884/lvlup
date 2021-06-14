package lvlup

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/rocketlaunchr/dataframe-go/imports"

	"github.com/alyu/configparser"

	"io/ioutil"
)

var ctx = context.Background()

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readConfig() {
	configparser.Delimiter = "="

	config, err := configparser.Read("sandbox-config.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)

	section, err := config.Section("DEFAULT")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf(section.ValueOf("apikey"))
	}

}

func readCSV(csvfile string) {
	// give it a csv file and
	data, err := ioutil.ReadFile(csvfile)
	check(err)
	df, err := imports.LoadFromCSV(ctx, strings.NewReader(string(data)))
	if err != nil {
		panic(err)
	}
	fmt.Print(df.Table())

}
