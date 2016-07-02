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
rma generage -s production -n deploy
rma production deploy
```

#### By host filter

```
HOSTS=server1,server2 rma production deploy
```

OR

```
rma --hosts=server1,server2 production deploy
```

#### By role filter

```
ROLES=app,web rma production deploy
```

OR

```
rma --roles=app,web production deploy
```
