package pdf

import (
	"fmt"
)

func Generator() {

	r := NewRequestPdf("")

	//html template path
	templatePath := "pdf/templates/sample.html"

	//path for download pdf
	outputPath := "./pdf/storage/pdf.pdf"

	//html template data
	templateData := struct {
		Data string
	}{
		Data: "مرحبا",
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, err := r.GeneratePDF(outputPath)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(ok, "pdf generated successfully")
		}
	} else {
		fmt.Println("err:", err)
	}
}
