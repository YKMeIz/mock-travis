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
	"io"
	"os"
	"path/filepath"
)

// Check if program is running inside docker container.
func IsContainer() bool {
	// .dockerinit is removed in docker v1.11.
	// Instead, using .dockerenv to determine if
	// program is running inside docker container.
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

// Read .travis.yml and retrieve string value given the arg to use.
func GetYml(arg string) string {
	viper.SetConfigName(".travis")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/home/")
	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}
	return viper.GetString(arg)
}

// Read .travis.yml and retrieve slice value given the arg to use.
func GetYmlSlice(arg string) []string {
	viper.SetConfigName(".travis")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/home/")
	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}
	return viper.GetStringSlice(arg)
}

// Output in color.
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

// Get current location.
func CurrLoc() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		os.Exit(1)
	}
	return dir
}

// Same as `mkdir` command.
func MkDir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		os.Exit(1)
	}
}

// Copy file.
// Note: Destination should contain destination file name.
func CopyFile(src, dst string) {
	srcFile, err := os.Open(src)
	defer srcFile.Close()
	check(err)

	destFile, err := os.Create(dst)
	defer destFile.Close()
	check(err)

	_, err = io.Copy(destFile, srcFile)
	check(err)

	err = destFile.Sync()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}
