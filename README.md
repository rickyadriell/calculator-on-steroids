# calculator-with-steriods

## Objective

The objective of this project is to create a calculator api including addition, subtractions, multiplication and division. It is simple api but it will implement the following.

    - CI/CD using Github Actions
    - Unit Testing
    - Integrated Testing
    - Containerization using Docker
    - Hosting using AWS

## API Guide

### Addition

```API
POST /add
body {
    "a": 1,
    "b": 2
}
response {
    "result": 3
}
```

### Subtraction

```
POST /subtract
body {
    "a": 1,
    "b": 2
}
response {
    "result": -1
}
```

### Multiplication

```
POST /multiply
body {
    "a": 1,
    "b": 2
}
response {
    "result": 2
}
```

### Division

```
POST /divide
body {
    "a": 1,
    "b": 2
}
response {
    "result": 0.5
}
```
