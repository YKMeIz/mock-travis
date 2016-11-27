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

func Init() {
	pull()
	mockCfg := utils.GetYmlSlice("mock_travis.mock_config")
	for i := 0; i < len(mockCfg); i++ {
		run(mockCfg[i])
		clean()
	}
}

func run(cfg string) {
	volLoc := utils.CurrLoc() + "/:" + utils.ShareDir
	if err := sh.Command("docker",
		"run",
		"--name",
		utils.ContainerName,
		"--cap-add=SYS_ADMIN",
		"--privileged=true",
		"-v",
		volLoc,
		"-i",
		utils.DockerImage,
		utils.ShareDir+"/mock-travis",
		cfg).Run(); err != nil {
		utils.ColorPrint("red",
			"OVERALL: Fail to build "+
				utils.GetYml("mock_travis.packages_name")+
				" and related build dependencies.")
		os.Exit(1)
	}
	utils.ColorPrint("yellow",
		"OVERALL: Successfully build RPMs and related build dependencies ("+cfg+").")
}
