/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/pkg/database/postgres"
	"github.com/segmentio/ksuid"

	"github.com/jackc/pgx/v4"
)

// Post table name.
const PostTable string = "post"

// Post repository interface.
type Post interface {
	Create(ctx context.Context, post domain.Post) (ksuid.KSUID, error)
	GetByID(ctx context.Context, id ksuid.KSUID) (domain.Post, error)
	Delete(ctx context.Context, id, authorID ksuid.KSUID) error
	Update(ctx context.Context, post domain.Post) error
}

// Post repository structure.
type PostRepository struct{ psql postgres.Postgres }

// Creating a new post repository.
func NewPostRepository(psql postgres.Postgres) *PostRepository {
	return &PostRepository{psql: psql}
}

// Creating a new post in postgres database.
func (r *PostRepository) Create(ctx context.Context, post domain.Post) (ksuid.KSUID, error) {
	var id string

	// Query to create post.
	query := fmt.Sprintf(`INSERT INTO "%s" (id, author_id, text) VALUES ($1, $2, $3) RETURNING "id"`, PostTable)

	// Scan post id.
	row := r.psql.QueryRow(ctx, query, post.Id.String(), post.AuthorId.String(), post.Text)
	if err := row.Scan(&id); err != nil {
		return ksuid.Nil, err
	}

	return ksuid.Parse(id)
}

// Getting a post by id in postgres database.
func (r *PostRepository) GetByID(ctx context.Context, id ksuid.KSUID) (domain.Post, error) {
	var post domain.Post

	// Query for get post by id.
	query := fmt.Sprintf(`SELECT "author_id", "text", "updated_at" FROM "%s" WHERE "id"=$1`, PostTable)

	row := r.psql.QueryRow(ctx, query, id.String())

	// Scanning query row.
	if err := row.Scan(&post.AuthorId, &post.Text, &post.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Post{}, &domain.Error{Code: domain.CodeNotFound, Message: "User not found"}
		}

		return domain.Post{}, &domain.Error{Code: domain.CodeInternal, Message: "Internal Server Error"}
	}

	return post, nil
}

// Deleting a post in postgres database.
func (r *PostRepository) Delete(ctx context.Context, id, authorID ksuid.KSUID) error {
	// Query for delete post by id.
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id=$1 AND author_id=$2`, PostTable)
	_, err := r.psql.Exec(ctx, query, id.String(), authorID.String())

	return err
}

// Updating a post in postgres database.
func (r *PostRepository) Update(ctx context.Context, post domain.Post) error {
	// Query for update post by id.
	query := fmt.Sprintf(`UPDATE "%s" SET text=$1, updated_at=now() WHERE "id"=$2 AND author_id=$3`, PostTable)
	_, err := r.psql.Exec(ctx, query, post.Text, post.Id.String(), post.AuthorId.String())

	return err
}
