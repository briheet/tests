# Testing

While i write applications via Test driven development, i tend to fuck up things and forget even the basics.

## For Testing

`make test`

## For Run

`make run`

## For Bench

`make bench`

## For coverage generation

`make gencover`
Now check your browser

## For building the image

Either
`docker build -t gotest:v1 .`
Or
`make image`

## For running the image

Either
`docker run -p 5001:5001 gotest:v1`
Or
`make runx`
