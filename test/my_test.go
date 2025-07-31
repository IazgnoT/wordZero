package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ZeroHawkeye/wordZero/pkg/document"
)

func TestWordTempProcessor(t *testing.T) {
	renderer := document.NewTemplateRenderer()

	// 加载模板
	renderer.LoadTemplateFromFile("complex_template", "/home/wenten/code/wordZero/templates/generated_complex_template.docx")

	// 分析模板结构
	analysis, err := renderer.AnalyzeTemplate("complex_template")
	if err != nil {
		log.Fatal(err)
	}

	// 查看变量列表
	fmt.Println("模板变量:")
	for varName := range analysis.Variables {
		fmt.Printf("  - {{%s}}\n", varName)
	}

	// 查看条件列表
	fmt.Println("条件变量:")
	for condName := range analysis.Conditions {
		fmt.Printf("  - {{#if %s}}\n", condName)
	}

	// 查看列表变量
	fmt.Println("循环列表:")
	for listName := range analysis.Lists {
		fmt.Printf("  - {{#each %s}}\n", listName)
	}
}
