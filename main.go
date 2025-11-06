package main

type Options struct {
	path    string
	outFile string
}

func readOptions() *Options {
	opts := &Options{}
	return opts
}

func main() {
	_ = readOptions()

	// TODO: Read in the data from the CSV file. Handle any errors. Sort the
	// data. Output the data to the terminal or to a file. Handle any errors.
}
