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
	"log"
	"os"
	"strings"
)

func Init() {
	setTmpDir()
	if repo := utils.GetYml("mock_travis.packages_extra_repo"); repo != "" {
		setRepo(repo)
	}
	if git := utils.GetYml("mock_travis.packages_buildrequires_git"); git != "" {
		setGit(git)
	}
	mock()
}

func setTmpDir() {
	utils.MkDir(utils.TmpDir)
	if err := sh.Command("cp", "-r", utils.ShareDir, utils.TmpDir+"/"+"SPEC").Run(); err != nil {
		utils.ColorPrint("red", "Fail to setup sources to temporary directory.")
		os.Exit(1)
	}
	utils.MkDir(utils.TmpDir + "/" + "SRPM")
	utils.MkDir(utils.TmpDir + "/" + "RPM")
	utils.MkDir(utils.TmpDir + "/" + "debugInfo")
	utils.MkDir(utils.TmpDir + "/" + "source")
}

func setRepo(repo string) {
	utils.ColorPrint("cyan", "Start setting extra package repository")
	repoInfo := `
[extra-local]
name=extra-local
baseurl=` + repo + `
gpgcheck=0
"""
`
	mockCfg := "/etc/mock/" + utils.MockConfig + ".cfg"
	input, err := ioutil.ReadFile(mockCfg)
	if err != nil {
		log.Fatalln(err)
		utils.ColorPrint("red", "Setting extra package repository failed.")
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if line == `"""` {
			lines[i] = repoInfo
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(mockCfg, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		utils.ColorPrint("red", "Setting extra package repository failed.")
		os.Exit(1)
	}
	utils.ColorPrint("green", "Setting extra package repository succeeded.")
}

func setGit(git string) {
	utils.ColorPrint("cyan", "Start setting git repository")
	gitUrl := "https://github.com/" + git
	_, _ = sh.Command("dnf", "-y", "install", "git").Output()
	if _, err := sh.Command("git", "clone", gitUrl, utils.TmpDir+"/"+"SPEC/GIT").Output(); err != nil {
		utils.ColorPrint("red", "Setting git repository failed.")
		os.Exit(1)
	}
	utils.ColorPrint("green", "Setting git repository succeeded.")
}

func mockInit() {
	utils.ColorPrint("cyan", "Start setting up "+utils.MockConfig+" mock environment")
	if _, err := sh.Command("/usr/bin/mock",
		"-r",
		utils.MockConfig,
		"--init").Output(); err != nil {
		utils.ColorPrint("red", "Setup "+utils.MockConfig+" mock environment failed.")
	}
	utils.ColorPrint("green", "Setup "+utils.MockConfig+" mock environment succeeded.")
}
