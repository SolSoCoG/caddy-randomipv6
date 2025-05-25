# caddy-randomipv6

Trying to create an IPv6_Randomizer to hide the users ip at reverse proxies. Made for my mastodon instance, previously using LUA with openresty.

After quite some config attempts, this is how it works in a Caddyfile.

I don't know why, but *X-Forwarded-For* kept getting overwritten until I did a [header_up](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy#headers) with ... *the X-Forwarded-For* Header? Which then used the modules generated random ipv6.

For randomizing [within php environment](https://caddyserver.com/docs/caddyfile/directives/php_fastcgi):

```
domain.tld {
  root /var/www/html/public
  file_server

  randomipv6
  php_fastcgi unix//run/php/php8.4-fpm.sock {
    header_up X-Forwarded-For {header.X-Forwarded-For}
  }
}
```


For randomizing for [a reverse proxy](https://caddyserver.com/docs/caddyfile/directives/reverse_proxy):

```
domain.tld {
  randomipv6
  reverse_proxy http://127.0.0.1:8080 {
    header_up X-Forwarded-For {header.X-Forwarded-For}
  }
}
```
