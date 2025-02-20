package dto

// DishChoice는 요리 선택지 입력 DTO
type DishChoice struct {
	Name  string `json:"name"`
	Extra int    `json:"extra,omitempty"`
}

// DishOption는 요리 옵션 입력 DTO
type DishOption struct {
	Name    string       `json:"name"`
	Choices []DishChoice `json:"choices,omitempty"`
	Extra   int          `json:"extra,omitempty"`
}

// CreateDish는 요리 생성 입력 DTO
type CreateDish struct {
	Name         string       `json:"name"`
	Price        int          `json:"price"`
	Photo        string       `json:"photo,omitempty"`
	Description  string       `json:"description"`
	RestaurantID int          `json:"restaurant_id"`
	Options      []DishOption `json:"options,omitempty"`
}

// UpdateDish는 요리 수정 입력 DTO
type UpdateDish struct {
	Name        string       `json:"name,omitempty"`
	Price       int          `json:"price,omitempty"`
	Photo       string       `json:"photo,omitempty"`
	Description string       `json:"description,omitempty"`
	Options     []DishOption `json:"options,omitempty"`
}

// DeleteDish는 요리 삭제 입력 DTO
type DeleteDish struct {
	DishID int `json:"dish_id"`
}

// DishInput는 특정 요리 조회 입력 DTO
type DishInput struct {
	DishID int `json:"dish_id"`
}

// DishsInput는 요리 목록 조회 입력 DTO
type DishsInput struct {
	Page int `json:"page"`
}

// SearchDish는 요리 검색 입력 DTO
type SearchDish struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
}
