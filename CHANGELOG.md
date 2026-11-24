# Changelog

## v0.0.9 (unreleased)

### Changed

- **Breaking** — Access generators through methods on `*faker.Faker`, for example `f.Person().Name()` and `f.Num().IntBt(1, 10)`. See the [English migration guide](docs/migration/v0.0.8-to-v0.0.9.md) or the [Japanese migration guide](docs/migration/v0.0.8-to-v0.0.9.ja.md) for the full field mapping and changes to client code.
- **Breaking** — The minimum Go version is now 1.26.4, up from 1.24.2.

### Removed

- **Breaking** — The exported `Faker.Rand` field has been removed. Use the six primitive generator accessors, or construct a standalone `core.Rand` for a custom random source.

### Added

- Added German (`de_DE`), British English (`en_GB`), French (`fr_FR`), Canadian English (`en_CA`) and Indian English (`en_IN`) providers for names, addresses, companies, job titles and domestic phone numbers.

### Fixed

- `Company().EIN()` returns an empty string when the locale has no U.S. EIN prefixes. It no longer generates a leading hyphen followed by seven digits or advances the random source in that case.
- Requests for unsupported full kana names now return an empty string without advancing the random source. Previously a missing kana callback could panic. A configured logger receives one warning for each unsupported EIN or full kana name request.
