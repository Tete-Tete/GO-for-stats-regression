# Go Regression Analysis Tool

This repository contains a Go implementation for regression analysis. It includes a `main.go` file, which serves as the entry point, and a `regression_test.go` file for testing the functionality of the regression features. The project performs linear regression analysis on Anscombe's Quartet datasets, using Go to highlight the importance of data visualization.

## Overview
The main goal of this project is to provide a simple and effective tool for performing regression analysis. It aims to help users explore relationships between variables using linear regression techniques. The Anscombe's Quartet consists of four datasets, each having nearly identical statistical properties but appearing very different when graphed. This project helps highlight the importance of visualizing data.

## Features
- **Linear Regression Analysis**: Implements linear regression to analyze data trends, providing slope and intercept for each dataset.
- **Unit Testing**: Includes `regression_test.go` to validate the implementation with various test cases.
- **Graphical Output**: Generates graphs for each dataset, allowing users to visually compare and understand differences that might not be evident through statistical summaries alone.

## Requirements
- Go programming language installed (version 1.16 or higher recommended).
- The required package, `github.com/montanaflynn/stats`, should be installed using:
  ```sh
  go get github.com/montanaflynn/stats
  ```

## Installation From Git and Setup
### Step 1: Clone the Repository
Clone this repository to your local machine:
```sh
git clone <https://github.com/Tete-Tete/GO-for-stats-regression.git>
```

### Step 2: Run the Application
To build and run the Go application, run the following commands in your terminal:
```sh
go build -o regression
./regression
```
This will create an executable file named `regression` in your current directory.

## Running the Project
To run the project and see the regression analysis results:
```sh
go run main.go
```
The program will output the slope and intercept for each dataset in the Anscombe Quartet, formatted to 8 decimal places for higher precision.

## Testing
### Running Tests
This project includes test cases to validate the linear regression calculations. To run tests, use:
```sh
go test
```
This command will execute all unit tests, including those in `regression_test.go`, which verify the correctness of the linear regression calculations.

## Key Functions
- **`calculateRegression(x, y []float64) (float64, float64, error)`**:
  - Takes `x` and `y` values as inputs.
  - Returns the slope, intercept, and any error encountered during the calculation.

- **`toSeries(x, y []float64) stats.Series`**:
  - Converts `x` and `y` arrays into a `stats.Series` type for easier statistical calculations.

- **`testRegression()`**:
  - Loops over all four datasets and performs regression calculations.
  - Outputs the slope and intercept for each dataset to highlight differences in regression results.

## Dataset
The Anscombe Quartet datasets are defined in the code as follows:

- **Dataset 1**: `x1` and `y1`
- **Dataset 2**: `x2` and `y2`
- **Dataset 3**: `x3` and `y3`
- **Dataset 4**: `x4` and `y4` (includes an outlier to highlight the difference visually)

## Output
The output consists of the regression coefficients (slope and intercept) for each dataset, formatted to 8 decimal places for precision. This precision helps in understanding the subtle differences between datasets that have otherwise identical statistical measures.

For example:
```
Dataset 1: Regression Coefficients - Slope = 0.50000000, Intercept = 3.00000000
Dataset 2: Regression Coefficients - Slope = 0.50000000, Intercept = 3.00000000
...
```

## Roles of Programs and Data
- **Programs**: The primary program is contained in `main.go`, which is responsible for executing the regression analysis. The test file (`regression_test.go`) is used for unit testing to ensure that the regression functions work as intended.
- **Data**: The data used for regression analysis is defined directly in the code. However, you can modify `main.go` to read from a file or database if needed.

## Conclusion
The final program will display the execution time for each linear regression performed on the datasets in Anscombe's Quartet, highlighting the efficiency of the implemented algorithm. Additionally, the program will generate graphs for each dataset, allowing users to visually compare the datasets and understand the differences that statistical analysis alone might not reveal. The graphical output is particularly useful for demonstrating the importance of visualizing data along with statistical analysis. Users may find it insightful to compare these results with implementations in other languages, such as Python or R, to evaluate performance and visual output differences.

