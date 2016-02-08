package cluster

import (
	"fmt"
	"strings"

	"github.com/golang/glog"
)

var times int

func CreateCluster(args []string) error {
	glog.V(2).Infof("Create args %s", args)
	for i := 0; i < times; i++ {
		fmt.Println(strings.Join(args, " "))
	}
	return nil
}
