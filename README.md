# cap

Remote multi-server automation tool

## Example

```
export PATH=$GOPATH/bin:$PATH
go get -u github.com/itpkg/cap
which cap
cap -h
```

### Deploy task

```
cap g -s production
vi config/stages/production.toml

cap g -t deploy
vi config/tasks/deploy.toml

cap r -s production -t deploy
```

#### By host filter

```
HOSTS=server1,server2 cap r -s production -t deploy
```

OR

```
cap r --hosts=server1,server2 -s production -t deploy
```

#### By role filter

```
ROLES=app,web cap r -s production -t deploy
```

OR

```
cap r --roles=app,web -s production -t deploy
```
