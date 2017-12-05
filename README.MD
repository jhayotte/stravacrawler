# Strava Crawler

This project is a concrete sample of a program that benefit from golang concurrency.

The goal is to grab a HTML value on many pages and do the sum.

`$ go build -race; ./stravacrawler`

Output:
```json
# 1.80
# 2.80
# 4.70
# 4.70
# 8.90
...
# 678.20
# 697.60
# 723.60
# 743.30
end: 743.30 km
```