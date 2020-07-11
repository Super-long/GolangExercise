package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Track struct{
	Title string
	Artist string
	Album string
	Year int
}

var tracks = []*Track{
	{"Go", "delilah","from the Roots up", 2012},
	{"Go", "moby","moby", 1992},
	{"Go", "moby","copy", 2020},
	{"Go Ahead", "Alicia keys","As I am", 2007},
	{"Ready 2 Go", "Martin Solceig","smash", 2011},
}

func printTracks(track []*Track){
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0,8,2,' ',0)
	fmt.Fprintf(tw, format, "Title", "Artist","Album", "Year")
	fmt.Fprintf(tw, format, "-----", "-----","-----", "-----")
	for _,t := range tracks{
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool{
	if x[i].Artist != x[j].Artist{
		return x[i].Artist < x[j].Artist
	}
	if x[i].Year != x[j].Year{
		return x[i].Year < x[j].Year
	}
	return false
}

func (x byArtist) Swap(i, j int){
	x[i], x[j] = x[j], x[i]
}

func main(){
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
}