package data

import (
	"github.com/travboz/backend-projects/personal-blog-api/internal/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt CustomDate         `bson:"created_at"`
	Tags      []string           `bson:"tags"`
}

func ValidateArticle(v *validator.Validator, article *Article) {
	// Use the Check() method to execute our validation checks. This will add the
	// provided key and error message to the errors map if the check does not evaluate
	// to true.
	v.Check(article.Content != "", "content", "must be provided")
	v.Check(len(article.Content) <= 1000, "content", "must not be more than 1000 bytes long")

	v.Check(article.Tags != nil, "tags", "must be provided")
	v.Check(len(article.Tags) >= 1, "tags", "must contain at least 1 tag")
	v.Check(len(article.Tags) <= 5, "tags", "must not contain more than 5 tags")
	// Note that we're using the Unique helper in the line below to check that all
	// values in the input.Gensres slice are unique.
	v.Check(validator.Unique(article.Tags), "tags", "must not contain duplicate tags")
}
