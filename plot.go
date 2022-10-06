package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"os"
	"strings"
	"time"
)

func main() {
	filename := "data/arnie120.tsv"
	columns := []string{"Latitude", "Longitude", "Altitude"}

	df := getTable(filename)
	df = renameColumns(df)
	_, _ = parseTimeStamp(df)

	output := df.Select(columns)
	fmt.Println(df.Names())
	fmt.Println(output.Names())
}

// parseTimeStamp constructs a timestamp from date and time columns in the Dataframe.
func parseTimeStamp(df dataframe.DataFrame) (series.Series, error) {
	day := df.Col("Day")
	date := df.Col("Date")
	t := df.Col("Time")

	if (day.Error() != nil) && (date.Error() != nil) {
		return series.Series{}, errors.New("dataframe does not have column with name 'Day' or 'Date'")
	}
	df.Select([]string{"Date", "Time"})
	return day.(t).Map(func(e series.Element) series.Element {
		d, _ := time.Parse("01/02/06 15:04:05", e.String())
		fmt.Println(e.String())
		e.Set(d.Unix())
		return e
		//return series.Element(&series.Elementd.Unix()))
	}), nil
}

func renameColumns(df dataframe.DataFrame) dataframe.DataFrame {
	for _, c := range df.Names() {
		newC := strings.ReplaceAll(c, "AEST", "")
		newC = strings.TrimSpace(newC)
		newC = strings.ReplaceAll(newC, " ", "")
		df = df.Rename(newC, c)
	}
	return df
}

func getTable(filename string) dataframe.DataFrame {
	f, err := os.Open(filename)
	if err != nil {
		return dataframe.DataFrame{}
	}
	return dataframe.ReadCSV(bufio.NewReader(f), dataframe.WithDelimiter('\t'))
}
