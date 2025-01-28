package models

type PaginationQuery struct {
    Page     int    `json:"page"`
    PageSize int    `json:"page_size"`
    Search   string `json:"search"`
    SortBy   string `json:"sort_by"`
    SortDir  string `json:"sort_dir"`
    MinPrice float64 `json:"min_price"`
    MaxPrice float64 `json:"max_price"`
    Author   string `json:"author"`
}

type PaginatedResponse struct {
    Data       interface{} `json:"data"`
    Total      int        `json:"total"`
    Page       int        `json:"page"`
    PageSize   int        `json:"page_size"`
    TotalPages int        `json:"total_pages"`
}