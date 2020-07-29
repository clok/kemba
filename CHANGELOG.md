<a name="unreleased"></a>
## [Unreleased]


<a name="v0.6.2"></a>
## [v0.6.2] - 2020-07-29
### Chore
- updated docs on exported struct


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


[Unreleased]: https://github.com/clok/kemba/compare/v0.6.2...HEAD
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
