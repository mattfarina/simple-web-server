package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mattfarina/simple-web-server/pkg/fs"

	"github.com/Masterminds/log-go"
	"github.com/Masterminds/log-go/impl/cli"
	"github.com/spf13/cobra"
)

func main() {

	var debug bool
	var trace bool
	var port string

	rootCmd := &cobra.Command{
		Use:   "simple-web-server",
		Short: "Serve a filesystem on the web",
		RunE: func(cmd *cobra.Command, args []string) error {

			http.Handle("/", fs.FileServer(http.Dir(args[0])))
			ul := fmt.Sprintf("localhost:%s", port)
			log.Infof("Receiving traffic on %s", ul)
			err := http.ListenAndServe(ul, nil)

			return err
		},
	}

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "phe port to serve on")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "provide debug level output")
	rootCmd.PersistentFlags().BoolVarP(&trace, "trace", "t", false, "provide trace level output, including debug level")
	rootCmd.PersistentFlags().Parse(os.Args)

	lgr := cli.NewStandard()
	if debug {
		lgr.Level = log.DebugLevel
	}
	if trace {
		lgr.Level = log.TraceLevel
	}
	log.Current = lgr

	log.Debugf("Port: %s", port)
	log.Debugf("Debug: %t", debug)
	log.Debugf("Trace: %t", trace)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
