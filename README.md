# csv-replace
Small tools for csv-parsing. At current time support only 1 csv file, and replace without regex. But this will be.

Install Go:
```
apt-get install -yq golang
go version
```

Install:
```
go build -ldflags="-s -w" -o csv-replace
sudo mv ./csv-replace /bin
```

Usage:
```
csv-replace -p /home/exports/1970-01-01.csv -f domains -r mysite.org
csv-replace -sp /home/exports/1970-01-01.csv -f domains -r mysite.org
```
