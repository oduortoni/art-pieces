package types

type Piece struct {
	Identifier int	`json:"identifier"`
	Title string	`json:"title"`
	Slug string		`json:"slug"`
	Value float64	`json:"value"`
	Description string	`json:"description"`
	Details string	`json:"details"`
}
