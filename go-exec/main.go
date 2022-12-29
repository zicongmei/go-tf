package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func main() {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.1.9")),
	}

	fmt.Println("installer version: ", installer.Version.String())
	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}
	fmt.Println("execPath: ", execPath)

	workingDir := "tfsample"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}
	fmt.Println("init done")

	err = tf.Apply(context.Background())
	if err != nil {
		log.Fatalf("error running apply: %s", err)
	}
	fmt.Println("apply done")

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	graph, err := tf.Graph(context.Background())
	if err != nil {
		log.Fatalf("error running graph: %s", err)
	}

	fmt.Println("state.FormatVersion: ", state.FormatVersion) // "0.1"
	fmt.Println(state.Values.Outputs)
	fmt.Println("graph: ", graph)

	err = tf.Destroy(context.Background())
	if err != nil {
		log.Fatalf("error running destroy: %s", err)
	}
	fmt.Println("destroy done")
}
