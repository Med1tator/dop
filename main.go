package main

import (
	"dop/cmd"
	"log"
)

// # 2
//var name word

func main() {
	// # 1
	//var name word
	//flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	//flag.StringVar(&name, "n", "Go语言编程之旅", "帮助信息")
	//flag.Parse()
	//
	//log.Printf("name: %s\n", name)

	// # 2
	//flag.Parse()
	//goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	//goCmd.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
	//
	//phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	//phpCmd.StringVar(&name, "n", "PHP语言编程之旅", "帮助信息")
	//
	//args := flag.Args()
	//switch args[0] {
	//case "go":
	//	_ = goCmd.Parse(args[1:])
	//case "php":
	//	_ = phpCmd.Parse(args[1:])
	//}
	//
	//log.Printf("name: %s\n", name)

	// # 3
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v\n", err)
	}
}
