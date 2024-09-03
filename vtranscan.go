package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Linevals struct {
	btime      int
	etime      int
	findingval string
	message    string
}

func (l Linevals) String() string {
	bseconds := l.btime / 1000
	bminutes := bseconds / 60
	bhours := bminutes / 60
	bseconds = bseconds % 60
	bminutes = bminutes % 60
	eseconds := l.etime / 1000
	eminutes := eseconds / 60
	ehours := eminutes / 60
	eseconds = eseconds % 60
	eminutes = eminutes % 60
	return fmt.Sprintf("Start: %v:%v:%v - End: %v:%v:%v : %s in (%s)\n",
		bhours, bminutes, bseconds, ehours, eminutes, eseconds,
		l.findingval, l.message)
}

//const filelocation = "eepy henya playing games and chatting dayo.tsv"

func main() {
	filelocation := flag.String("i", "", "whisper .tsv file")
	storeagelocation := flag.String("o", "bakatimestamps", "file name of timestamped location of words searched")

	flag.Parse()
	fmt.Println("Opening File - " + *filelocation + "\n")
	writefile, err := os.Create(*storeagelocation)
	defer writefile.Close()
	if err != nil {
		panic(err)
	}
	bakas := []string{"ばか", "バカ", "馬鹿", "fool", "idiot", "moron", "stupid", "loser"}
	file, err := os.Open(*filelocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tsvs := csv.NewReader(file)
	tsvs.Comma = '\t'
	for {
		data, err := tsvs.Read()

		if err != nil {
			fmt.Println(err)
			break
		}
		for _, b := range bakas {
			if strings.Contains(data[2], b) {
				mseconds, err := strconv.Atoi(data[0])
				if err != nil {
					panic(err)
				}
				bseconds := mseconds / 1000
				bminutes := bseconds / 60
				bhours := bminutes / 60
				bseconds = bseconds % 60
				bminutes = bminutes % 60

				mseconds, err = strconv.Atoi(data[1])
				if err != nil {
					panic(err)
				}

				//eseconds := mseconds / 1000
				//eminutes := eseconds / 60
				//ehours := eminutes / 60
				//eseconds = eseconds % 60
				//eminutes = eminutes % 60
				_, err = fmt.Fprintf(writefile, "T: %0.2v:%0.2v:%0.2v %8s -> (%s) \n",
					bhours, bminutes, bseconds,
					b, data[2])
				if err != nil {
					panic(err)
				}
			}

		}
	}

}
