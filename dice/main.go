package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gebv/go-dice/filepathx"
	"github.com/gebv/go-dice/gen"
	"github.com/gebv/go-dice/parse"
)

var VERSION = "???"
var BUILDATE = "???"

var fToolFormat = flag.String("tool-format", "goreturns", "Tool which format the code. Available: goreturns.")

func main() {
	log.SetFlags(0)
	log.SetPrefix("dice: ")
	flag.Parse()

	log.Println("Verions", VERSION)
	log.Println("Build date", BUILDATE)

	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("failed get work directory", err)
	}

	file := os.Getenv("GOFILE")
	packName := os.Getenv("GOPACKAGE")

	srcFileName := filepath.Join(baseDir, file)

	log.Println("Process file:", srcFileName, packName)

	// /foo/bar/<orig_file_name>_dice.go
	var dstFileName = filepathx.FileNameWithSuffix(srcFileName, "dice")
	var dstFile *os.File

	for _, node := range parse.Parse(srcFileName) {
		grules := parseGenRules(node.DocComments())
		if len(grules) == 0 {
			// skip if did not find any command
			continue
		}
		if dstFile == nil {
			dstFile, err = os.Create(dstFileName)
			if err != nil {
				log.Fatalln("failed create file", err)
			}

			packageComments := fmt.Sprintf(
				defaultPackageComments,
				srcFileName,
				flag.Args(),
				packName,
				VERSION,
			)
			gen.Package(dstFile, packName, packageComments)
		}

		for _, g := range grules {
			if g.Match(node) {
				printGenRuleInfo(dstFile, g, node)
				gen.Comments(dstFile, "dice: START code")
				fmt.Fprintln(dstFile)
				if err := g.Gen(dstFile, g.Cfg, node); err != nil {
					gen.Comments(dstFile, fmt.Sprintf("dice: ERROR %q", err))
				}
				fmt.Fprintln(dstFile)
				gen.Comments(dstFile, "dice: END code")
				fmt.Fprintln(dstFile)
			}
		}
	}
	dstFile.Close()

	// formatting after generation
	formatCode(dstFileName)

	switch *fToolFormat {
	case "goreturns":
		// go imports will fix all the imports
		runGoreturns(dstFileName)
	}

	// fixes formatting for generated code
	runGofmt(dstFileName)
}

func formatCode(file string) {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed open file=%q %v", file, err)
	}

	ffileBytes, err := format.Source(fileBytes)
	if err != nil {
		log.Fatal("failed format code", err)
	}

	if err := ioutil.WriteFile(file, ffileBytes, 0644); err != nil {
		log.Fatal("failed write formated code", err)
	}
}

func runGoreturns(file string) {
	cmd := exec.Command("goreturns", "-w", file)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("failed goreturns", err)
	}
}

func runGofmt(file string) {
	cmd := exec.Command("gofmt", "-s", "-w", file)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("failed goreturns", err)
	}
}

func parseGenRules(in string) []*gen.GeneratorRule {
	in = strings.TrimSpace(in)
	if in == "" {
		return []*gen.GeneratorRule{}
	}
	res := []*gen.GeneratorRule{}
	for _, ann := range gen.ParseAnnotations(in) {
		grule := gen.NewGeneratorRule(ann)
		if grule == nil {
			continue
		}
		if grule.IsValid() {
			res = append(res, grule)
		}
	}
	return res
}

func printGenRuleInfo(w io.Writer, grule *gen.GeneratorRule, node parse.Node) {
	gen.Comments(w, "dice: START meta information")
	gen.Comments(w, fmt.Sprintf("File position: %q", node.NodePosition))
	gen.Comments(w, fmt.Sprintf("Annotation name: %q", grule.AnnotaionName))
	buff := new(bytes.Buffer)
	err := json.NewEncoder(buff).Encode(grule.Cfg)
	if err != nil {
		log.Fatal("Failed encode genrule config", err)
	}
	gen.Comments(w, "Rule config:")
	gen.Comments(w, buff.String())
	gen.Comments(w, "dice: END meta information")
}

// TODO: customize?
var defaultPackageComments = `Code generated by github.com/gebv/go-dice with details DO NOT EDIT.
File: %v
Porcess with args: %v
Package name: %v
Version: %v`
