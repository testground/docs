package main

import (
    "context"
    "fmt"
    "os"

    "dagger.io/dagger"
)

func main() {
    if err := build(context.Background()); err != nil {
        fmt.Println(err)
    }
}

func build(ctx context.Context) error {
    fmt.Println("Building with Dagger")

    // initialize Dagger client
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        return err
    }
    defer client.Close()

    // get reference to the local project
    src := client.Host().Workdir()

    // get `node` image
    node := client.Container().From("node:18-alpine")

    // mount cloned repository into `golang` image
    node = node.WithMountedDirectory("/src", src).WithWorkdir("/src")


    // define the application build command
    path := "build/"
    node = node.
			Exec(dagger.ContainerExecOpts{
        Args: []string{"npm", "install", "docsify-cli", "-g"},
    	}).
			Exec(dagger.ContainerExecOpts{
				Args: []string{"gitbook", "build", "--gitbook=4.0.0-alpha.2"},
			})

    // get reference to build output directory in container
    output := node.Directory(path)

    // write contents of container build/ directory to the host
    _, err = output.Export(ctx, path)
    if err != nil {
        return err
    }

    return nil
}
