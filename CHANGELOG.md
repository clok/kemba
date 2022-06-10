<a name="unreleased"></a>
## [Unreleased]


<a name="v1.2.0"></a>
## [v1.2.0] - 2022-06-10
### Features
- support mid tag wildcard for enabling logger ([#31](https://github.com/clok/kemba/issues/31))


<a name="v1.1.2"></a>
## [v1.1.2] - 2022-06-10
### Bug Fixes
- **deps:** update all non-major dependencies ([#30](https://github.com/clok/kemba/issues/30))
- **deps:** update module github.com/stretchr/testify to v1.7.1 ([#27](https://github.com/clok/kemba/issues/27))
- **deps:** update module github.com/gookit/color to v1.4.2 ([#21](https://github.com/clok/kemba/issues/21))

### Chore
- update changelog for v1.1.2
- create FUNDING.yml
- update README
- **deps:** update golangci/golangci-lint-action action to v3 ([#26](https://github.com/clok/kemba/issues/26))
- **deps:** update actions/setup-go action to v3 ([#25](https://github.com/clok/kemba/issues/25))
- **deps:** update actions/checkout action to v3 ([#24](https://github.com/clok/kemba/issues/24))
- **deps:** update all non-major dependencies ([#22](https://github.com/clok/kemba/issues/22))
- **deps:** update goreleaser/goreleaser-action action to v3 ([#28](https://github.com/clok/kemba/issues/28))
- **deps:** update jandelgado/gcov2lcov-action action to v1.0.9 ([#29](https://github.com/clok/kemba/issues/29))


<a name="v1.1.1"></a>
## [v1.1.1] - 2021-04-05
### Chore
- update changelog for v1.1.1

### Ci
- add makefile and go releaser
- **renovate:** fix syntax error in renovate config


<a name="v1.1.0"></a>
## [v1.1.0] - 2021-03-15
### Chore
- **ci:** port to using golangci-lint github action
- **go.mod:** bump to go 1.16

### Features
- **release:** v1.1.0
- **release:** v1.1.0


<a name="v1.0.0"></a>
## [v1.0.0] - 2021-03-03
### Chore
- **deps:** update actions/checkout action to v2 ([#11](https://github.com/clok/kemba/issues/11))
- **deps:** update module gookit/color to v1.3.0 ([#7](https://github.com/clok/kemba/issues/7))
- **deps:** update actions/setup-go action to v2 ([#12](https://github.com/clok/kemba/issues/12))
- **deps:** update module stretchr/testify to v1.7.0 ([#16](https://github.com/clok/kemba/issues/16))
- **deps:** update module gookit/color to v1.3.6 ([#15](https://github.com/clok/kemba/issues/15))
- **deps:** update jandelgado/gcov2lcov-action action to v1.0.8 ([#14](https://github.com/clok/kemba/issues/14))
- **deps:** update module gookit/color to v1.3.2 ([#13](https://github.com/clok/kemba/issues/13))
- **deps:** update module gookit/color to v1.3.1 ([#8](https://github.com/clok/kemba/issues/8))
- **deps:** update coverallsapp/github-action action to v1.1.2 ([#9](https://github.com/clok/kemba/issues/9))
- **deps:** update jandelgado/gcov2lcov-action action to v1.0.7 ([#10](https://github.com/clok/kemba/issues/10))
- **github actions:** add go proxy warming
- **lint:** fix linting error on struct
- **renovate:** add gomodTidy option
- **renovate:** add extension for group:allNonMajor

### Features
- **release:** v1.0.0


<a name="v0.7.1"></a>
## [v0.7.1] - 2020-08-21
### Bug Fixes
- return pointer to color object

### Features
- **release:** v0.7.1


<a name="v0.7.0"></a>
## [v0.7.0] - 2020-08-21
### Features
- export PickColor helper method
- **release:** v0.7.0


<a name="v0.6.4"></a>
## [v0.6.4] - 2020-08-21
### Chore
- **dependencies:** remove duplicate sync of gookit/color
- **deps:** update module kr/pretty to v0.2.1 ([#6](https://github.com/clok/kemba/issues/6))
- **deps:** update module gookit/color to v1.2.7 ([#4](https://github.com/clok/kemba/issues/4))
- **deps:** add renovate.json ([#3](https://github.com/clok/kemba/issues/3))
- **renovate:** update config
- **renovate:** clean up dupe config
- **renovate:** add config file

### Features
- **release:** v0.6.4


<a name="v0.6.3"></a>
## [v0.6.3] - 2020-08-02
### Chore
- update readme with new badges, including awesome-go

### Features
- **release:** v0.6.3


<a name="v0.6.2"></a>
## [v0.6.2] - 2020-07-29
### Chore
- updated docs on exported struct

### Features
- **release:** v0.6.2


<a name="v0.6.1"></a>
## [v0.6.1] - 2020-07-27
### Chore
- do not export base kemba struct

### Features
- **release:** v0.6.1


<a name="v0.6.0"></a>
## [v0.6.0] - 2020-07-27
### Features
- **release:** v0.6.0
- **time delta:** added log event time deltas to first line of logged buffer events. Resolution in ms


<a name="v0.5.0"></a>
## [v0.5.0] - 2020-07-21
### Features
- **env:** support reading both the DEBUG and KEMBA env vars
- **release:** v0.5.0


<a name="v0.4.0"></a>
## [v0.4.0] - 2020-07-21
### Chore
- updated README

### Features
- **color:** deterministic color picking based on tag input for a new logger
- **release:** v0.4.0


<a name="v0.3.0"></a>
## [v0.3.0] - 2020-07-19
### Features
- **extend:** capability to extend existing logger with appended tag values

### Release
- **v0.3.0:** update CHANGELOG


<a name="v0.2.3"></a>
## [v0.2.3] - 2020-07-17
### Chore
- updated README usage notes

### Release
- **v0.2.3:** update CHANGELOG


<a name="v0.2.2"></a>
## [v0.2.2] - 2020-07-17
### Chore
- updated README example code.

### Release
- **v0.2.2:** update CHANGELOG


<a name="v0.2.1"></a>
## [v0.2.1] - 2020-07-17
### Chore
- updates to docs

### Release
- **v0.2.1:** update CHANGELOG


<a name="v0.2.0"></a>
## [v0.2.0] - 2020-07-17
### Bug Fixes
- adjust go.mod package name

### Chore
- updated docs and README
- added github actions & test coverage ([#2](https://github.com/clok/kemba/issues/2))
- update readme with example output
- added changelog

### Release
- **v0.2.0:** update CHANGELOG

### Tech Debt
- reduced color list to more agreeable list of select colors


<a name="v0.1.0"></a>
## v0.1.0 - 2020-07-16
### Chore
- update README
- update README, add contribution guide and initial go.mod

### Features
- **logger:** initial implementation of Printf, Println and Log ([#1](https://github.com/clok/kemba/issues/1))
- **tag:** enable regex tag processing for filtering of log calls

### Initial
- commit


[Unreleased]: https://github.com/clok/kemba/compare/v1.2.0...HEAD
[v1.2.0]: https://github.com/clok/kemba/compare/v1.1.2...v1.2.0
[v1.1.2]: https://github.com/clok/kemba/compare/v1.1.1...v1.1.2
[v1.1.1]: https://github.com/clok/kemba/compare/v1.1.0...v1.1.1
[v1.1.0]: https://github.com/clok/kemba/compare/v1.0.0...v1.1.0
[v1.0.0]: https://github.com/clok/kemba/compare/v0.7.1...v1.0.0
[v0.7.1]: https://github.com/clok/kemba/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/clok/kemba/compare/v0.6.4...v0.7.0
[v0.6.4]: https://github.com/clok/kemba/compare/v0.6.3...v0.6.4
[v0.6.3]: https://github.com/clok/kemba/compare/v0.6.2...v0.6.3
[v0.6.2]: https://github.com/clok/kemba/compare/v0.6.1...v0.6.2
[v0.6.1]: https://github.com/clok/kemba/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/clok/kemba/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/clok/kemba/compare/v0.4.0...v0.5.0
[v0.4.0]: https://github.com/clok/kemba/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/clok/kemba/compare/v0.2.3...v0.3.0
[v0.2.3]: https://github.com/clok/kemba/compare/v0.2.2...v0.2.3
[v0.2.2]: https://github.com/clok/kemba/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/clok/kemba/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/clok/kemba/compare/v0.1.0...v0.2.0
