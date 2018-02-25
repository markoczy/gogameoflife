# Game of Life Simulator in Go

## Usage

### Command Line Arguments

| value   | type   | description                 | default          |
| ------- | ------ | --------------------------- | ---------------- |
| debug   | flag   | enable debug logs           | unset (false)    |
| density | float  | seed density, if applicable | 0.1              |
| fps     | int    | frames per second           | 20               |
| height  | int    | height of the panel         | 600              |
| scale   | int    | pixel scaling               | 4                |
| seeder  | string | seed method                 | "random.default" |
| width   | int    | width of the panel          | 600              |

> Note: You can always type `gol -h` to retreive this information

### Examples

- Start in Debug Mode: `gol -debug`
- Change Seed Density: `gol -density=0.2`
- Set window dimensions: `gol -width=2400 -height=1200`
- Use glider seeder: `gol -seeder=random.gliders`

> Other seeder implementations are work in progress, at the moment there's only
`default.random` which randomizes each cell separatly and `random.gliders` which 
creates a set of gliders at random