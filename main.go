package main

import (
	"context"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	_ "github.com/macwinux/benthos-knative-plugin/bloblang"
	"github.com/benthosdev/benthos/v4/public/service"
)

func main() {
	service.RunCLI(context.Background())
}
