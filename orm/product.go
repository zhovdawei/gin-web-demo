package orm

type Product struct {
	ID    uint
	Name  string
	Color string
	Price float32
}

func (product Product) Insert() {
	DataSource.Create(&product)
}

func (product Product) Delete() {
	DataSource.Delete(&product)
}

func (product Product) Update() {
	DataSource.Model(&product).Update("color", "草原绿")
	//DataSource.Update(&product, "color", "草原绿")
}

func (product *Product) Selete() {
	DataSource.First(&product)
}
