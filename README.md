# gorock

<img width="543" alt="image" src="https://github.com/howen02/gorock/assets/108785851/ea23e655-c0e4-42fe-90f9-4dc62dc8cec1">

## Why
I created gorok because I often found myself switching between free-tier ngrok domains during my software engineering internship. It was quite a hassle and I figured to make a CLI tool since I was learning Go.

## How it works
gorok simply helps you replace all the ngrok domains in your `.env` so you won't have to manually find and change them.

## Usage
1. Download the source code
2. Create an `accounts.txt` file to store the necessary information in the following format 👇🏻
```
email, authToken, domain
abc@gmail.com, eR7FpT2bL6S3oC9zN8W2E0hV6Y4M7I9uG2X3D5K1eR7FpT2bL, stylish-sapphire-snail.ngrok-free.app
123@gmail.com, A5bG8pR3fT2LbS1oCzN8W2E0hV6Y4M7I9uG2X3D5K1eR7FpT2, fuzzy-flamingo-frog.ngrok-free.app
```
3. Run `go build main.go`
4. Move the binary file and `accounts.txt` to your development space
5. Start using gorok with `./main` !
   
<img width="546" alt="image" src="https://github.com/howen02/gorock/assets/108785851/21506cce-1397-4e9f-abc5-73649eeeb6fc">
