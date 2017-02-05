package models

// Image is an array of bytes of any image
type Image []byte

// Product model of something we have to sell
type Product struct {
	Title       string
	Price       float32
	Description string
	Images      []Image
}

// IndexViewModel is used by our index page
type IndexViewModel struct {
	Title    string
	Products []Product
}
