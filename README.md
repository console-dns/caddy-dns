# caddy-dns

## 安装说明

安装最新版本的 golang 和 xcaddy，然后使用如下命令创建可用的 caddy 可执行文件

```bash
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
~/go/bin/xcaddy build --with=github.com/console-dns/caddy-dns@main
./caddy list-modules | grep console
```

## 使用说明

编辑 Caddyfile , 按照类似于如下的配置填入内容

```conf

*.example.com {
    tls  {
        ca https://ca.example.com/acme/acme/directory
        resolvers <your internal dns>:53
        propagation_timeout 10s
        dns console {
            server https://dns.example.com
            token <client token>
        }
    }
}
```