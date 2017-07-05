package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/ginuerzh/gost"
	"github.com/golang/glog"
)

var (
	options struct {
		ChainNodes, ServeNodes flagStringList
	}
)

func init() {
	/*var (
		configureFile string
		printVersion  bool
	)

	flag.StringVar(&configureFile, "C", "gost.json", "configure file")
	flag.Var(&options.ChainNodes, "F", "forward address, can make a forward chain")
	flag.Var(&options.ServeNodes, "L", "listen address, can listen on multiple ports")
	flag.BoolVar(&printVersion, "V", false, "print version")
	flag.Parse()

		if err := loadConfigureFile(configureFile); err != nil {
			glog.Fatal(err)
		}

		if glog.V(5) {
			http2.VerboseLogs = true
		}

			if flag.NFlag() == 0 {
				flag.PrintDefaults()
				return
			}

		if printVersion {
			fmt.Fprintf(os.Stderr, "gost %s (%s)\n", gost.Version, runtime.Version())
			return
		}*/
}

func Gost(conf Config) {
	options.ServeNodes = conf.ServeNodes
	options.ChainNodes = conf.ChainNodes
	fmt.Println(options.ServeNodes)
	fmt.Println(options.ChainNodes)
	chain := gost.NewProxyChain()
	if err := chain.AddProxyNodeString(options.ChainNodes...); err != nil {
		glog.Fatal(err)
	}
	chain.Init()

	var wg sync.WaitGroup
	for _, ns := range options.ServeNodes {
		serverNode, err := gost.ParseProxyNode(ns)
		if err != nil {
			glog.Fatal(err)
		}

		wg.Add(1)
		go func(node gost.ProxyNode) {
			defer wg.Done()
			certFile, keyFile := node.Get("cert"), node.Get("key")
			if certFile == "" {
				certFile = gost.DefaultCertFile
			}
			if keyFile == "" {
				keyFile = gost.DefaultKeyFile
			}
			cert, err := gost.LoadCertificate(certFile, keyFile)
			if err != nil {
				glog.Fatal(err)
			}
			server := gost.NewProxyServer(node, chain, &tls.Config{Certificates: []tls.Certificate{cert}})
			glog.Fatal(server.Serve())
		}(serverNode)
	}
	wg.Wait()
}

func loadConfigureFile(configureFile string) error {
	if configureFile == "" {
		return nil
	}
	content, err := ioutil.ReadFile(configureFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(content, &options); err != nil {
		return err
	}
	return nil
}

type flagStringList []string

func (this *flagStringList) String() string {
	return fmt.Sprintf("%s", *this)
}
func (this *flagStringList) Set(value string) error {
	*this = append(*this, value)
	return nil
}
