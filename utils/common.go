// Copyright © 2016 nrechn <nrechn@gmail.com>
//
// This file is part of mock-travis.
//
// mock-travis is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// mock-travis is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with mock-travis. If not, see <http://www.gnu.org/licenses/>.
//

package utils

const (
	ContainerName = "mock-build"
	DockerImage   = "nrechn/fedora-mock"
	ShareDir      = "/home"
	TmpDir        = "/var/tmp/mock-travis"
	LocalRepo     = `
[mock-local]
name=mock-local
baseurl=file:///var/tmp/mock-travis/RPM/
gpgcheck=0
"""
`
)

var (
	RebuildList []string
	StillFail   []string
	MockConfig  string
)
