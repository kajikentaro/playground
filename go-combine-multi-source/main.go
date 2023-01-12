package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"regexp"
)

func isNotTestFile(fi fs.FileInfo) bool {
	filename := fi.Name()

	r := regexp.MustCompile(`_test.go$`)
	isTestFile := r.MatchString(filename)
	return !isTestFile
}

func main() {
	// 名前を標準入力から受け取って開く
	var dirname string
	fmt.Print("Please enter target directory: ")
	fmt.Scan(&dirname)
	fset := token.NewFileSet()
	f, err := parser.ParseDir(fset, dirname, isNotTestFile, parser.Mode(0))
	if err != nil {
		log.Fatal(err)
	}

	// importの重複しない一覧を作成する
	importMap := map[string]bool{}
	ast.Inspect(f["main"], func(n ast.Node) bool {
		if v, isOk := n.(*ast.ImportSpec); isOk {
			importMap[v.Path.Value] = true
		}
		return true
	})

	// importのDeclを作成
	importDecl := ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{},
	}
	for k := range importMap {
		importDecl.Specs = append(importDecl.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: k,
			},
		})
	}

	// import以外のDeclのリストを作成
	theOtherDecls := []ast.Decl{}
	for _, file := range f["main"].Files {
		for _, decl := range file.Decls {
			if v, isOk := decl.(*ast.GenDecl); isOk {
				if v.Tok == token.IMPORT {
					continue
				}
			}
			theOtherDecls = append(theOtherDecls, decl)
		}
	}

	// 両方をくっつけて新しいfileにする
	newFile := &ast.File{
		Name:  ast.NewIdent("main"),
		Decls: append([]ast.Decl{&importDecl}, theOtherDecls...),
	}

	// result.go.tmp を作成
	file, err := os.OpenFile("result.go.tmp", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	format.Node(file, token.NewFileSet(), newFile)
}
