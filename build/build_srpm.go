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
	"strings"
)

func buildsrpm() {
	specDir := utils.TmpDir + "/" + "SPEC"
	_ = filepath.Walk(utils.ShareDir, copyFiles)
	_ = filepath.Walk(specDir, mockBuildSRPM)
}

func mockBuildSRPM(filePath string, f os.FileInfo, err error) error {
	if filepath.Ext(f.Name()) == ".spec" {
		specDir := path.Dir(filePath)
		specFile := specDir + "/" + f.Name()
		name := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		utils.ColorPrint("cyan", "Start downloading "+name+" source files")
		if _, errDown := sh.Command("spectool", "-g", specFile, "-C", utils.TmpDir+"/"+"source").Output(); errDown != nil {
			utils.ColorPrint("red", "Fail to download "+name+" source file")
			os.Exit(1)
		}
		utils.ColorPrint("green", "Download "+name+" source succeeded.")
		utils.ColorPrint("cyan", "Start building "+name+" SRPM ("+utils.MockConfig+").")
		if _, errBuild := sh.Command("/usr/bin/mock",
			"-r",
			utils.MockConfig,
			"--resultdir",
			utils.TmpDir+"/"+"SRPM",
			"--buildsrpm",
			"--sources",
			utils.TmpDir+"/"+"source",
			"--spec",
			specFile).Output(); errBuild != nil {
			utils.ColorPrint("red", "Build "+name+" SRPM failed ("+utils.MockConfig+").")
			os.Exit(1)
		}
		utils.ColorPrint("green", "Build "+name+" SRPM succeeded ("+utils.MockConfig+").")
	}
	return nil
}

func copyFiles(filePath string, f os.FileInfo, err error) error {
	src := path.Dir(filePath) + "/" + f.Name()
	dst := utils.TmpDir + "/source/" + f.Name()
	if f.IsDir() {
		return nil
	}
	utils.CopyFile(src, dst)
	return nil
}
