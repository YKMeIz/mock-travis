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
	"github.com/nrechn/mock-travis/utils"
	"os"
)

func mock() {
	mockInit()
	buildsrpm()
	buildrpm()
	if utils.RebuildList != nil {
		updateRepo()
		rebuildRPM()
		if len(utils.StillFail) != 0 {
			utils.ColorPrint("red", "Still build failed packages:")
			for i := 0; i < len(utils.StillFail); i++ {
				utils.ColorPrint("red", utils.StillFail[i])
			}
			os.Exit(1)
		}
	}
}
