# uWave Self-Learning Challenge
Submission for uWave's self learning challenge. This README details how to get the backend service started.

Access the plans to the challenge <a href="https://drive.google.com/drive/folders/1ntg9ZdPEqFMkd3WRzKy7eoi5BvCB0mYj?usp=sharing">here</a>.

# Getting started
There are two ways to get the service going, either through Docker or interfacing with Go directly. Below are instructions for both, feel free to use either one.

## Docker Guide
1. Build the Docker image from the Dockerfile.
```
docker build -t uwave:latest ./
```

2. Create and run the container. 
```
docker run -d -p 80:80 --name uwave uwave:latest
```
<i>(Note: We are exposing and binding port 80 from the Docker container to your localhost.)</i>

3. Simply type 'localhost' in your web browser to view the documentation to begin.
4. End the service with the following command.
```
docker stop uwave
```
5. To remove the containers and images, use the following commands.
```
docker rm -f uwave
docker rmi uwave:latest
```

## Go Guide
1. Build the Go package.
```
go build .
```
2. Run the Go package.
```
./uwave_challenge
```
<i>(Note: This command would differ according to your OS. On a Windows system, the command would include '.exe' at the end.)</i>

3. End the program with ```ctrl + c``` on your terminal.