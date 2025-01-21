package middlewares

import (
	"fmt"
	"mcs_api/src/models"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

func CreateServicePdf(company *models.Company, machine *models.Machine, service *models.Service, outputPath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Configurar la fuente
	pdf.SetFont("Arial", "B", 12)

	// Dibujar encabezado
	pdf.CellFormat(190, 10, "FICHA TECNICA", "0", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 10, "MANTENIMIENTO CORRECTIVO", "0", 1, "C", false, 0, "")

	// Información del cliente
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(25, 8, "Cliente", "0", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(63.3, 8, "Empresa/Taller: "+company.Name, "0", 0, "L", false, 0, "")
	pdf.CellFormat(63.4, 8, "Encargado: "+company.Manager, "0", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 8, "Cel: "+"69804340", "0", 1, "R", false, 0, "")

	// Información de la máquina
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(25, 8, "Maquina", "0", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(47.5, 8, "Maquina: "+"recta", "0", 0, "L", false, 0, "")
	pdf.CellFormat(47.5, 8, "Marca: "+"Juki", "0", 0, "L", false, 0, "")
	pdf.CellFormat(47.5, 8, "Modelo: "+machine.Model, "0", 0, "L", false, 0, "")
	pdf.CellFormat(47.5, 8, "Serie: "+machine.Serial, "0", 1, "R", false, 0, "")

	// Estado de la maquina
	pdf.Ln(5)
	pdf.CellFormat(190, 8, "ESTADO DE LA MAQUINA", "1", 1, "C", false, 0, "")
	for i := 0; i < 6; i++ {
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(10, 8, "P8", "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 8, "Mecanismo de puntada", "1", 0, "L", false, 0, "")
		pdf.SetTextColor(98, 98, 98) // operacion limitada
		// pdf.SetTextColor(8, 101, 253) // operacion limitada
		// pdf.SetTextColor(182, 31, 43) // no opertivo
		// pdf.SetTextColor(6, 104, 36) // operativo
		pdf.CellFormat(90, 8, "Operativo", "1", 1, "L", false, 0, "")
	}
	pdf.SetTextColor(0, 0, 0)

	// Problemas encontrados y soluciones
	pdf.Ln(5)
	pdf.CellFormat(100, 8, "PROBLEMAS ENCONTRADOS", "1", 0, "C", false, 0, "")
	pdf.CellFormat(90, 8, "SOLUCIONES", "1", 1, "C", false, 0, "")
	for i := 0; i < 6; i++ {
		pdf.CellFormat(10, 8, "P8", "1", 0, "C", false, 0, "")
		pdf.CellFormat(90, 8, "Buje ya gastado", "1", 0, "L", false, 0, "")
		pdf.CellFormat(90, 8, "Se reemplazo el buje", "1", 1, "L", false, 0, "")
	}

	// Materiales
	pdf.Ln(5)
	pdf.CellFormat(190, 8, "MATERIALES", "1", 1, "C", false, 0, "")
	pdf.CellFormat(20, 8, "Cantidad", "1", 0, "C", false, 0, "")
	pdf.CellFormat(120, 8, "Material", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "Costo Bs.", "1", 1, "C", false, 0, "")
	for _, v := range service.Materials {
		fmt.Println(v.Price)
		fmt.Println(strconv.FormatFloat(v.Price, 'f', -1, 64))
		pdf.CellFormat(20, 8, strconv.Itoa(v.Number), "1", 0, "C", false, 0, "")
		pdf.CellFormat(120, 8, v.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(50, 8, strconv.FormatFloat(v.Price, 'f', -1, 64), "1", 1, "C", false, 0, "")
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
