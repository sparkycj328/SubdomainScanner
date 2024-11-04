# SubdomainScanner
Another Subdomain Fuzzer

## Table of Contents
 - [Background](#background)
 - [Setup and Usage](#setup-and-usage)

### Background
Program will make requests to either the specified DNS server or the default DNS server (8.8.8.8).

### Setup and Usage
1. **Setup**
2. **Usage**
    - `domain`: The target domain. Required to be set by user
    - `wordlist`: The wordlist to be used for fuzzing. Required to be set by user
    - `count`: The amount of concurrent workers to be used. Default is 100
    - `server`: The DNS server to which requests will be sent. Default is 8.8.8.8