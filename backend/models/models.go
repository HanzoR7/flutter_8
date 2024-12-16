package main

type Product struct {
	ID          int
	Name        string
	Image       string
	Cost        int
	Describtion string
}

type User struct {
	ID    int
	Image string
	Name  string
	Phone string
	Mail  string
}

type ShopCartProduct struct {
	ID     int
	Count  int
	UserID int
}

type favoriteProduct struct {
	ID     int
	UserID int
}

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
