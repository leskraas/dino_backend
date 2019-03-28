package synthesis

import (
	"fmt"
	"log"
	"os/exec"
)

func executeYosys(path string) {
	app := "yosys/yosys"
	arg0 := []string{"-p", "prep -auto-top; write_json larstestOutAigmap.json", path}
	yosysCmd := exec.Command(app, arg0...) //, arg1, arg2)

	err := yosysCmd.Run()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

}

func executeNetlistsvg() {
	app := "node"
	arg := []string{"netlistsvg/bin/netlistsvg", "larstestOutAigmap.json", "-o", "larsern2Aigmap.svg"}
	netlistsvgCmd := exec.Command(app, arg...)

	err := netlistsvgCmd.Run()
	if err != nil {
		log.Fatalf("netlistsvgCmd.Run() failed with '%s'\n", err)
	}
}

func displaySvg() {
	app := "display"
	arg := "larsern2Aigmap.svg"
	displayCmd := exec.Command(app, arg)
	err := displayCmd.Run()

	if err != nil {
		log.Fatalf("displayCmd.Run() failed with '%s'\n", err)
	}
}

func RunSynthesis() {
	fmt.Println("Running")
	executeYosys("yosys/tests/simple/larserik/up_counter.v")
	executeNetlistsvg()
	displaySvg()
	fmt.Println("Done")

}

// /Users/larserikskraastad/Documents/ntnu/master/program/yosys/./yosys -p "prep -auto-top; aigmap; write_json testOutJson.json" /Users/larserikskraastad/Documents/ntnu/master/program/yosys/tests/simple/larserik/counter.v
