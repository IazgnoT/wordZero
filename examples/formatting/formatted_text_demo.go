package main

import (
	"log"

	"github.com/IazgnoT/wordZero/pkg/document"
)

func main() {
	doc := document.New()

	// 创建不同格式的文本
	formats := []*document.TextFormat{
		{FontFamily: "Arial", FontSize: 12, Bold: true},
		{FontFamily: "Times New Roman", FontSize: 14, Italic: true},
		{FontFamily: "Courier New", FontSize: 10, FontColor: "FF0000"},
		{FontFamily: "Calibri", FontSize: 16, Bold: true, Italic: true},
	}

	texts := []string{
		"这是粗体文本",
		"这是斜体文本",
		"这是红色文本",
		"这是粗体斜体文本",
	}

	for i, text := range texts {
		para := doc.AddFormattedParagraph(text, formats[i])
		para.SetAlignment(document.AlignCenter)
	}

	// 保存文档
	err := doc.Save("examples/output/formatted_text_demo.docx")
	if err != nil {
		log.Fatalf("保存文档失败: %v", err)
	}

	log.Println("格式化文本演示文档创建成功！")
}
