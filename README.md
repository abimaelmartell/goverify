# Microservice for Email Verification


## Installation
```
go get github.com/abimaelmartell/goverify

goverify
```

It will start a web server at localhost:8080


## Usage

### GET /verify?email={EMAIL}

Sample Response

    {
      result: "NoMxServersFound",
      mailbox_exists: false,
      is_catch_all: false,
      is_disposable: false,
      email: "abimex@gmail.com",
      domain: "gmail.com",
      user: "abimex"
    }

#### result (String)

| Code | Description |
| --- | --- |
| NoMxServersFound | The server is not running a mail server |
| InvalidEmailAddress | The email address is not valid |

#### mailbox_exists (Bool)
True if the email address exists on the server

#### is_catch_all (Bool)
True if the email server is catch all

#### is_disposable (Bool)
True if the email server is listed as a disposable email provider

#### email (String)
The email address

#### domain (String)
The domain name of the email address

#### user (String)
The username for the domain
