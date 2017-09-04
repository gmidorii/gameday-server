# gameday-server

### build
```
% go build

% ./gameday-server -h
Usage of ./gameday-server:
  -c toml
        config file type toml (default "./config.toml")
```

### run
```
% ./gameday-server -c config.toml
```

### Config
config file type is `toml`  

ex)
```toml
outernalurl="URL"

# Animal DB is supposed mysql
[Animal]
user="user name"
password="passwd"
host="host name"
```

### All API
| path      | content          |
|-----------|------------------|
| /ping     | health check     |
| /outernal | outernal connect |
| /animal   | animal db info   |
