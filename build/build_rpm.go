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
	"os"
	"path"
	"path/filepath"
)

func buildrpm() {
	srpmDir := utils.TmpDir + "/" + "SRPM"
	_ = filepath.Walk(srpmDir, mockBuildRPM)
}

func mockBuildRPM(filePath string, f os.FileInfo, err error) error {
	if filepath.Ext(f.Name()) == ".rpm" {
		srpmDir := path.Dir(filePath)
		srpmFile := srpmDir + "/" + f.Name()
		name := f.Name()[0 : len(f.Name())-13]
		utils.ColorPrint("cyan", "Start building "+name+" binary RPM ("+utils.MockConfig+").")
		_, err = sh.Command("/usr/bin/mock",
			"-r",
			utils.MockConfig,
			"--resultdir",
			utils.TmpDir+"/"+"RPM",
			"--rebuild",
			srpmFile,
		).Output()
		if err != nil {
			utils.ColorPrint("red", "Build "+name+" binary RPM failed ("+utils.MockConfig+").")
			utils.RebuildList = append(utils.RebuildList, srpmFile)
		} else {
			utils.ColorPrint("green", "Build "+name+" binary RPM succeeded ("+utils.MockConfig+").")
		}
	}
	return nil
}

func rebuildRPM() {
	utils.ColorPrint("yellow", "Start rebuilding for the binary RPMs built failed.")
	for i := 0; i < cap(utils.RebuildList); i++ {
		fileFullName := path.Base(utils.RebuildList[i])
		utils.ColorPrint("cyan", "Start rebuild "+fileFullName+" ("+utils.MockConfig+").")
		if _, err := sh.Command("/usr/bin/mock",
			"-r",
			utils.MockConfig,
			"--resultdir",
			utils.TmpDir+"/"+"RPM",
			"--rebuild",
			utils.RebuildList[i],
		).Output(); err != nil {
			utils.ColorPrint("red", "Rebuild "+fileFullName+" failed ("+utils.MockConfig+").")
			utils.StillFail = append(utils.StillFail, fileFullName)
			os.Exit(1)
		}
		utils.ColorPrint("green", "Rebuild "+fileFullName+" succeeded ("+utils.MockConfig+").")
	}
}
