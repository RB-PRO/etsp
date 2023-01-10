package etsp

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func MakeWorkBook() (*excelize.File, error) {
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
func CloseXlsx(f *excelize.File) error {
	if err := f.SaveAs("etsp.xlsx"); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func Filter(f *excelize.File, ssheet string) {
	f.AutoFilter(ssheet, "A1:D4", &excelize.AutoFilterOptions{
		Column: "B", Expression: "x != blanks",
	})
}

func WriteOneLine(f *excelize.File, ssheet string, row int, SearchBasicRes SearchBasicResponse, SearchBasicIndex int, GetPartsRemainsByCodeRes GetPartsRemainsByCodeResponse, GetPartsRemainsByCodeIndex int) {
	// SearchBasic
	writeHeadOne(f, ssheet, 1, row, SearchBasicRes.Data.Items[SearchBasicIndex].Code, "")
	writeHeadOne(f, ssheet, 2, row, SearchBasicRes.Data.Items[SearchBasicIndex].Name, "")
	writeHeadOne(f, ssheet, 3, row, SearchBasicRes.Data.Items[SearchBasicIndex].Note, "")
	writeHeadOne(f, ssheet, 4, row, SearchBasicRes.Data.Items[SearchBasicIndex].UniqueNumber, "")
	writeHeadOne(f, ssheet, 5, row, SearchBasicRes.Data.Items[SearchBasicIndex].OmegaNumber, "")
	writeHeadOne(f, ssheet, 6, row, SearchBasicRes.Data.Items[SearchBasicIndex].SkubaNumber, "")
	writeHeadOne(f, ssheet, 7, row, SearchBasicRes.Data.Items[SearchBasicIndex].Group, "")
	writeHeadOne(f, ssheet, 8, row, SearchBasicRes.Data.Items[SearchBasicIndex].Subgroup, "")
	writeHeadOne(f, ssheet, 9, row, SearchBasicRes.Data.Items[SearchBasicIndex].CodeImagePart, "")
	writeHeadOne(f, ssheet, 10, row, existBool(SearchBasicRes.Data.Items[SearchBasicIndex].HasPartAttendant), "")
	writeHeadOne(f, ssheet, 11, row, existBool(SearchBasicRes.Data.Items[SearchBasicIndex].IsSklad), "")
	writeHeadOne(f, ssheet, 12, row, existBool(SearchBasicRes.Data.Items[SearchBasicIndex].IsShops), "")
	writeHeadOne(f, ssheet, 13, row, existBool(SearchBasicRes.Data.Items[SearchBasicIndex].IsShipment), "")
	writeHeadOne(f, ssheet, 14, row, existBool(SearchBasicRes.Data.Items[SearchBasicIndex].IsOutside), "")
	writeHeadOne(f, ssheet, 15, row, SearchBasicRes.Data.Items[SearchBasicIndex].ClientArticle, "")
	writeHeadOne(f, ssheet, 16, row, existBool(SearchBasicRes.Data.Items[SearchBasicIndex].IsAnalog), "")
	// GetPartsRemainsByCode
	writeHeadOne(f, ssheet, 17, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].ApproximateIncomeDateInFreeStatus, "")
	writeHeadOne(f, ssheet, 18, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].GoodsCode, "")
	writeHeadOne(f, ssheet, 19, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].GoodsComment, "")
	writeHeadOne(f, ssheet, 20, row, strconv.Itoa(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].GoodsUnitID), "")
	writeHeadOne(f, ssheet, 21, row, existBool(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].HasPriceQuantityDependence), "")
	writeHeadOne(f, ssheet, 22, row, existBool(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].IsNotRefundable), "")
	writeHeadOne(f, ssheet, 23, row, existBool(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].IsSubstandard), "")
	writeHeadOne(f, ssheet, 24, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].ManufacturerName, "")
	writeHeadOne(f, ssheet, 25, row, fmt.Sprintf("%g", GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].Price), "")
	writeHeadOne(f, ssheet, 26, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].Quantity, "")
	writeHeadOne(f, ssheet, 27, row, strconv.Itoa(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].QuantityValue), "")
	writeHeadOne(f, ssheet, 28, row, existBool(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].QuantityWithRestrictions), "")
	writeHeadOne(f, ssheet, 29, row, strconv.Itoa(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].RemainsStatusID), "")
	writeHeadOne(f, ssheet, 30, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].RemainsStatusName, "")
	writeHeadOne(f, ssheet, 31, row, strconv.Itoa(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].StorageID), "")
	writeHeadOne(f, ssheet, 32, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].StorageName, "")
	writeHeadOne(f, ssheet, 33, row, strconv.Itoa(GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].StoragePosition), "")
	writeHeadOne(f, ssheet, 33, row, fmt.Sprintf("%g", GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].Weight), "")
}

// Вернуть true-false
func existBool(boolean bool) string {
	if boolean {
		return "True"
	} else {
		return "False"
	}
}
func WriteHead(f *excelize.File, ssheet string) {
	// SearchBasic
	writeHeadOne(f, ssheet, 1, 1, "Code", "Код детали (используется для показа остатков на складах/магазинах, , в GetPartsRemainsByCode2 и других методах PartsRemains)")
	writeHeadOne(f, ssheet, 2, 1, "Name", "Название детали;")
	writeHeadOne(f, ssheet, 3, 1, "Note", "Описание детали;")
	writeHeadOne(f, ssheet, 4, 1, "UniqueNumber", "Уникальный номер;")
	writeHeadOne(f, ssheet, 5, 1, "OmegaNumber", "Код омеги (только по отдельному доступу*);")
	writeHeadOne(f, ssheet, 6, 1, "SkubaNumber", "Код скубы (только по отдельному доступу*);")
	writeHeadOne(f, ssheet, 7, 1, "Group", "Группа (применяемость);")
	writeHeadOne(f, ssheet, 8, 1, "Subgroup", "Подгруппа (код подгруппы);")
	writeHeadOne(f, ssheet, 9, 1, "CodeImagePart", "Код изображения детали (используется для получения изображения детали);")
	writeHeadOne(f, ssheet, 10, 1, "HasPartAttendant", "Наличие сопутствующих деталией (признак используется для вызова списка сопустствующих деталей);")
	writeHeadOne(f, ssheet, 11, 1, "IsSklad", "Наличие на складе (True - имеется);")
	writeHeadOne(f, ssheet, 12, 1, "IsShops", "Наличие в розничной сети (True - имеется);")
	writeHeadOne(f, ssheet, 13, 1, "IsShipment", "Наличие в пути (True - имеется);")
	writeHeadOne(f, ssheet, 14, 1, "IsOutside", "Наличие на внешних складах (True - имеется);")
	writeHeadOne(f, ssheet, 15, 1, "ClientArticle", "Артикул клиента (только по отдельному доступу*);")
	writeHeadOne(f, ssheet, 16, 1, "IsAnalog", "Признак аналога (True - аналог).")
	// GetPartsRemainsByCode
	writeHeadOne(f, ssheet, 17, 1, "ApproximateIncomeDateInFreeStatus", "Ориентировочная дата поступления товара в свободный остаток.")
	writeHeadOne(f, ssheet, 18, 1, "GoodsCode", "Артикул товара")
	writeHeadOne(f, ssheet, 19, 1, "GoodsComment", "Описание товара")
	writeHeadOne(f, ssheet, 20, 1, "GoodsUnitID", "Идентификатор товара")
	writeHeadOne(f, ssheet, 21, 1, "HasPriceQuantityDependence", "Признак наличия цены от количества (если True, сведения будут в группе QuantityDependence, отфильтровывать по комбинации GoodsUnitId+StorageId)")
	writeHeadOne(f, ssheet, 22, 1, "IsNotRefundable", "Признак невозвратности (True - товар нельзя вернуть)")
	writeHeadOne(f, ssheet, 23, 1, "IsSubstandard", "Признак некондиции (True - некондиция)")
	writeHeadOne(f, ssheet, 24, 1, "ManufacturerName", "Производитель")
	writeHeadOne(f, ssheet, 25, 1, "Price", "Цена клиента")
	writeHeadOne(f, ssheet, 26, 1, "Quantity", "Кол-во товара в виде строки**")
	writeHeadOne(f, ssheet, 27, 1, "QuantityValue", "Кол-во товара в виде числа**")
	writeHeadOne(f, ssheet, 28, 1, "QuantityWithRestrictions", "Признак ограниченного доступа к просмотру количества. Если доступ ограничен, то товар больше 8 показывается как >8")
	writeHeadOne(f, ssheet, 29, 1, "RemainsStatusID", "Идентификатор статуса товара")
	writeHeadOne(f, ssheet, 30, 1, "RemainsStatusName", "Статус товара (свободный остаток, оформление, в пути и др.)")
	writeHeadOne(f, ssheet, 31, 1, "StorageID", "Идентификатор склада")
	writeHeadOne(f, ssheet, 32, 1, "StorageName", "Название склада")
	writeHeadOne(f, ssheet, 33, 1, "StoragePosition", "Позиция склада (для сортировки)")
	writeHeadOne(f, ssheet, 33, 1, "Weight", "Вес (кг.)")
}

func writeHeadOne(f *excelize.File, ssheet string, col int, row int, val string, comment string) {
	collumn, _ := excelize.ColumnNumberToName(col)
	f.SetCellValue(ssheet, collumn+strconv.Itoa(row), val)
	if comment != "" {
		f.AddComment(ssheet, excelize.Comment{
			Cell:   collumn + strconv.Itoa(row),
			Author: "Бот",
			Runs: []excelize.RichTextRun{
				{Text: "Бот: ", Font: &excelize.Font{Bold: true}},
				{Text: comment},
			},
		})
	}
}
