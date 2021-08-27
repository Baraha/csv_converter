package api

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

func LoadFile(ctx *fasthttp.RequestCtx) {

	file, err := ctx.FormFile("file")

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("file", file)

	s := strings.Split(file.Filename, ".")
	fmt.Println("file.Filename", file.Filename, "and file name", s[0])

	readFile, err := file.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(readFile)
	var headers map[int]string
	headers = make(map[int]string)
	var data = map[int]map[int]string{}
	cnt := 0
	separator := ","
	firstIter := true
	for scanner.Scan() {
		s := scanner.Text()

		s = strings.ReplaceAll(s, " ", "")
		switch {
		case firstIter == true && strings.Contains(s, ";"):
			separator = ";"
		case firstIter == true && strings.Contains(s, ":"):
			separator = ":"
		case firstIter == true && strings.Contains(s, "-"):
			separator = "-"
		}
		fmt.Println("sycle text", s)
		fmt.Println("separator", separator)

		if strings.Contains(s, `"`) && firstIter == true {

			for index, value := range strings.Split(s, separator) {

				headers[index] = value
				fmt.Println("header data headers[index]", headers[index])
			}
		}

		if firstIter == false {
			data[cnt] = make(map[int]string)
			for index, value := range strings.Split(s, separator) {

				data[cnt][index] = value
				fmt.Println("value", value)
				fmt.Println("lines data data[index]", data[cnt])
			}
			cnt++

		}

		firstIter = false

	}

	fmt.Println("END PROCESS : headers", headers, "lines", data)

	// lines, err := csv.NewReader(readFile).ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// var (
	// 	ddl = fmt.Sprintf(`
	// 		CREATE TABLE %s (
	// 			salary   INT

	// 		) Engine=Memory
	// 	`, s[0])

	// 	dds = fmt.Sprintf(`INSERT INTO %s (salary) VAlUES (?)`, s[0])
	// )
	// connect, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err2 := connect.Exec(ddl)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// tx, _ := connect.Begin()
	// _, err3 := tx.Exec(dds, 100)
	// if err3 != nil {
	// 	log.Fatal(err3)
	// }
	// if err := tx.Commit(); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("end process")
}
