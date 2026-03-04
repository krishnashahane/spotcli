# spotcli CLI spec (v0.2.0)

One-liner: Spotify power CLI using web cookies; search + playback control.
Parser: Kong.
Cookies: sweetcookie (local sweetcookie).
Output: human by default; `--plain` or `--json`.
Color: on by default; respects `NO_COLOR`, `TERM=dumb`, `--no-color`.
Platforms: macOS, Linux, Windows.

## Usage

```
spotcli [global flags] <command> [args]
```

## Global flags

- `-h, --help`
- `--version`
- `-q, --quiet`
- `-v, --verbose`
- `-d, --debug`
- `--json`
- `--plain`
- `--no-color`
- `--config <path>` default: `os.UserConfigDir()/spotcli/config.toml`
- `--profile <name>` default: `default`
- `--timeout <dur>` default: `10s`
- `--market <cc>` default: account market or `US`
- `--language <tag>` default: `en`
- `--device <name|id>` default: active device
- `--engine <auto|web|connect|applescript>` default: `connect` (`applescript` is macOS-only)
- `--no-input`

## Commands

### auth

- `spotcli auth status`
- `spotcli auth import`
  - flags: `--browser <chrome|brave|edge|firefox|safari>` default: `chrome`
  - `--browser-profile <name>`
  - `--cookie-path <file>`
  - `--domain <host>` default `spotify.com`
- `spotcli auth clear`

### search

- `spotcli search track <query> [--limit N] [--offset N]`
- `spotcli search album <query> [--limit N] [--offset N]`
- `spotcli search artist <query> [--limit N] [--offset N]`
- `spotcli search playlist <query> [--limit N] [--offset N]`
- `spotcli search episode <query> [--limit N] [--offset N]`
- `spotcli search show <query> [--limit N] [--offset N]`

### info

- `spotcli track info <id|url>`
- `spotcli album info <id|url>`
- `spotcli artist info <id|url>`
- `spotcli playlist info <id|url>`
- `spotcli show info <id|url>`
- `spotcli episode info <id|url>`

### playback

- `spotcli play [<id|url>]` (track/album/playlist/show)
  - optional: `--type <track|album|playlist|show|episode>` for raw IDs
  - artist URIs play top tracks (starts with the first)
- `spotcli pause`
- `spotcli next`
- `spotcli prev`
- `spotcli seek <ms|mm:ss>`
- `spotcli volume <0-100>`
- `spotcli shuffle <on|off>`
- `spotcli repeat <off|track|context>`
- `spotcli status`

### queue

- `spotcli queue add <id|url>`
- `spotcli queue show`
- `spotcli queue clear` (not supported by Spotify API yet)

### library

- `spotcli library tracks list [--limit N]`
- `spotcli library tracks add <id|url...>`
- `spotcli library tracks remove <id|url...>`
- `spotcli library albums list [--limit N]`
- `spotcli library albums add <id|url...>`
- `spotcli library albums remove <id|url...>`
- `spotcli library artists list [--limit N] [--after <artist-id>]`
- `spotcli library artists follow <id|url...>`
- `spotcli library artists unfollow <id|url...>`
- `spotcli library playlists list [--limit N]`

### playlists

- `spotcli playlist create <name> [--public] [--collab]`
- `spotcli playlist add <playlist> <track...>`
- `spotcli playlist remove <playlist> <track...>`
- `spotcli playlist tracks <playlist> [--limit N]`

### devices

- `spotcli device list`
- `spotcli device set <name|id>`

## Output contract

- stdout: primary results; human or machine modes.
- stderr: warnings/errors/logs.
- `--plain`: stable, line-oriented, tab-separated fields.
- `--json`: stable, documented keys per command.

## Engines

- `auto`: connect first; fall back to web for unsupported features or rate limits.
- `connect`: internal connect-state endpoints for playback; GraphQL for search/info.
- `web`: Web API endpoints; search/info/playback auto-fallback to connect when rate limited.

## Exit codes

- `0` success
- `1` generic failure
- `2` invalid usage/validation
- `3` auth/cookies missing or invalid
- `4` network/timeouts

## Config / env

- Env prefix: `SPOTCLI_`
- Precedence: flags > env > config
- Secrets: never via flags; use browser cookies only.
- Overrides:
  - `SPOTCLI_TOTP_SECRET_URL` (http(s) or `file://...`)
  - `SPOTCLI_CONNECT_VERSION` (connect playback client version)

## Examples

- `spotcli auth import --browser chrome`
- `spotcli search track "weezer" --limit 5 --plain`
- `spotcli play spotify:track:7hQJA50XrCWABAu5v6QZ4i`
- `spotcli device list --json`
- `spotcli playlist create "Road Trip" --public`
