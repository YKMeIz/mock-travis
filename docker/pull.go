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

package docker

import (
	"github.com/codeskyblue/go-sh"
	"github.com/nrechn/mock-travis/utils"
	"os"
)

func pull() {
	utils.ColorPrint("cyan", "Start pulling "+utils.DockerImage+" docker image...")
	if _, err := sh.Command("docker", "pull", utils.DockerImage).Output(); err != nil {
		utils.ColorPrint("red", "Pull "+utils.DockerImage+" docker image failed.")
		os.Exit(1)
	}
	utils.ColorPrint("green", "Pull "+utils.DockerImage+" docker image succeeded.")
}
