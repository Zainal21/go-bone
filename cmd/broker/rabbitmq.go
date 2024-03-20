package broker

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Zainal21/go-bone/app/bootstrap"
	"github.com/Zainal21/go-bone/app/controller"
	"github.com/Zainal21/go-bone/pkg/config"
	"github.com/Zainal21/go-bone/pkg/logger"
)

var (
	flags  = flag.NewFlagSet("rabbit", flag.ExitOnError)
	name   = flags.String("name", "", "queue and exchange name")
	topics = flags.String("topics", "", "topic to subscribes")
	help   = flags.Bool("guide", false, "Print Help")
)

func ServeRabbitMQ() {
	flags.Usage = usage
	_ = flags.Parse(os.Args[2:])

	args := flags.Args()

	if (len(args) == 0 && (*name == "" || *topics == "")) || *help {
		flags.Usage()
		return
	}

	var topicList []string
	for _, t := range strings.Split(*topics, "|") {
		if t != "" {
			topicList = append(topicList, t)
		}
	}

	logger.SetJSONFormatter()
	cnf, err := config.LoadAllConfigs()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to load configuration file: %v", err))
	}
	mController := controller.NewLogController()

	subs := bootstrap.RegistryRabbitMQSubscriber(*name, cnf, mController)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed connect to RabbitMQ Cause: %v", err))
	}

	err = subs.Listen(topicList)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to Listen Topic Cause: %v", err))
	}
}

func usage() {
	fmt.Println(usageCommands)
}

var (
	usageCommands = `
	Usage:
		go run main.go rabbit [flags]

	Flags:
	--name 		string    	name of queue and exchange
	--topics 	string		separate with pipeline "|" if want multiple bindings
  	--guide          		print help
`
)
