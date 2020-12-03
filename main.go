package main

import (
	"fmt"
	"k8s/clientgo/incluster"
	"k8s/clientgo/outcluster"
)

func main() {
	/*
		args := len(os.Args)
		var inOrOut string
		if args <= 1 {
			fmt.Println("Usage: need argument in or out")
			return
		} else if len(os.Args) > 2 {
			fmt.Println("Two arguments is enough")
			return
		} else {
			inOrOut = os.Args[1]
		}

		if inOrOut == "in" {
			incluster.GetK8sInfoFromIncluster()
		} else if inOrOut == "out" {
			outcluster.GetK8sInfoFromOutcluster()
		} else {
			fmt.Println("Usage: need argument in or out")
			return
		}
	*/
	var err error

	err = incluster.GetK8sInfoFromIncluster()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = outcluster.GetK8sInfoFromOutcluster()
	if err != nil {
		fmt.Println(err.Error())
	}
}
