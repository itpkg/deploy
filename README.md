# rma

Remote multi-server automation tool

## Example

```
go get github.com/itpkg/rma
rma -h
rma init
```

### Deploy task

```
rma g -s production -n deploy
rma r -s production -n deploy
```

#### By host filter

```
HOSTS=server1,server2 rma r -s production -n deploy
```

OR

```
rma r --hosts=server1,server2 -s production -n deploy
```

#### By role filter

```
ROLES=app,web rma r -s production -n deploy
```

OR

```
rma r --roles=app,web -s production -n deploy
```
