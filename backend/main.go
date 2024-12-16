package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID          int
	Name        string
	Image       string
	Cost        float64
	Describtion string
	Favorite    bool
	ShopCart    bool
	Count       int
}

type User struct {
	ID    int
	Image string
	Name  string
	Phone string
	Mail  string
}

var users = []User{{
	ID:    1,
	Image: "https://avatars.githubusercontent.com/u/125373457?v=4",
	Name:  "Калагов Марат",
	Phone: "8(900)000-00-00",
	Mail:  "@marattttt.ka",
}}

var products = []Product{
	{
		ID:          1,
		Name:        "Nike Air Jordan 1",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/3fa/500_500_1/ptbxe5k14n91ojfpq6vv28e3hljakg62.jpg",
		Cost:        15990,
		Describtion: "Классические баскетбольные кроссовки с культовым дизайном. Nike Air Jordan 1 — это первая модель кроссовок, созданная для легендарного баскетболиста Майкла Джордана. Обладают стильным внешним видом и отличной поддержкой.",
	},
	{
		ID:          2,
		Name:        "Adidas Stan Smith",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/e48/500_500_1/xdzcj237213peo9n3hcv1tm3obnyk0ub.jpg",
		Cost:        7590,
		Describtion: "Adidas Stan Smith — это классические кроссовки, которые стали иконой стиля. Их простой дизайн и комфорт делают их идеальными для повседневной носки.",
	},
	{
		ID:          3,
		Name:        "Nike Air Max 90",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/4aa/500_500_1/zythg01bjon2ucpg0qfi2qvb1p13euzl.jpg",
		Cost:        19990,
		Describtion: "Nike Air Max 90 — это первая модель с видимой амортизацией Air. Они обеспечивают отличную поддержку и комфорт, идеально подходят для активного образа жизни.",
	},
	{
		ID:          4,
		Name:        "Converse All Star",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/97c/500_500_1/4xr91kquh8xuriigo8wwrg8mpckbcvyd.JPG",
		Cost:        9390,
		Describtion: "Converse Chuck Taylor All Star — это универсальные кеды, которые подходят к любому стилю. Их стильный и простой дизайн делает их популярными среди молодежи.",
	},
	{
		ID:          5,
		Name:        "Adidas Superstar",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/d29/500_500_1/y3hijlixmv4p5qqey340xgl4ax6jkl4a.jpg",
		Cost:        11890,
		Describtion: "Adidas Superstar — это культовые кроссовки, которые изначально были созданы для баскетбола. Их уникальный дизайн и удобство сделали их популярными во всем мире.",
	},
	{
		ID:          6,
		Name:        "Nike Dunk",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/0db/500_500_1/mupah495oocq0h9vq2ntrjgn212fkc2y.JPG",
		Cost:        16790,
		Describtion: "Nike Dunk — это стильные и удобные кроссовки, которые подходят как для баскетбольной площадки, так и для повседневной жизни. Доступны в различных цветах и дизайнах.",
	},
	{
		ID:          7,
		Name:        "New Balance 574",
		Image:       "https://static.street-beat.ru/upload/resize_cache/iblock/4ae/500_500_1/vmx03wmdx054s63qxfc4op3inb9m76d7.jpg",
		Cost:        15190,
		Describtion: "New Balance 574 — это идеальные кроссовки для тех, кто ценит стиль и комфорт. Их мягкая подошва и поддержка делают их идеальными для долгих прогулок.",
	},
}

// обработчик для GET-запроса, возвращает список продуктов
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)

	newProduct.ID = products[len(products)-1].ID + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

//Добавление маршрута для получения одного продукта

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Обновление продукта по id
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Ищем продукт для обновления
	for i, Product := range products {
		if Product.ID == id {
			products[i].Name = updatedProduct.Name
			products[i].Image = updatedProduct.Image
			products[i].Cost = updatedProduct.Cost
			products[i].Describtion = updatedProduct.Describtion
			products[i].Favorite = updatedProduct.Favorite
			products[i].ShopCart = updatedProduct.ShopCart
			products[i].Count = updatedProduct.Count

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func getFavoriteProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	var favoriteProducts = []Product{}

	for _, Product := range products {
		if Product.Favorite == true {
			w.Header().Set("Content-Type", "application/json")
			favoriteProducts = append(favoriteProducts, Product)
		}
	}

	json.NewEncoder(w).Encode(favoriteProducts)
}

func getShopCartProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	var shopCartProducts = []Product{}

	for _, Product := range products {
		if Product.ShopCart == true {
			w.Header().Set("Content-Type", "application/json")
			shopCartProducts = append(shopCartProducts, Product)
		}
	}

	json.NewEncoder(w).Encode(shopCartProducts)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Получаем ID из URL
	idStr := r.URL.Path[len("/users/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Декодируем обновлённые данные продукта
	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Ищем продукт для обновления
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Image = updatedUser.Image
			users[i].Phone = updatedUser.Phone
			users[i].Mail = updatedUser.Mail

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)           // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)  // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)       // Получить продукт по ID
	http.HandleFunc("/products/update/", updateProductHandler) // Обновить продукт
	http.HandleFunc("/products/delete/", deleteProductHandler) // Удалить продукт

	http.HandleFunc("/favorite_products", getFavoriteProductsHandler)
	http.HandleFunc("/shop_cart_products", getShopCartProductsHandler)

	http.HandleFunc("/users/", getUserByIDHandler)
	http.HandleFunc("/users/update/", updateUserHandler)

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
