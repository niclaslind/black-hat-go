### Usage


#### This will output a successful run

```zsh
    $ go run main.go someC0mpl3xP@ssw0rd
    2021/05/14 10:59:06 hash = $2a$10$qYWi9QRnYHJRy6btoyHkCO95tGfbK8uOpD0NQmAd3Q9yR0b3wxoIC
    2021/05/14 10:59:06 [+] Authentication successful
```

#### This will output a failed run

```zsh
    $ go run main.go someWrongPassword
    2021/05/14 10:58:45 hash = $2a$10$tuPMbnk3AAv6mR0cJGHjOuW5yWB9VQJ826AfW9oVwcARaeqeuBHdC
    2021/05/14 10:58:45 [!] Auhentication failed
```