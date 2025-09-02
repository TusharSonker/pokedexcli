## pokedexcli

Command‑line Pokédex written in Go. It lets you explore Pokémon location areas, list and inspect Pokémon, and simulate catching them while caching PokéAPI responses locally for faster repeated queries.

### Features
* List paginated location areas (`map`, `mapb` for backward pagination)
* Explore a specific location area to see encountered Pokémon (`explore <location-area>`)
* Catch Pokémon with a simple probability mechanic (`catch <pokemon-name>`)
* Maintain an in‑memory Pokédex of caught Pokémon (`pokedex`)
* Inspect detailed stats & types of caught Pokémon (`inspect <pokemon-name>`)
* In‑memory TTL cache for PokéAPI responses (configurable expiration)

### Requirements
* Go 1.23+
* Internet connection (for live PokéAPI calls)

### Install / Build
```bash
git clone https://github.com/TusharSonker/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
./pokedexcli
```

Or run directly:
```bash
go run .
```

### Usage (REPL Commands)
| Command | Syntax | Description |
|---------|--------|-------------|
| help | `help` | Show help menu |
| map | `map` | List next page of location areas |
| mapb | `mapb` | List previous page of location areas |
| explore | `explore <location-area>` | Show Pokémon in a location area |
| catch | `catch <pokemon-name>` | Attempt to catch a Pokémon |
| pokedex | `pokedex` | List caught Pokémon |
| inspect | `inspect <pokemon-name>` | Show stats/types for a caught Pokémon |
| exit | `exit` | Quit the CLI |

Example session:
```
> help
> map
> map            # again for next page
> explore canalave-city-area
> catch pikachu
> pokedex
> inspect pikachu
> exit
```

### Catch Probability
The current implementation generates a random number and compares it to a threshold. Feel free to refine by incorporating base experience or other stats for a more nuanced mechanic.

### Internal Architecture
* `main.go` + `repl.go`: REPL loop & command dispatch
* `command_*.go`: Individual command callbacks
* `internal/pokeapi`: Thin client for PokéAPI (locations & Pokémon)
* `internal/pokecache`: Simple TTL cache with background reaper goroutine

### Caching
The cache stores raw JSON responses keyed by full request URL and purges entries older than the configured interval (set when constructing the client in `main.go`). This reduces redundant HTTP calls while exploring or re‑listing pages.

### Testing
Run all tests:
```bash
go test ./...
```

### Potential Improvements
* Improve error handling (return actual errors instead of `nil` with empty structs)
* Fix/rename small typos (`GetLocationAres` → `GetLocationAreas`)
* Add mutex protection to cache reads (currently only writes lock)
* Richer catch probability algorithm & feedback (e.g., show percentage)
* Persist caught Pokémon to disk (JSON or BoltDB) between runs
* Add colored output for better UX
* Add integration tests hitting a mock PokéAPI server

### Contributing
1. Fork the repo
2. Create a branch: `git checkout -b feature/your-feature`
3. Commit changes: `git commit -m "feat: add your feature"`
4. Push: `git push origin feature/your-feature`
5. Open a Pull Request

### License
MIT (add a LICENSE file if publishing formally)

### Disclaimer
This project uses the public PokéAPI. Pokémon © Nintendo / Game Freak / Creatures. This is an educational CLI tool.

