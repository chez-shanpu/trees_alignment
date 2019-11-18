package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chez-shanpu/repo2tree/model"
	"io/ioutil"
	"log"
)

func main() {
	var treeFilePaths []string
	var trees [2]*model.NodeInfo

	flag.Parse()
	treeFilePaths = flag.Args()
	if len(treeFilePaths) != 2 {
		log.Fatal("Number of argument is wrong.")
	}
	for key := range treeFilePaths {
		trees[key] = readTreeJson(treeFilePaths[key])
	}
	dist := layerAlignmentDistanceTotal(trees[0].RootNode, trees[1].RootNode)
	fmt.Printf("Alignment distance between %s and %s is %v",
		trees[0].RepositoryName, trees[1].RepositoryName, dist)
}

func readTreeJson(filePath string) (tree *model.NodeInfo) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, tree)
	if err != nil {
		log.Fatal(err)
	}
	return
}
