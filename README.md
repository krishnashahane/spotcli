# spotcli

A terminal-first Spotify client written in Go. Search Spotify, control playback, manage your library and playlists, inspect devices, and produce machine-readable output for scripts.

> **Important:** spotcli uses Spotify web-player/internal endpoints and browser session cookies. These endpoints are not the same as Spotify's supported public Web API and may change without notice. spotcli does **not** bypass or remove Spotify rate limits.

## Features

- Search tracks, albums, artists, playlists, shows, and episodes
- Playback: play, pause, next, previous, seek, volume, shuffle, repeat
- Queue management
- Library and followed-artist management
- Playlist creation and track management
- Device listing and transfer
- Browser-cookie authentication through `sweetcookie`
- Human, plain-text, and JSON output
- Multiple engines: `connect`, `web`, `auto`, and macOS-only `applescript`
- Configurable profiles, market, language, device, and request timeout

## Requirements

- Go 1.24 or newer for building from source
- A Spotify account
- A supported browser profile containing an active Spotify web session for cookie authentication
- macOS for the `applescript` engine

## Install

### From source

```bash
go install github.com/krishnashahane/spotcli/cmd/spotcli@latest
```

Verify the installation:

```bash
spotcli --version
spotcli --help
```

### Build locally

```bash
git clone https://github.com/krishnashahane/spotcli.git
cd spotcli
go test ./...
go build ./cmd/spotcli
```

## Quick start

Import an existing Spotify browser session:

```bash
spotcli auth import --browser chrome
```

Then try:

```bash
spotcli search track "weezer" --limit 5
spotcli play spotify:track:7hQJA50XrCWABAu5v6QZ4i
spotcli status
spotcli device list
```

If your browser uses a named profile:

```bash
spotcli auth import --browser chrome --browser-profile "Profile 1"
```

Check authentication without exposing cookie values:

```bash
spotcli auth status
```

## Authentication and cookies

spotcli reads Spotify cookies from a supported browser or from a previously exported cookie file.

Browser import:

```bash
spotcli auth import --browser chrome
```

By default, spotcli targets the `spotify.com` cookie domain and stores its cached cookie file under the application's configuration directory.

You can inspect or clear the cached session:

```bash
spotcli auth status
spotcli auth clear
```

Treat the cookie cache like a password. Do not commit it, paste it into issues, or share it. spotcli writes cached cookies with restrictive local file permissions.

## Engines

### `connect`

Default engine. Uses Spotify's internal Connect/playback services for playback, queue, device operations, and internal search/info paths.

### `web`

Uses Spotify Web API-style endpoints authenticated with a browser-derived web-player token. Playback/search/info have fallback behavior where supported.

### `auto`

Chooses between Connect and Web API-style paths to maximize supported operations. It can fall back when an operation is unsupported.

### `applescript`

Uses the installed Spotify macOS application through AppleScript where supported, with Web API-style fallback where available. This engine is macOS-only.

Because these engines depend on Spotify's web-player behavior, upstream changes can temporarily break functionality. The project cannot guarantee compatibility with every Spotify release.

## Usage

```text
spotcli [global flags] <command> [args]
```

Global flags:

- `--config <path>` — configuration file
- `--profile <name>` — profile name
- `--timeout <duration>` — HTTP timeout; default `10s`
- `--market <cc>` — market/country code
- `--language <tag>` — language/locale
- `--device <name|id>` — target playback device
- `--engine <auto|web|connect|applescript>` — API engine
- `--json` — structured JSON output
- `--plain` — line-oriented output
- `--no-color` — disable human-output colors
- `-q, --quiet` — reduce output
- `-v, --verbose` — verbose output
- `-d, --debug` — debug output
- `--no-input` — disable interactive prompts

`--json` and `--plain` are mutually exclusive.

### Commands

- `auth status|import|clear`
- `search track|album|artist|playlist|show|episode`
- `track info`, `album info`, `artist info`, `playlist info`, `show info`, `episode info`
- `play`, `pause`, `next`, `prev`, `seek`, `volume`, `shuffle`, `repeat`, `status`
- `queue add|show`
- `library tracks|albums|artists|playlists`
- `playlist create|add|remove|tracks`
- `device list|set`

For the complete command specification:

```text
docs/spec.md
```

## Scripting

JSON output is suitable for programs:

```bash
spotcli --json search track "daft punk" --limit 5
```

Plain output is useful for shell pipelines:

```bash
spotcli --plain search artist "radiohead"
```

## Configuration

Configuration is stored under the platform's user configuration directory unless `--config` is supplied.

Useful environment overrides include:

- `SPOTCLI_CONFIG`
- `SPOTCLI_PROFILE`
- `SPOTCLI_TIMEOUT`
- `SPOTCLI_MARKET`
- `SPOTCLI_LANGUAGE`
- `SPOTCLI_DEVICE`
- `SPOTCLI_ENGINE`
- `SPOTCLI_JSON`
- `SPOTCLI_PLAIN`
- `SPOTCLI_NO_COLOR`
- `SPOTCLI_QUIET`
- `SPOTCLI_VERBOSE`
- `SPOTCLI_DEBUG`
- `SPOTCLI_NO_INPUT`

Advanced overrides:

- `SPOTCLI_TOTP_SECRET_URL` — use a custom HTTPS or local-file TOTP secret source
- `SPOTCLI_CONNECT_VERSION` — override the Connect client version sent to Spotify

Do not point secret-related configuration at untrusted files or remote sources.

## Security

spotcli handles browser session cookies and access tokens. The project therefore follows a few basic safeguards:

- Cookie caches are written with mode `0600`.
- Configuration directories are created with mode `0700`.
- Remote response bodies used for API errors and JSON decoding are bounded.
- Custom TOTP secret sources reject plain HTTP.
- Cookie import is restricted to Spotify domains/subdomains.
- IDs inserted into Connect endpoint paths are URL-escaped.
- Access tokens used in the dealer WebSocket URL are query-encoded.

For a local installation, also protect the browser profile and operating-system account that can access the Spotify session.

## Troubleshooting

### Authentication fails

1. Make sure Spotify is open and logged in in the selected browser.
2. Re-import the cookies.
3. Check `spotcli auth status`.
4. Try a different supported browser profile if you have multiple profiles.

### Playback works but search/info fails

Try another engine:

```bash
spotcli --engine auto status
spotcli --engine web search track "weezer"
```

Spotify can change internal web-player APIs independently of spotcli releases.

### A command reports unsupported operation

Some operations are available through the Web API-style engine but not through the internal Connect engine. Use `--engine web` or `--engine auto` where appropriate.

## Development

Run the full test suite before submitting changes:

```bash
go test ./...
go vet ./...
```

Format code with the project's formatting scripts:

```bash
./scripts/format.sh
./scripts/lint.sh
```

The CI workflow also runs the configured linter and coverage checks.

## Legal and responsible use

spotcli is an independent project and is not affiliated with Spotify.

Use it only with accounts and browser sessions you are authorized to use. Respect Spotify's Terms of Service, applicable laws, and the privacy/security of anyone whose account or device you access.

## License

MIT
