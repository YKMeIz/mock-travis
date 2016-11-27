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

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func IsContainer() bool {
	// .dockerinit is removed in docker v1.11
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

func GetYml(arg string) string {
	viper.SetConfigName(".travis")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/home/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
		os.Exit(1)
	}
	return viper.GetString(arg)
}

func GetYmlSlice(arg string) []string {
	viper.SetConfigName(".travis")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/home/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
		os.Exit(1)
	}
	return viper.GetStringSlice(arg)
}

func ColorPrint(colorOption, msg string) {
	switch colorOption {
	case "red":
		fmt.Println("\033[31m\033[1m" + msg + "\033[0m\033[39m")

	case "green":
		fmt.Println("\033[32m\033[1m" + msg + "\033[0m\033[39m")

	case "yellow":
		fmt.Println("\033[33m\033[1m" + msg + "\033[0m\033[39m")

	case "cyan":
		fmt.Println("\033[36m\033[1m" + msg + "\033[0m\033[39m")

	default:
		fmt.Println(msg)
	}
}

func CurrLoc() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	return dir
}

func MkDir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
		os.Exit(1)
	}
}
