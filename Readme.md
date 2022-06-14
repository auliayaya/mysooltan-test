# My Sooltan Backend Developer Assessment Test


Hai everyone, this is repo for completing  My Sooltan Backend Developer Assessment Test.


## Notes

Here i explain a little bit about what i've done. Because my development environment is using windows, i use example using windows directory.

I create CLI Program:

- Get file from local directory  and converting file log to plainText as default
  - `mysooltan  -p dpkg.log`
  - `mysooltan -p C:\Users\ACER\dpkg.log -t text -o  C:\Users\ACER\CI\GOLANG/mysooltan-test`
- Converting file log to Plaintext with Flag `-t text` and save it to current directory
  - `mysooltan  -p dpkg.log -t text .`
- Converting file log to JSON with Flag `-t json` and save it to current directory
  - `mysooltan  -p dpkg.log -t json .`
- Converting file log to JSON with Flag `-t json` and save it to choosen directory path
  - `mysooltan -p dpkg.log -t text -o  C:\Users\ACER\CI\GOLANG/mysooltan-test`
- Getting detail information about current CLI command
  - `mysooltan -h`


If you wanna try this repo on your local machine, you can follow the instruction in subcontent getting started

## Getting Started

- Clone this repository
- Run `go mod tidy`
- Then run `make cli` on your command line terminal
