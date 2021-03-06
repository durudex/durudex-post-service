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

package service

import (
	"context"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/repository/postgres"
	"github.com/segmentio/ksuid"
)

// Post interface.
type Post interface {
	Create(ctx context.Context, post domain.Post) (ksuid.KSUID, error)
	GetByID(ctx context.Context, id ksuid.KSUID) (domain.Post, error)
	Delete(ctx context.Context, id, authorID ksuid.KSUID) error
	Update(ctx context.Context, post domain.Post) error
}

// Post service structure.
type PostService struct{ repos postgres.Post }

// Creating a new post service.
func NewPostService(repos postgres.Post) *PostService {
	return &PostService{repos: repos}
}

// Creating a new post.
func (s *PostService) Create(ctx context.Context, post domain.Post) (ksuid.KSUID, error) {
	// Validate a post.
	if err := post.Validate(); err != nil {
		return ksuid.Nil, err
	}

	// Create a new post.
	id, err := s.repos.Create(ctx, post)
	if err != nil {
		return ksuid.Nil, err
	}

	return id, nil
}

// Getting a post by id.
func (s *PostService) GetByID(ctx context.Context, id ksuid.KSUID) (domain.Post, error) {
	// Get post by id.
	post, err := s.repos.GetByID(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

// Deleting a post.
func (s *PostService) Delete(ctx context.Context, id, authorID ksuid.KSUID) error {
	return s.repos.Delete(ctx, id, authorID)
}

// Updating a post.
func (s *PostService) Update(ctx context.Context, post domain.Post) error {
	// Validate a post.
	if err := post.Validate(); err != nil {
		return err
	}

	return s.repos.Update(ctx, post)
}
