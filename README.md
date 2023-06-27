# Maze Generator

## About
I wanted to re-create the (now-defunct) [noops challenge maze generator](https://api.noopschallenge.com/mazebot/random) API so that I could continue to use the maze solvers I had created.

## Usage
```
mazebot - a maze generator API

Usage:
  mazebot [flags]
  mazebot [command]

Available Commands:
  api         Starts the mazebot API
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  print       Generates a maze and prints it out

Flags:
  -h, --help   help for mazebot

Use "mazebot [command] --help" for more information about a command.
```

### API
TODO:
- [ ] Handle solution submissions
- [ ] Support "race" mode 

```
Starts the mazebot API

Usage:
  mazebot api [flags]

Flags:
  -h, --help       help for api
      --host ip    Specifies the host addresses to respond to (default 127.0.0.1)
      --port int   Changes the server port (default 8080)
```

Matching the `noops` API, clients can send `GET` requests to [/mazebot/random?minSize=12&maxSize=12](http://localhost:8080/mazebot/random?minSize=12&maxSize=12). And example response looks like:

```json
{
    "endingPosition": [4,4],
    "map": [
        ["X","X","X","X","X","X","X","X","X","X","X","X"],
        ["X","X"," "," ","X"," ","X"," "," ","X"," ","X"],
        ["X"," ","X"," "," "," "," "," ","X"," "," ","X"],
        ["X"," "," "," ","X"," ","X","X","X"," ","X","X"],
        ["X","X"," ","X","B","X"," "," "," "," "," ","X"],
        ["X"," "," "," "," "," "," ","X"," ","X"," ","X"],
        ["X","X"," ","X"," ","X"," "," ","X","X"," ","X"],
        ["X"," "," "," ","X"," ","X"," "," "," ","X","X"],
        ["X","X"," ","X"," "," ","X","A","X"," "," ","X"],
        ["X","X"," "," "," ","X","X"," ","X"," ","X","X"],
        ["X"," "," ","X"," "," "," ","X"," "," "," ","X"],
        ["X","X","X","X","X","X","X","X","X","X","X","X"]
    ],
    "name": "12x12",
    "path": "/mazebot/random",
    "startingPosition": [8,7]
}
```

### Print
In order to help visualize the results of the maze generation, you can simply print out a generated maze to the CLI. For easier visual validation of the generated paths, the `X` wall character is replaced with a `▓` character. Ideally, a monospaced font would be utilized in order to show accurate proportions.
```
Generates a maze and prints it out

Usage:
  mazebot print size [flags]

Flags:
  -h, --help   help for print
```

Example usage:
```
> mazebot print 24

Generating a 24x24 maze
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
▓▓ ▓ ▓A▓  ▓  ▓  ▓▓▓    ▓
▓       ▓  ▓ ▓ ▓ ▓ ▓ ▓ ▓
▓ ▓▓ ▓ ▓  ▓         ▓  ▓
▓▓  ▓  ▓ ▓ ▓▓ ▓ ▓ ▓▓▓▓ ▓
▓▓▓  ▓         ▓ ▓  ▓  ▓
▓▓  ▓▓ ▓ ▓▓ ▓ ▓   ▓ ▓ ▓▓
▓  ▓    ▓ ▓  ▓ ▓▓      ▓
▓ ▓ ▓ ▓▓   ▓▓   ▓ ▓▓▓▓ ▓
▓ ▓      ▓▓ ▓ ▓▓      ▓▓
▓  ▓ ▓ ▓▓   ▓ ▓▓ ▓▓ ▓  ▓
▓ ▓   ▓   ▓▓      ▓  ▓ ▓
▓  ▓ ▓ ▓ ▓ ▓ ▓▓ ▓  ▓  ▓▓
▓▓     ▓ ▓     ▓  ▓▓ ▓ ▓
▓▓ ▓ ▓     ▓▓ ▓▓ ▓     ▓
▓   ▓▓▓ ▓ ▓  ▓  ▓ ▓▓ ▓ ▓
▓▓ ▓ ▓  ▓   ▓ ▓ ▓ ▓ ▓ ▓▓
▓▓    ▓ ▓ ▓            ▓
▓  ▓ ▓▓ ▓ ▓ ▓▓▓▓▓▓▓ ▓▓ ▓
▓ ▓    ▓ ▓▓       ▓ ▓  ▓
▓ ▓ ▓▓     ▓ ▓ ▓ ▓ ▓ ▓ ▓
▓  ▓▓ ▓ B▓ ▓  ▓      ▓▓▓
▓▓    ▓ ▓   ▓  ▓ ▓ ▓   ▓
▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
Took 343.666µs
```
