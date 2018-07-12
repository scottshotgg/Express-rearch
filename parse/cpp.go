package parse

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/scottshotgg/ExpressRedo/token"
)

/*
	This should really be an interface that can be implemented with
	certain functions so that other transpilations can be implemeneted
*/

var (
	f   *os.File
	r   *rand.Rand
	err error
)

func translateArray(t token.Value) {
	fmt.Printf("%+v\n", t)
	trueValue, ok := t.True.([]token.Token)
	if !ok {
		fmt.Println("shit look at t")
		fmt.Println("trueValue", trueValue)
		os.Exit(9)
	}

	// assuming only single type arrays until I have time to do multi type arrays in C
	arrayType := t.Type
	arrayValue := func() (valueString string) {
		for i, v := range trueValue {
			sprintString := "%v"
			if v.Value.Type == "string" {
				sprintString = "\"" + sprintString + "\""
			}
			valueString += fmt.Sprintf(sprintString, v.Value.True)
			if i != len(trueValue)-1 {
				valueString += ", "
			}
		}
		return
	}()
	thing := arrayType + " " + t.Name + "[] = { " + arrayValue + " };\n"
	fmt.Println(thing)
	_, err = f.Write([]byte(thing))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}
}

func translateVariableStatement(t token.Value) error {
	// if the token type is var make a var statement in C

	switch t.Type {
	case "var":
		// int abc = 5;
		// Any zyx = Any{ "int", &abc };
		varName := t.Name + strconv.Itoa(int(r.Uint32()))
		thing := strings.Join([]string{t.Acting, varName, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		thing += "Any " + t.Name + " = Any{ \"" + t.Acting + "\", &" + varName + " };\n"
		fmt.Println(thing)
		_, err = f.Write([]byte(thing))
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(9)
		}
		return nil

	case "object":
		// translateObject(t)
		return nil

	case "array":
		translateArray(t)
		return nil

	// In the case of the object we need to essentially instantiate a struct that will be used even if only temporarily
	// could just use that json library for now but wtf
	// fmt.Println("std::map<string, " + +"> " + t.Name)
	case "string":
		thing := "std::" + strings.Join([]string{t.Type, t.Name, "=", fmt.Sprintf("\"%v\"", t.True)}, " ") + ";\n"
		fmt.Println(thing)
		_, err = f.Write([]byte(thing))
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(9)
		}
		return nil

	case "int":
		fallthrough
	case "float":
		fallthrough
	case "bool":
		thing := strings.Join([]string{t.Type, t.Name, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		fmt.Println(thing)
		_, err = f.Write([]byte(thing))
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(9)
		}
		return nil

	default:
		fmt.Println("am i an error ???")
		return errors.New("i am not nil")
	}
	return errors.New("why am i here")
}

func translateIf(t token.Value) {
	fmt.Println("wtf")
	fmt.Printf("t %+v\n", t)

	_, err = f.Write([]byte(fmt.Sprintf("if (%s) ", t.String)))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	// metadata, ok := t.Metadata
	// if !ok {
	// 	fmt.Println("omfg error")
	// 	os.Exit(9)
	// }
	fmt.Println("metadata", t.Metadata)

	fmt.Println("t.True", t.True)
	// os.Exit(9)

	// // // body, ok := opMap["body"].True.([]token.Value)
	// // // if !ok {
	// // // 	fmt.Println("omfg error")
	// // // 	os.Exit(9)
	// // // }
	TranslateBlock(token.Value{
		Type: token.Block,
		True: t.True,
	})
	_, err = f.Write([]byte("\n"))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}
}

func translateLoop(t token.Value) error {
	fmt.Printf("t: %+v", t)
	if t.Type != "for" {
		return errors.New("blah")
	}

	fmt.Println()
	fmt.Println("t.Metadata", t.Metadata)
	// os.Exit(9)

	// tValue, ok := t.Metadata)
	// if !ok {
	// 	return errors.New("not the type")
	// }
	// fmt.Println("tValue", tValue)

	for k, v := range t.Metadata {
		fmt.Println("k, v", k, v)
	}

	loop := fmt.Sprintf("{\nint %s=%d;\nwhile (%s<%d) {\n",
		t.Metadata["start"].(token.Value).Name, t.Metadata["start"].(token.Value).True.(int), t.Metadata["start"].(token.Value).Name,
		t.Metadata["end"].(token.Value).True.(int))
	fmt.Println("loop", loop)
	_, err = f.Write([]byte(loop))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	fmt.Println("wtf is this")
	TranslateBlock(token.Value{
		Type: token.Block,
		True: t.True,
	})

	loopEnding := fmt.Sprintf("%s+=%d;}\n}\n", t.Metadata["start"].(token.Value).Name, t.Metadata["step"].(token.Value).True.(int))
	fmt.Println(loopEnding)
	_, err = f.Write([]byte(loopEnding))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	return nil
}

func TranslateBlock(tv token.Value) {
	_, err = f.Write([]byte("{\n"))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	insideBlock := tv.True.([]token.Value)
	fmt.Println("insideBlock", insideBlock[0])

	for _, t := range insideBlock {
		fmt.Println("insideBlock t", t)

		if err = translateVariableStatement(t); err != nil {
			fmt.Println("i am here translateVariableStatement", err)
			if err = translateLoop(t); err != nil {
				translateIf(t)
			}
		}
	}

	f.Write([]byte("}"))
}

func (p *Parser) Transpile(block token.Value) ([]string, error) {
	fmt.Println("yo waddup")

	fmt.Println("block", block)

	// fmt.Println(p.source)

	// fmt.Println("tokens", len(p.source))
	for _, value := range block.True.([]token.Value) {
		fmt.Println()
		fmt.Printf("value %+v\n", value)
	}
	f, err = os.Create("main.expr.shit")
	if err != nil {
		fmt.Println("got an err creating file")
		os.Exit(9)
	}

	// TODO: check all f.Write errors I guess
	f.Write([]byte("#include <map>\n#include <string>\n"))
	f.Write([]byte("struct Any { std::string type; void* data; };\n"))
	f.Write([]byte("int main()"))

	TranslateBlock(block)

	f.Close()

	return []string{}, nil
}