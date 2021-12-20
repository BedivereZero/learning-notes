package main

import (
	"flag"

	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)

	flag.Parse()
	defer klog.Flush()

	klog.Info("hello, world.")

	for i := 0; i < 10; i++ {
		klog.V(klog.Level(i)).InfoS("hello", "depth", i, "type", "info")
	}
}
