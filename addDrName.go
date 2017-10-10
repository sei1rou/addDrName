package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	flag.Parse()

	//ログファイル準備
	logfile, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	failOnError(err)
	defer logfile.Close()

	log.SetOutput(logfile)

	//入力ファイル準備
	infile, err := os.Open(flag.Arg(0))
	failOnError(err)
	defer infile.Close()

	//書き込みファイル準備
	outfile, err := os.Create("./test.csv")
	failOnError(err)
	defer outfile.Close()

	reader := csv.NewReader(transform.NewReader(infile, japanese.ShiftJIS.NewDecoder()))
	reader.Comma = '\t'
	writer := csv.NewWriter(transform.NewWriter(outfile, japanese.ShiftJIS.NewEncoder()))
	writer.Comma = '\t'
	writer.UseCRLF = true

	log.Print("Start\r\n")

	for {
		record, err := reader.Read() // １行読み出す
		if err == io.EOF {
			break
		} else {
			failOnError(err)
		}

		drNamePos := len(record)
		if record[drNamePos-1] != "医師名" {
			record[drNamePos-1] = "寺門　節雄"
		}

		out_record := record

		//１行書き出す
		writer.Write(out_record)
	}
	writer.Flush()
	log.Print("Finesh !\r\n")

}
