package middlewares

import (
	"mcs_api/src/models"

	"github.com/jung-kurt/gofpdf"
)

func CreateServicePdf(company *models.Company, machine *models.Machine, service *models.Service, outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Configurar la fuente
	pdf.SetFont("Arial", "B", 12)

	// Dibujar encabezado
	pdf.CellFormat(190, 10, "ORDEN DE TRABAJO", "0", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 10, "MANTENIMIENTO CORRECTIVO", "0", 1, "C", false, 0, "")

	// Información del cliente
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(30, 8, "Cliente:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(100, 8, "Nombre: "+company.Name, "0", 0, "L", false, 0, "")
	pdf.CellFormat(60, 8, "Cel: ____________________", "0", 1, "L", false, 0, "")

	// Información de la máquina
	pdf.Ln(5)
	pdf.CellFormat(30, 8, "Maquina:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, "Marca: "+machine.BrandId, "0", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, "Modelo: "+machine.Model, "0", 0, "L", false, 0, "")
	pdf.CellFormat(40, 8, "Serie: "+machine.Serial, "0", 1, "L", false, 0, "")

	// Detalle del servicio
	pdf.Ln(5)
	pdf.CellFormat(190, 8, "DETALLE DEL SERVICIO", "1", 1, "C", false, 0, "")
	for i := 0; i < 6; i++ {
		pdf.CellFormat(190, 8, "", "1", 1, "L", false, 0, "")
	}

	// Problemas encontrados y soluciones
	pdf.Ln(5)
	pdf.CellFormat(95, 8, "PROBLEMAS ENCONTRADOS", "1", 0, "C", false, 0, "")
	pdf.CellFormat(95, 8, "SOLUCIONES", "1", 1, "C", false, 0, "")
	for i := 0; i < 6; i++ {
		pdf.CellFormat(95, 8, "", "1", 0, "L", false, 0, "")
		pdf.CellFormat(95, 8, "", "1", 1, "L", false, 0, "")
	}

	// Observaciones y materiales
	pdf.Ln(5)
	pdf.CellFormat(95, 8, "OBSERVACIONES", "1", 0, "C", false, 0, "")
	pdf.CellFormat(55, 8, "MATERIALES", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 8, "Costo Bs.", "1", 1, "C", false, 0, "")
	for i := 0; i < 6; i++ {
		pdf.CellFormat(95, 8, "", "1", 0, "L", false, 0, "")
		pdf.CellFormat(55, 8, "", "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 8, "", "1", 1, "L", false, 0, "")
	}

	// Firmas
	pdf.Ln(5)
	pdf.CellFormat(95, 8, "FIRMA CLIENTE:", "0", 0, "L", false, 0, "")
	pdf.CellFormat(95, 8, "FIRMA RESPONSABLE:", "0", 1, "L", false, 0, "")
	pdf.Ln(2)
	pdf.CellFormat(95, 8, "Nombre: ____________________", "0", 0, "L", false, 0, "")
	pdf.CellFormat(95, 8, "Nombre: ____________________", "0", 1, "L", false, 0, "")
	pdf.CellFormat(95, 8, "CI: ____________________", "0", 0, "L", false, 0, "")
	pdf.CellFormat(95, 8, "CI: ____________________", "0", 1, "L", false, 0, "")

	// Fecha y total
	pdf.Ln(5)
	pdf.CellFormat(95, 8, "FECHA Y HORA DE ENTREGA: ____________________", "0", 0, "L", false, 0, "")
	pdf.CellFormat(95, 8, "TOTAL COSTO DEL SERVICIO: ____________________", "0", 1, "L", false, 0, "")
	//

	return pdf.OutputFileAndClose(outputPath)
}
