# rma

Remote multi-server automation tool

## Example

```
go get github.com/itpkg/rma
rma -h
```

### Deploy task

```
rma g -s production
vi config/stages/production.toml

rma g -t deploy
vi config/tasks/deploy.toml

rma r -s production -t deploy
```

#### By host filter

```
HOSTS=server1,server2 rma r -s production -t deploy
```

OR

```
rma r --hosts=server1,server2 -s production -t deploy
```

#### By role filter

```
ROLES=app,web rma r -s production -t deploy
```

OR

```
rma r --roles=app,web -s production -t deploy
```
