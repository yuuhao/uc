package utils

import "fmt"

func checkError(err error) {
	if err != nil {
		fmt.Printf("zooKeeper err: %+v", err)
	}
}

func InitZookeeper() {

}
