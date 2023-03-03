package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

const (
	fileName = "./handlingHttpRequest.webservices/Database/productDB.json"
)

var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]Product)}

func Init() {
	fmt.Println("loading products...")
	prodMap, err := loadProductMap()
	if err != nil {
		log.Fatal(err)
	}
	productMap.m = prodMap
	fmt.Printf("%d products loaded...\n", len(productMap.m))
}

func loadProductMap() (map[int]Product, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := os.ReadFile(fileName)
	productList := make([]Product, 0)
	err = json.Unmarshal([]byte(file), &productList)
	if err != nil {
		log.Fatal(err)
	}
	productMap := make(map[int]Product)
	for _, p := range productList {
		productMap[p.Id] = p
	}
	return productMap, nil
}

func GetProduct(productId int) *Product {
	productMap.RLock()
	defer productMap.RUnlock()
	if product, ok := productMap.m[productId]; ok {
		return &product
	}
	return nil
}

func RemoveProduct(productID int) {
	productMap.Lock()
	defer productMap.Unlock()
	delete(productMap.m, productID)
}

func GetProductList() []Product {
	productMap.RLock()
	products := make([]Product, 0, len(productMap.m))
	for _, value := range productMap.m {
		products = append(products, value)
	}
	productMap.RUnlock()
	return products
}

func getProductIds() []int {
	productIds := []int{}
	productMap.RLock()
	for key := range productMap.m {
		productIds = append(productIds, key)
	}
	productMap.RUnlock()
	sort.Ints(productIds)
	return productIds
}

func getNextProductID() int {
	productIds := getProductIds()
	return productIds[len(productIds)-1] + 1
}

func AddOrUpdateProduct(product Product) (int, error) {
	// if the product id is set, update, otherwise add
	addOrUpdateID := -1
	if product.Id > 0 {
		oldProduct := GetProduct(product.Id)
		// if it exists, replace it, otherwise return error
		if oldProduct == nil {
			return 0, fmt.Errorf("product id [%d] doesn't exist", product.Id)
		}
		addOrUpdateID = product.Id
	} else {
		addOrUpdateID = getNextProductID()
		product.Id = addOrUpdateID
	}
	productMap.Lock()
	productMap.m[addOrUpdateID] = product
	productMap.Unlock()
	return addOrUpdateID, nil
}
