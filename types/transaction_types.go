package types

type ArticleOwnership struct {
    ArticleID string  `json:"articleId"`
    OwnerID   string  `json:"ownerId"`
    Price     float64 `json:"price"`
}
