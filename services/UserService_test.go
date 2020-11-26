package UserService_test

var _ = Describe("UserService", func() {
	// var (
	//     longBook  Book
	//     shortBook Book
	// )
	//
	// BeforeEach(func() {
	//     longBook = Book{
	//         Title:  "Les Miserables",
	//         Author: "Victor Hugo",
	//         Pages:  1488,
	//     }
	//
	//     shortBook = Book{
	//         Title:  "Fox In Socks",
	//         Author: "Dr. Seuss",
	//         Pages:  24,
	//     }
	// })

	Describe("UserService", func() {
		It("throws error ", func() {
			Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
		})

	})
})

// var storage *gorm.DB
//
// func Storage() *gorm.DB {
// 	return storage
// }
//
// func InitDB() {
// 	dsn := "host=localhost user=memories_user password=memories_password dbname=memories_dev port=5432 sslmode=disable"
// 	var err error
// 	storage, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 	}
// }
