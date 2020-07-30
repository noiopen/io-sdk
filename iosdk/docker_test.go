package main

import "fmt"

func ExampleDockerVersion() {
	//*DryRunFlag = false
	DryRunPush("19.03.5", "!no docker")
	out, err := dockerVersion()
	fmt.Println(out, err)
	out, err = dockerVersion()
	fmt.Println(out, err)
	// Output:
	// docker version --format {{.Server.Version}}
	// 19.03.5 <nil>
	// docker version --format {{.Server.Version}}
	//  no docker
}

func ExampleDockerIP() {
	DryRunPush("192.168.192.168", "Error: cannot find")
	fmt.Println(*dockerIP("dummy"))
	fmt.Println(dockerIP("dummy"))
	// Output:
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} dummy
	// 192.168.192.168
	// docker inspect --format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} dummy
	// <nil>
}

func ExampleDockerInfo() {
	DryRunPush("!bad", "Info: hello")
	out, err := dockerInfo()
	fmt.Println(1, err, out+"*")
	out, err = dockerInfo()
	fmt.Println(2, err, out+"*")
	// Output:
	// docker info
	// 1 Docker is not running *
	// docker info
	// 2 <nil> Info: hello*
}
