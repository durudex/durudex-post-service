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

package config

const (
	// Config defaults.
	defaultConfigPath string = "configs/main"

	// Server defaults.
	defaultServerHost string = "post.service.durudex.local"
	defaultServerPort string = "8005"

	// TLS server defaults.
	defaultTLSEnable bool   = true
	defaultTLSCACert string = "./certs/rootCA.pem"
	defaultTLSCert   string = "./certs/post.service.durudex.local-cert.pem"
	defaultTLSKey    string = "./certs/post.service.durudex.local-key.pem"
)
