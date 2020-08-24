package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"net"

	"github.com/kinvolk/k8s-egress-filtering-benchmark/pkg/filters/bpf"
	"github.com/kinvolk/k8s-egress-filtering-benchmark/pkg/filters/iptables"
	"github.com/kinvolk/k8s-egress-filtering-benchmark/pkg/ipnetsgenerator"
)

var (
	iface       string
	countParam  int
	ipnetsParam string
	seed        int64
	filterType  string
)

type filter interface {
	SetUp(nets []net.IPNet, iface string) error
	CleanUp()
}

func init() {
	flag.StringVar(&iface, "iface", "", "Iface to attach the filter to")
	flag.IntVar(&countParam, "count", 0, "Number of entries to generate")
	flag.StringVar(&ipnetsParam, "ipnets", "", "List of ipnets and their weigth to generate (ex. 24:0.7,16:0.1)")
	flag.Int64Var(&seed, "seed", 0, "Seed to use for the random generator")
	flag.StringVar(&filterType, "filter", "", "Type of filter to use (bpf, iptables)")
}

func main() {
	flag.Parse()

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	ipnetsReq := ipnetsgenerator.ParseIPNetsParam(countParam, ipnetsParam)
	nets := ipnetsgenerator.GenerateIPNets(ipnetsReq, seed)

	var filter filter

	switch filterType {
	case "bpf":
		filter = bpf.New()
	case "iptables":
		filter = iptables.New()
	default:
		fmt.Printf("%q is not a valid filter type", filterType)
		os.Exit(1)
	}

	if err := filter.SetUp(nets, iface); err != nil {
		fmt.Printf("error setting up filter %s", err)
		return
	}

	defer filter.CleanUp()

	// TODO: run the actual benchmark here

	var input string
	fmt.Println("Print enter to exit: ")
	fmt.Scanf("%s", &input)
}
