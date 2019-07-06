package core

import (
	"log"
	"os"
	"os/user"
)

// Built-in Data set
type _BuiltInDataSet struct {
	url    string
	path   string
	sep    string
	header bool
	loader func(string, string, bool) *DataSet
}

var builtInDataSets = map[string]_BuiltInDataSet{
	// MovieLens: https://grouplens.org/datasets/movielens/
	"ml-100k": {
		url:    "https://cdn.sine-x.com/datasets/movielens/ml-100k.zip",
		path:   "ml-100k/u.data",
		sep:    "\t",
		header: false,
		loader: LoadDataFromCSV,
	},
	"ml-1m": {
		url:    "https://cdn.sine-x.com/datasets/movielens/ml-1m.zip",
		path:   "ml-1m/ratings.dat",
		sep:    "::",
		header: false,
		loader: LoadDataFromCSV,
	},
	"ml-10m": {
		url:    "https://cdn.sine-x.com/datasets/movielens/ml-10m.zip",
		path:   "ml-10M100K/ratings.dat",
		sep:    "::",
		header: false,
		loader: LoadDataFromCSV,
	},
	"ml-20m": {
		url:    "https://cdn.sine-x.com/datasets/movielens/ml-20m.zip",
		path:   "ml-20m/ratings.csv",
		sep:    ",",
		header: false,
		loader: LoadDataFromCSV,
	},
	// Netflix:
	"netflix": {
		url:    "https://cdn.sine-x.com/datasets/netflix/netflix-prize-Data.zip",
		path:   "netflix/training_set.txt",
		loader: LoadDataFromNetflix,
	},
	"filmtrust": {
		url:    "https://cdn.sine-x.com/datasets/filmtrust/filmtrust.zip",
		path:   "filmtrust/ratings.txt",
		sep:    " ",
		header: false,
		loader: LoadDataFromCSV,
	},
	"epinions": {
		url:    "https://cdn.sine-x.com/datasets/epinions/epinions.zip",
		path:   "epinions/ratings_data.txt",
		sep:    " ",
		header: true,
		loader: LoadDataFromCSV,
	},
}

// The Data directories
var (
	GorseDir   string
	DataSetDir string
	TempDir    string
)

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error while init() file built_in.go in package core "+
			"see https://github.com/zhenghaoz/gorse/issues/3", err)
	}

	GorseDir = usr.HomeDir + "/.gorse"
	DataSetDir = GorseDir + "/dataset"
	TempDir = GorseDir + "/temp"

	// create all folders
	if err = os.MkdirAll(DataSetDir, os.ModePerm); err != nil {
		panic(err)
	}
	if err = os.MkdirAll(TempDir, os.ModePerm); err != nil {
		panic(err)
	}
}
