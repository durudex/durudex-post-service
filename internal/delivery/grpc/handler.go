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

package grpc

import (
	"github.com/durudex/durudex-post-service/internal/service"

	"google.golang.org/grpc"
)

// GRPC handler structure.
type Handler struct{ service *service.Service }

// Creating a new gRPC handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// Registration services handlers.
func (h *Handler) RegisterHandlers(srv *grpc.Server) {}