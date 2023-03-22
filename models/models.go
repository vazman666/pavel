package models

import (
	"database/sql"
	"sync"
)

type Product struct {
	Id       int     `json:"id"`
	Partnum  string  `json:"partnum`
	Qty      int     `json:"qty"`
	Price    float32 `json:"price"`
	Summ     float32 `json:"summ"`
	Provider string  `json:"provider"`
	Name     string  `json:"name"`
	Remark   string  `json:"remark"`
	Date     string  `json:"date"`
	Color    string  `json:"color"`
}

var Database *sql.DB
var M sync.Mutex
