## Ants

#### Ants Simulation on Hexagonal Grid
This program simulates the behavior of ants on a hexagonal grid using the Raylib-Go library. The simulation includes ants moving around, picking up and dropping leaves, and resting. The grid is displayed graphically, with ants and leaves represented by hexagons.

#### Features
- Hexagonal Grid: The simulation takes place on a hexagonal grid.
- Ants: Ants move around the grid, pick up leaves, and rest.
- Leaves: Leaves are scattered around the grid and can be picked up and dropped by ants.
- Graphical Display: The grid, ants, and leaves are displayed using Raylib-Go.

#### How to Run
- Ensure you have Go installed.
- Clone the repository.
- Run ```go mod tidy``` to install dependencies.
- Execute the program with ```go run main.go```.

#### Dependencies
- Raylib-Go

The resulting simulation can be seen in the newAnts.mp4 file

## Forest

#### Forest Fire Simulation
This program simulates the spread of fire in a forest using a grid-based approach. The forest is represented by a grid where each cell can be a tree, fire, empty space, or an axe. The user can interact with the simulation by moving the axe around the grid.

#### Features
- Grid-Based Forest: The forest is represented by a 20x20 grid.
- Fire Spread Simulation: Fire spreads to adjacent trees over time.
- User Interaction: The user can move an axe around the grid to cut down trees.
- Burned Percentage Calculation: The program calculates and displays the percentage of the forest that has been burned.
- Plot Generation: Generates a plot showing the percentage of the forest burned over time. (forest_fire.png)

#### How to Run
- Ensure you have Go installed.
- Clone the repository.
- Run ```go mod tidy``` to install dependencies.
- Execute the program with go run main.go.

#### Dependencies
- eiannone/keyboard
- gonum/plot

## Labs

Lab folders feature course work, including creating a server and experimenting with various mathematical properties
