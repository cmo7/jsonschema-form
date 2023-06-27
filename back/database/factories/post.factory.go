package factories

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"
)

type PostFactory struct{}

func (PostFactory) CreateOne() *models.Post {

	user := UserFactory{}.CreateOne()
	post := models.Post{
		Title:   Faker.Lorem().Sentence(7),
		Content: Faker.Lorem().Paragraph(5),
		User:    *user,
	}
	return &post
}

func (PostFactory) CreateMany(count int) []*models.Post {
	var posts []*models.Post
	for i := 0; i < count; i++ {
		post := PostFactory{}.CreateOne()
		posts = append(posts, post)
	}
	return posts
}

func (PostFactory) CreateOneWithData(data *models.Post) *models.Post {
	post := models.Post{}
	if data.Title != "" {
		post.Title = data.Title
	} else {
		post.Title = Faker.Lorem().Sentence(7)
	}

	if data.Content != "" {
		post.Content = data.Content
	} else {
		post.Content = Faker.Lorem().Paragraph(5)
	}

	if data.User.GetId().String() != "" {
		post.User = data.User
	} else {
		user := UserFactory{}.CreateOne()
		database.DB.Create(&user)
		post.User = *user
	}

	return &post
}

func (PostFactory) CreateManyWithData(count int, data *models.Post) []*models.Post {
	var posts []*models.Post
	post := &models.Post{}
	for i := 0; i < count; i++ {
		post = PostFactory{}.CreateOneWithData(data)
		posts = append(posts, post)
	}
	return posts
}
