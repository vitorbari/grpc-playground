package main

import (
	"context"
	"fmt"

	"github.com/vitorbari/grpc-playground/13-crud-api-mongodb-golang/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type server struct{}

func (s *server) CreateBlog(ctx context.Context, r *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := r.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprint("Cannot get inserted ID"),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (s *server) ReadBlog(ctx context.Context, r *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	id := r.GetBlogId()

	blog := &blogItem{}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid blog ID: %s", id),
		)
	}

	filter := bson.D{{"_id", oid}}
	err = collection.FindOne(ctx, filter).Decode(blog)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Blog not found: %s", id),
		)
	} else if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       blog.ID.Hex(),
			AuthorId: blog.AuthorID,
			Title:    blog.Title,
			Content:  blog.Content,
		},
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, r *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	blog := r.GetBlog()

	blogItem := &blogItem{}

	oid, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid blog ID: %s", blog.GetId()),
		)
	}

	filter := bson.D{{"_id", oid}}
	err = collection.FindOne(ctx, filter).Decode(blogItem)
	if err == mongo.ErrNoDocuments {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Blog not found: %s", blog.GetId()),
		)
	} else if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	blogItem.AuthorID = blog.GetAuthorId()
	blogItem.Title = blog.GetTitle()
	blogItem.Content = blog.GetContent()

	result, err := collection.ReplaceOne(ctx, filter, blogItem)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", err),
		)
	}

	if result.ModifiedCount != 1 {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("UpdateRequest modified %d records", result.ModifiedCount),
		)
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       blogItem.ID.Hex(),
			AuthorId: blogItem.AuthorID,
			Title:    blogItem.Title,
			Content:  blogItem.Content,
		},
	}, nil
}

func (s *server) DeleteBlog(ctx context.Context, r *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	oid, err := primitive.ObjectIDFromHex(r.GetBlogId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid blog ID: %s", r.GetBlogId()),
		)
	}

	filter := bson.D{{"_id", oid}}
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog in MongoDB: %v", err),
		)
	}

	return &blogpb.DeleteBlogResponse{BlogId: r.GetBlogId()}, nil
}

func (s *server) ListBlogs(r *blogpb.ListBlogsRequest, stream blogpb.BlogService_ListBlogsServer) error {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		blogItem := &blogItem{}
		err := cursor.Decode(&blogItem)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Failed to decode blog item: %v", err),
			)
		}

		stream.Send(&blogpb.ListBlogsResponse{
			Blog: &blogpb.Blog{
				Id:       blogItem.ID.Hex(),
				AuthorId: blogItem.AuthorID,
				Title:    blogItem.Title,
				Content:  blogItem.Content,
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return nil
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}
