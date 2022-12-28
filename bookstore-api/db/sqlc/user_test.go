package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alekseiapa/mini-go-projects/book-store/util"
)

func createRandomBook(t *testing.T) Book {
	arg := CreateBookParams{
		Name:        util.RandomBookName(),
		Publication: util.RandomBookPublication(),
	}
	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)
	require.Equal(t, arg.Name, book.Name)
	require.Equal(t, arg.Publication, book.Publication)
	require.NotZero(t, book.Uuid)
	return book
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

func TestGetBook(t *testing.T) {
	book1 := createRandomBook(t)
	book2, err := testQueries.GetBook(context.Background(), book1.Uuid)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.Uuid, book2.Uuid)
	require.Equal(t, book1.Publication, book2.Publication)
	require.Equal(t, book1.Name, book2.Name)
}

func TestUpdateBookName(t *testing.T) {
	book1 := createRandomBook(t)
	arg := UpdateBookNameParams{Uuid: book1.Uuid, Name: book1.Name}
	book2, err := testQueries.UpdateBookName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book2)
	require.Equal(t, book1.Name, book2.Name)
	require.Equal(t, book1.Uuid, book2.Uuid)
	require.Equal(t, book1.Publication, book2.Publication)
}

func TestDeleteBook(t *testing.T) {
	book1 := createRandomBook(t)
	err := testQueries.DeleteBook(context.Background(), book1.Uuid)
	require.NoError(t, err)

	book2, err := testQueries.GetBook(context.Background(), book1.Uuid)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book2)
}

func TestListBooks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBook(t)
	}
	arg := ListBooksParams{
		Limit:  10,
		Offset: 10,
	}

	books, err := testQueries.ListBooks(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, books, 10)

	for _, book := range books {
		require.NotEmpty(t, book)
	}
}
