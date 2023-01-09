package app

import (
	"github.com/xuri/excelize/v2"
)

func makeWorkBook() (*excelize.File, error) {
	// Создать книгу Excel
	f := excelize.NewFile()
	// Create a new sheet.
	_, err := f.NewSheet("main")
	if err != nil {
		return f, err
	}
	f.DeleteSheet("Sheet1")
	return f, nil
}
func closeXlsx(f *excelize.File) error {
	if err := f.SaveAs("etsp.xlsx"); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func writeHead(f *excelize.File, ssheet string) {
	writeHeadOne(f, ssheet, 1, "Code", "Код детали (используется для показа остатков на складах/магазинах, , в GetPartsRemainsByCode2 и других методах PartsRemains)")
	writeHeadOne(f, ssheet, 2, "Name", "Название детали;")
	writeHeadOne(f, ssheet, 3, "Note", "Описание детали;")
	writeHeadOne(f, ssheet, 4, "UniqueNumber", "Уникальный номер;")
	writeHeadOne(f, ssheet, 5, "OmegaNumber", "Код омеги (только по отдельному доступу*);")
	writeHeadOne(f, ssheet, 6, "SkubaNumber", "Код скубы (только по отдельному доступу*);")
	writeHeadOne(f, ssheet, 7, "Group", "Группа (применяемость);")
	writeHeadOne(f, ssheet, 8, "Subgroup", "Подгруппа (код подгруппы);")
	writeHeadOne(f, ssheet, 9, "CodeImagePart", "Код изображения детали (используется для получения изображения детали);")
	writeHeadOne(f, ssheet, 10, "HasPartAttendant", "Наличие сопутствующих деталией (признак используется для вызова списка сопустствующих деталей);")
	writeHeadOne(f, ssheet, 11, "IsSklad", "Наличие на складе (True - имеется);")
	writeHeadOne(f, ssheet, 12, "IsShops", "Наличие в розничной сети (True - имеется);")
	writeHeadOne(f, ssheet, 13, "IsShipment", "Наличие в пути (True - имеется);")
	writeHeadOne(f, ssheet, 14, "IsOutside", "Наличие на внешних складах (True - имеется);")
	writeHeadOne(f, ssheet, 15, "ClientArticle", "Артикул клиента (только по отдельному доступу*);")
	writeHeadOne(f, ssheet, 16, "IsAnalog", "Признак аналога (True - аналог).")
}

func writeHeadOne(f *excelize.File, ssheet string, col int, val string, comment string) {
	collumn, _ := excelize.ColumnNumberToName(col)
	f.SetCellValue(ssheet, collumn+"1", val)
	f.AddComment(ssheet, excelize.Comment{
		Cell:   collumn + "1",
		Author: "Бот",
		Runs: []excelize.RichTextRun{
			{Text: "Бот: ", Font: &excelize.Font{Bold: true}},
			{Text: comment},
		},
	})
}
