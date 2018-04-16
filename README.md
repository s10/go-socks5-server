# go-socks5-proxy

Simple SOCKS5 server using go-socks5 with auth

# Start container with proxy
```docker run -d --name socks5-proxy -p 1080:1080 -e PROXY_AUTH='{"<PROXY_USER_1>": "<PROXY_PASSWORD_1>", "<PROXY_USER_2>": "<PROXY_PASSWORD_2>"} esten/go-socks5-proxy```

where `PROXY_AUTH` is a JSON map of authorized user names and the respective passwords.

# Test running service
```curl --socks5 <docker machine ip>:1080 -U <PROXY_USER>:<PROXY_PASSWORD> https://api.ipify.org```

Result must show docker host IP address.
