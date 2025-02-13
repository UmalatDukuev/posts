package graph

import (
	"context"
	"testing"
	"time"

	"posts/graph/model"
	"posts/inmemory"

	"github.com/stretchr/testify/assert"
)

func setupTestResolver() *Resolver {
	return &Resolver{InMemoryStorage: inmemory.NewInMemoryDB()}
}

func TestCreatePost(t *testing.T) {
	resolver := setupTestResolver()
	ctx := context.Background()

	input := model.NewPost{
		Title:       "Test Post",
		Content:     "This is a test post",
		Author:      strPtr("John Doe"),
		PublishedAt: strPtr(time.Now().Format(time.RFC3339)),
	}

	post, err := resolver.Mutation().CreatePost(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Test Post", post.Title)
	assert.Equal(t, "This is a test post", post.Content)
	assert.Equal(t, "John Doe", post.Author)
}

func TestUpdatePost(t *testing.T) {
	resolver := setupTestResolver()
	ctx := context.Background()

	input := model.NewPost{
		Title:   "Old Title",
		Content: "Old Content",
		Author:  strPtr("John Doe"),
	}
	post, _ := resolver.Mutation().CreatePost(ctx, input)

	updateInput := model.NewPost{
		Title:   "New Title",
		Content: "Updated Content",
	}
	updatedPost, err := resolver.Mutation().UpdatePost(ctx, post.ID, &updateInput)
	assert.NoError(t, err)
	assert.NotNil(t, updatedPost)
	assert.Equal(t, "New Title", updatedPost.Title)
	assert.Equal(t, "Updated Content", updatedPost.Content)
}

func TestCreateComment(t *testing.T) {
	resolver := setupTestResolver()
	ctx := context.Background()

	post, _ := resolver.Mutation().CreatePost(ctx, model.NewPost{
		Title:   "Test Post",
		Content: "This is a test post",
		Author:  strPtr("Alice"),
	})

	commentInput := model.NewComment{
		PostID:  post.ID,
		Author:  "Bob",
		Content: "Nice post!",
	}
	comment, err := resolver.Mutation().CreateComment(ctx, commentInput)

	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, "Bob", comment.Author)
	assert.Equal(t, "Nice post!", comment.Content)
}

func TestCreateCommentWhenCommentsDisabled(t *testing.T) {
	resolver := setupTestResolver()
	ctx := context.Background()

	post, _ := resolver.Mutation().CreatePost(ctx, model.NewPost{
		Title:           "No Comments Post",
		Content:         "Comments are disabled",
		Author:          strPtr("Admin"),
		PublishedAt:     strPtr(time.Now().Format(time.RFC3339)),
		CommentsAllowed: false,
	})

	commentInput := model.NewComment{
		PostID:  post.ID,
		Author:  "Bob",
		Content: "Trying to comment...",
	}
	comment, err := resolver.Mutation().CreateComment(ctx, commentInput)

	assert.Error(t, err)
	assert.Nil(t, comment)
}

func TestGetAllPosts(t *testing.T) {
	resolver := setupTestResolver()
	ctx := context.Background()

	resolver.Mutation().CreatePost(ctx, model.NewPost{Title: "Post 1", Content: "Content 1", Author: strPtr("Alice")})
	resolver.Mutation().CreatePost(ctx, model.NewPost{Title: "Post 2", Content: "Content 2", Author: strPtr("Bob")})

	posts, err := resolver.Query().GetAllPosts(ctx)
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
}

func TestGetOnePost(t *testing.T) {
	resolver := setupTestResolver()
	ctx := context.Background()

	post, _ := resolver.Mutation().CreatePost(ctx, model.NewPost{
		Title:   "Unique Post",
		Content: "This is a unique post",
		Author:  strPtr("Charlie"),
	})

	foundPost, err := resolver.Query().GetOnePost(ctx, post.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundPost)
	assert.Equal(t, "Unique Post", foundPost.Title)
	assert.Equal(t, "This is a unique post", foundPost.Content)
}

func strPtr(s string) *string {
	return &s
}
