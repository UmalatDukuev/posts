// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID              string     `json:"id"`
	Content         string     `json:"content"`
	Author          *User      `json:"author"`
	PostID          string     `json:"postId"`
	ParentCommentID *string    `json:"parentCommentId,omitempty"`
	CreatedAt       string     `json:"createdAt"`
	Replies         []*Comment `json:"replies,omitempty"`
}

type CommentStatus struct {
	IsError bool     `json:"isError"`
	Message *string  `json:"message,omitempty"`
	Comment *Comment `json:"comment,omitempty"`
}

type CreateCommentInput struct {
	Content         string  `json:"content"`
	AuthorID        string  `json:"authorId"`
	PostID          string  `json:"postId"`
	ParentCommentID *string `json:"parentCommentId,omitempty"`
}

type CreatePostInput struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	AuthorID      string `json:"authorId"`
	AllowComments bool   `json:"allowComments"`
}

type DeleteStatus struct {
	IsError bool    `json:"isError"`
	Message *string `json:"message,omitempty"`
}

type Mutation struct {
}

type Post struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	AuthorID      string     `json:"authorId"`
	AllowComments bool       `json:"allowComments"`
	Comments      []*Comment `json:"comments,omitempty"`
}

type PostStatus struct {
	IsError bool    `json:"isError"`
	Message *string `json:"message,omitempty"`
	Post    *Post   `json:"post,omitempty"`
}

type Query struct {
}

type Subscription struct {
}

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    *string `json:"email,omitempty"`
}
