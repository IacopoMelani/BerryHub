FROM node:lts-alpine as stage

# Add Maintainer Info
LABEL maintainer="Iacopo melani <iacopomelani4@gmail.com>"

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# build app for production with minification
RUN npm run build


# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Iacopo melani <iacopomelani4@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/BerryHub

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Copy the Pre-built binary file from the previous stage
COPY --from=stage /app/dist .

EXPOSE 8888

CMD ["BerryHub"]