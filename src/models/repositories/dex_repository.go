package repositories

import (
	"fmt"
	"go-server/src/models/entities"
	"go-server/src/modules/dex/dto"
	"sync"

	"gorm.io/gorm"
)

type IDexRepository interface {
	ListDexes(conds ...interface{}) ([]entities.DexEntity, error)
	ListDexesWithPagination(listDexDto dto.ListDexDto) ([]entities.DexEntity, int64, error)
}
type DexRepository struct {
	db *gorm.DB
}

func NewDexRepository(db *gorm.DB) *DexRepository {
	return &DexRepository{db: db}
}

func (dexRepository *DexRepository) ListDexes(conds ...interface{}) ([]entities.DexEntity, error) {
	var dexes []entities.DexEntity
	err := dexRepository.db.Order("created_at desc").Find(&dexes, conds).Error
	return dexes, err
}
func (dexRepository *DexRepository) ListDexesWithPagination(listDexDto dto.ListDexDto) ([]entities.DexEntity, int64, error) {
	var dexes []entities.DexEntity
	var totalDex int64
	query := dexRepository.db.Model(&dexes)
	if listDexDto.Network != "" {
		query.Where("network = ?", listDexDto.Network)
	}
	errCount := query.Count(&totalDex).Error
	if errCount != nil {
		return nil, totalDex, errCount
	}
	errListDex := query.Limit(listDexDto.Limit).Offset(listDexDto.Offset).Find(&dexes).Error
	if errListDex != nil {
		return nil, totalDex, errListDex
	}
	return dexes, totalDex, nil
}

func CountDexs(query *gorm.DB, totalDex *int64, ch chan<- int64, wg *sync.WaitGroup) {
	fmt.Println("Startt")
	query.Count(totalDex)
	fmt.Println(*totalDex)
	ch <- *totalDex
	wg.Done()
}

func (dexRepository *DexRepository) ListDexesWithPaginationChannel(listDexDto dto.ListDexDto) ([]entities.DexEntity, int64, error) {
	var dexes []entities.DexEntity
	var dexes2 []entities.DexEntity
	var totalDex int64
	query := dexRepository.db.Model(&dexes)
	query2 := dexRepository.db.Model(&dexes2)
	if listDexDto.Network != "" {
		query.Where("network = ?", listDexDto.Network)
		query2.Where("network = ?", listDexDto.Network)
	}
	var wg = sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int64)
	go CountDexs(query2, &totalDex, ch, &wg)
	errListDex := query.Limit(listDexDto.Limit).Offset(listDexDto.Offset).Find(&dexes).Error
	if errListDex != nil {
		return nil, totalDex, errListDex
	}
	fmt.Println(dexes)
	i := <-ch
	fmt.Println("Done i")
	fmt.Println(i)
	return dexes, i, nil
}
