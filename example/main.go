package main

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var args = `{
          "input"                     : "data/data/feature_class.json",
          "output"                    : "data/output/out_feat_class.json",
          "constraints"               : "data/data/feature_class_const.json",
          "simplification_type"       : "DP",
          "threshold"                 : 50.0,
          "minimum_distance"          : 20.0,
          "relax_distance"            : 10.0,
          "is_feature_class"          : true,
          "planar_self"               : true,
          "non_planar_self"           : true,
          "avoid_new_self_intersects" : true,
          "geometric_relation"        : true,
          "distance_relation"         : true,
          "homotopy_relation"         : true
        }`

var simplifyExec = "simplify"
var execDir string

func init() {
	var ex, err = os.Executable()
	if err != nil {
		panic(err)
	}
	execDir = filepath.Dir(ex)
	simplifyExec = fmt.Sprintf("%v/%v", execDir, simplifyExec)
}

func main() {
	fmt.Println(simplifyExec)
	runPlugin(encode64(args))
}

func runPlugin(arg string) {
	var cmd *exec.Cmd
	cmd = exec.Command(simplifyExec, "-b", arg)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Execution Error: %s\n", err)
	}
}

func encode64(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}
