# Mock-Travis
[![GitHub license](https://img.shields.io/badge/license-GPL%20V3.0-red.svg?style=flat-square)](https://raw.githubusercontent.com/YKMeIz/bspwm-config/main/LICENSE)
[![Build Status](http://img.shields.io/travis/YKMeIz/mock-travis.svg?style=flat-square)](https://travis-ci.org/YKMeIz/mock-travis)
[![Go Report Card](https://goreportcard.com/badge/github.com/YKMeIz/mock-travis)](https://goreportcard.com/report/github.com/YKMeIz/mock-travis)
[![Gitter](https://badges.gitter.im/YKMeIz/mock-travis.svg)](https://gitter.im/YKMeIz/mock-travis?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Mock-Travis utilizes [Travis CI](https://travis-ci.org/) to provide a continuous integration practice for testing **`spec`** files of RedHat RPM packages. That is to say, Mock-Travis can be utilized to build source rpm packages and binary rpm packages in order to check whether **`spec`** files are written correctly.

> **Note**: RPM package **spec** files need to be stored in the GitHub repository.


### How to use Mock-Travis?
Setting up Mock-Travis is quite simple. All you need to do are just two things:
- Put **`.travis.yml`** your GitHub repository.
- Enable autobuild on [Travis CI](https://travis-ci.org/) website.

> Assumption: You should know how to add file and push to your GitHub repository; how to [sign in to Travis CI](https://travis-ci.org/auth) with your GitHub account, and go to [profile page](https://travis-ci.org/profile) and enable [Travis CI](https://travis-ci.org/) for the repository you want to build.


**Example of `.travis.yml` file**
```yml
#! -------------------------------------------------------------------------
#! Mock-Travis Settings
#! -------------------------------------------------------------------------

mock_travis:
  # Set a particular test configuration for mock.
  mock_config: fedora-25-x86_64

  # Use RPM spec files from a GitHub repository
  # to generate extra part of buildrequires packages.
  # For example
  #         If your spec repository is 
  #         "https://github.com/YKMeIz/Sway-Fedora",
  #         git could be set to `YKMeIz/Sway-Fedora`
  packages_buildrequires_git: 
  # Use extra/external repository during building packages.
  # This option allows mock to access an additional repository
  # plus defaults repositories based on the mock config.
  # gpgcheck is disabled in this option.
  # For example
  #         If add FZUG as an extra repository
  #         packages_extra_repo should set to
  #         "https://repo.fdzh.org/FZUG/testing/24/x86_64/"
  packages_extra_repo:


#! -------------------------------------------------------------------------
#! DO NOT EDIT THE FOLLOWING SETTINGS
#! UNLESS YOU KNOW WHAT YOU ARE DOING
#! -------------------------------------------------------------------------

sudo: required

services:
  - docker

script:
  - wget -q https://github.com/YKMeIz/mock-travis/releases/download/latest/mock-travis 
        && chmod +x mock-travis 
        && ./mock-travis
```

You can simply copy and paste the example above, or download the [example.travis.yml](https://raw.githubusercontent.com/YKMeIz/mock-travis/main/example.travis.yml).

> **Note**: Please remember to rename **[example.travis.yml](https://raw.githubusercontent.com/YKMeIz/mock-travis/main/example.travis.yml)** to **.travis.yml** if you download **[example.travis.yml](https://raw.githubusercontent.com/YKMeIz/mock-travis/main/example.travis.yml)**.


### How Mock-Travis works?
When you make a push to your GitHub repository, it will trigger a [Travis CI](https://travis-ci.org/) build. The build process will run a docker container and do the following things:
- Initialize mock config.
- Build source packages.
- Build binary packages, and record packages build failed.
- Create local repository and add it to the mock config.
- Rebuild binary packages based on the failed records.


### Advantages
- No need to test mock build on your own computer. It is quite hard to run mock if you use other GNU/Linux distros than RedHat related GNU/Linux distros.
- No need to worry about build requires. Mock-travis gives build faild packages another try with local repository which contains packages just built. It should be sufficient to solve the missing buildrequires issue.
- Beautiful output. Mock-travis generates colorful bold information to exhibit the build process and status. The [Travis CI](https://travis-ci.org/) will show the whole build log. The results of each step can be found easily in build log as they will be shown in colored bold words.

Here is an example of build log:
![Travis-CI log](https://github.com/YKMeIz/mock-travis/raw/main/misc/travis-ci-log.png)

> Click the picture to view the raw file.


### Projects use Mock-Travis
[YKMeIz/Sway-Fedora](https://github.com/YKMeIz/Sway-Fedora) is a real world example of utilizing Mock-Travis. You can check its [.travis.yml](https://github.com/YKMeIz/Sway-Fedora/blob/main/.travis.yml) file, or [Travis CI build log](https://travis-ci.org/YKMeIz/Sway-Fedora).



### Limitations
- The customizability is still low. Only provide a few options currently.
- Run `mock` in docker container to test packages is few minutes slower than test locally.


### ToDo
- Basically, it is to make mock-travis more customizable and run faster.


### Contributing
- If you have any suggestion, idea, or bug report; feel free to open an issue on this repository.
