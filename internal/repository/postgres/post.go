/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package postgres

import (
	"context"
	"fmt"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/pkg/database/postgres"

	"github.com/gofrs/uuid"
)

// Post table name.
const PostTable string = "post"

// Post repository interface.
type Post interface {
	Create(ctx context.Context, authorID uuid.UUID, text string) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, text string) error
}

// Post repository structure.
type PostRepository struct{ psql postgres.Postgres }

// Creating a new post repository.
func NewPostRepository(psql postgres.Postgres) *PostRepository {
	return &PostRepository{psql: psql}
}

// Creating a new post in postgres database.
func (r *PostRepository) Create(ctx context.Context, authorID uuid.UUID, text string) (uuid.UUID, error) {
	var id uuid.UUID

	// Query to create post.
	query := fmt.Sprintf(`INSERT INTO "%s" (author_id, text) VALUES ($1, $2) RETURNING "id"`, PostTable)

	// Scan post id.
	row := r.psql.QueryRow(ctx, query, authorID, text)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// Getting a post by id in postgres database.
func (r *PostRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	var post domain.Post

	post.ID = id

	// Query for get post by id.
	query := fmt.Sprintf(`SELECT "author_id", "text", "created_at", "updated_at" FROM "%s" WHERE "id"=$1`, PostTable)

	row := r.psql.QueryRow(ctx, query, id)

	// Scanning query row.
	err := row.Scan(&post.AuthorID, &post.Text, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

// Deleting a post in postgres database.
func (r *PostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// Query for delete post by id.
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id=$1`, PostTable)
	_, err := r.psql.Exec(ctx, query, id)

	return err
}

// Updating a post in postgres database.
func (r *PostRepository) Update(ctx context.Context, id uuid.UUID, text string) error {
	// Query for update post by id.
	query := fmt.Sprintf(`UPDATE "%s" SET "text"=$1 WHERE "id"=$2`, PostTable)
	_, err := r.psql.Exec(ctx, query, text, id)

	return err
}