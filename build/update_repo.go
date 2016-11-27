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

package build

import (
	"github.com/codeskyblue/go-sh"
	"github.com/nrechn/mock-travis/utils"
	"io/ioutil"
	"os"
	"strings"
)

func createRepo() {
	if _, err := sh.Command("createrepo", utils.TmpDir+"/"+"RPM").Output(); err != nil {
		utils.ColorPrint("red", "Update local repository failed")
		os.Exit(1)
	}
}

func updateRepo() {
	utils.ColorPrint("cyan", "Start updating local repository")
	createRepo()
	mockCfg := "/etc/mock/" + utils.MockConfig + ".cfg"
	input, err := ioutil.ReadFile(mockCfg)
	if err != nil {
		utils.ColorPrint("red", "Update local repository failed")
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if line == `"""` {
			lines[i] = utils.LocalRepo
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(mockCfg, []byte(output), 0644)
	if err != nil {
		utils.ColorPrint("red", "Update local repository failed")
		os.Exit(1)
	}
	utils.ColorPrint("green", "Update local repository succeeded")
}
