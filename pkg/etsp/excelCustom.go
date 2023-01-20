package etsp

import (
	"github.com/xuri/excelize/v2"
)

func WriteHeadCustom(f *excelize.File, ssheet string) {
	writeHeadOne(f, ssheet, 1, 1, "Название детали", "")                                                                                                                                  // Name
	writeHeadOne(f, ssheet, 2, 1, "Описание детали", "")                                                                                                                                  // Note
	writeHeadOne(f, ssheet, 3, 1, "Артикул товара", "")                                                                                                                                   // GoodsCode
	writeHeadOne(f, ssheet, 4, 1, "Описание товара", "")                                                                                                                                  // GoodsComment
	writeHeadOne(f, ssheet, 5, 1, "Наличие на складе (True - имеется)", "")                                                                                                               // IsSklad
	writeHeadOne(f, ssheet, 6, 1, "Производитель", "")                                                                                                                                    // IsSklad
	writeHeadOne(f, ssheet, 7, 1, "Номер производителя", "")                                                                                                                              // ManufacturerName
	writeHeadOne(f, ssheet, 8, 1, "Цена клиента", "")                                                                                                                                     // Price
	writeHeadOne(f, ssheet, 9, 1, "Название склада", "")                                                                                                                                  // StorageName
	writeHeadOne(f, ssheet, 10, 1, "Кол-во товара в виде числа", "")                                                                                                                      // QuantityValue
	writeHeadOne(f, ssheet, 11, 1, "Вес (кг)", "")                                                                                                                                        // Weight
	writeHeadOne(f, ssheet, 12, 1, "Код детали", "")                                                                                                                                      // Code
	writeHeadOne(f, ssheet, 13, 1, "Уникальный номер", "")                                                                                                                                // UniqueNumber
	writeHeadOne(f, ssheet, 14, 1, "Код омеги (только по отдельному доступу)", "")                                                                                                        // OmegaNumber
	writeHeadOne(f, ssheet, 15, 1, "Код скубы (только по отдельному доступу)", "")                                                                                                        // SkubaNumber
	writeHeadOne(f, ssheet, 16, 1, "Группа (применяемость)", "")                                                                                                                          // Group
	writeHeadOne(f, ssheet, 17, 1, "Подгруппа (код подгруппы)", "")                                                                                                                       // Subgroup
	writeHeadOne(f, ssheet, 18, 1, "Код изображения детали (используется для получения изображения детали)", "")                                                                          // CodeImagePart
	writeHeadOne(f, ssheet, 19, 1, "Наличие сопутствующих деталией (признак используется для вызова списка сопустствующих деталей)", "")                                                  // HasPartAttendant
	writeHeadOne(f, ssheet, 20, 1, "Наличие в розничной сети (True - имеется)", "")                                                                                                       // IsShops
	writeHeadOne(f, ssheet, 21, 1, "Наличие в пути (True - имеется)", "")                                                                                                                 // IsShipment
	writeHeadOne(f, ssheet, 22, 1, "Наличие на внешних складах (True - имеется)", "")                                                                                                     // IsOutside
	writeHeadOne(f, ssheet, 23, 1, "Артикул клиента (только по отдельному доступу)", "")                                                                                                  // ClientArticle
	writeHeadOne(f, ssheet, 24, 1, "Признак аналога (True - аналог)", "")                                                                                                                 // IsAnalog
	writeHeadOne(f, ssheet, 25, 1, "Ориентировочная дата поступления товара в свободный остаток", "")                                                                                     // ApproximateIncomeDateInFreeStatus
	writeHeadOne(f, ssheet, 26, 1, "Идентификатор товара", "")                                                                                                                            // GoodsUnitID
	writeHeadOne(f, ssheet, 27, 1, "Признак наличия цены от количества (если True, сведения будут в группе QuantityDependence, отфильтровывать по комбинации GoodsUnitId+StorageId)", "") // HasPriceQuantityDependence
	writeHeadOne(f, ssheet, 28, 1, "Признак невозвратности (True - товар нельзя вернуть)", "")                                                                                            // IsNotRefundable
	writeHeadOne(f, ssheet, 29, 1, "Признак некондиции (True - некондиция)", "")                                                                                                          // IsSubstandard
	writeHeadOne(f, ssheet, 30, 1, "Кол-во товара в виде строки", "")                                                                                                                     // Quantity
	writeHeadOne(f, ssheet, 31, 1, "Признак ограниченного доступа к просмотру количества. Если доступ ограничен, то товар больше 8 показывается как >8", "")                              // QuantityWithRestrictions
	writeHeadOne(f, ssheet, 32, 1, "Идентификатор статуса товара", "")                                                                                                                    // RemainsStatusID
	writeHeadOne(f, ssheet, 33, 1, "Статус товара (свободный остаток, оформление, в пути и др.)", "")                                                                                     // RemainsStatusName
	writeHeadOne(f, ssheet, 34, 1, "Идентификатор склада", "")                                                                                                                            // StorageID
	writeHeadOne(f, ssheet, 35, 1, "Позиция склада (для сортировки)", "")                                                                                                                 // StoragePosition
}
func WriteOneLineCustom(f *excelize.File, ssheet string, row int, SearchBasicRes SearchBasicResponse, SearchBasicIndex int, GetPartsRemainsByCodeRes GetPartsRemainsByCodeResponse, GetPartsRemainsByCodeIndex int) {

	writeHeadOne(f, ssheet, 1, row, SearchBasicRes.Data.Items[SearchBasicIndex].Name, "")
	writeHeadOne(f, ssheet, 2, row, SearchBasicRes.Data.Items[SearchBasicIndex].Note, "")
	writeHeadOne(f, ssheet, 3, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].GoodsCode, "")
	writeHeadOne(f, ssheet, 4, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].GoodsComment, "")
	writeHeadOne(f, ssheet, 5, row, SearchBasicRes.Data.Items[SearchBasicIndex].IsSklad, "")
	writeHeadOne(f, ssheet, 6, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].ManufacturerName, "")
	writeHeadOne(f, ssheet, 7, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].ManufacturerNumber, "")
	writeHeadOne(f, ssheet, 8, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].Price, "")
	writeHeadOne(f, ssheet, 9, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].StorageName, "")
	writeHeadOne(f, ssheet, 10, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].QuantityValue, "")
	writeHeadOne(f, ssheet, 11, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].Weight, "")
	writeHeadOne(f, ssheet, 12, row, SearchBasicRes.Data.Items[SearchBasicIndex].Code, "")
	writeHeadOne(f, ssheet, 13, row, SearchBasicRes.Data.Items[SearchBasicIndex].UniqueNumber, "")
	writeHeadOne(f, ssheet, 14, row, SearchBasicRes.Data.Items[SearchBasicIndex].OmegaNumber, "")
	writeHeadOne(f, ssheet, 15, row, SearchBasicRes.Data.Items[SearchBasicIndex].SkubaNumber, "")
	writeHeadOne(f, ssheet, 16, row, SearchBasicRes.Data.Items[SearchBasicIndex].Group, "")
	writeHeadOne(f, ssheet, 17, row, SearchBasicRes.Data.Items[SearchBasicIndex].Subgroup, "")
	writeHeadOne(f, ssheet, 18, row, SearchBasicRes.Data.Items[SearchBasicIndex].CodeImagePart, "")
	writeHeadOne(f, ssheet, 19, row, SearchBasicRes.Data.Items[SearchBasicIndex].HasPartAttendant, "")
	writeHeadOne(f, ssheet, 20, row, SearchBasicRes.Data.Items[SearchBasicIndex].IsShops, "")
	writeHeadOne(f, ssheet, 21, row, SearchBasicRes.Data.Items[SearchBasicIndex].IsShipment, "")
	writeHeadOne(f, ssheet, 22, row, SearchBasicRes.Data.Items[SearchBasicIndex].IsOutside, "")
	writeHeadOne(f, ssheet, 23, row, SearchBasicRes.Data.Items[SearchBasicIndex].ClientArticle, "")
	writeHeadOne(f, ssheet, 24, row, SearchBasicRes.Data.Items[SearchBasicIndex].IsAnalog, "")
	writeHeadOne(f, ssheet, 25, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].ApproximateIncomeDateInFreeStatus, "")
	writeHeadOne(f, ssheet, 26, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].GoodsUnitID, "")
	writeHeadOne(f, ssheet, 27, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].HasPriceQuantityDependence, "")
	writeHeadOne(f, ssheet, 28, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].IsNotRefundable, "")
	writeHeadOne(f, ssheet, 29, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].IsSubstandard, "")
	writeHeadOne(f, ssheet, 30, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].Quantity, "")
	writeHeadOne(f, ssheet, 31, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].QuantityWithRestrictions, "")
	writeHeadOne(f, ssheet, 32, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].RemainsStatusID, "")
	writeHeadOne(f, ssheet, 33, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].RemainsStatusName, "")
	writeHeadOne(f, ssheet, 34, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].StorageID, "")
	writeHeadOne(f, ssheet, 35, row, GetPartsRemainsByCodeRes.Data.Remains[GetPartsRemainsByCodeIndex].StoragePosition, "")

}
