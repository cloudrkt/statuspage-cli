# statuspage-cli

Statuspage.io CLI tool 

## Synopsis

This commandline tool interacts with the statuspage.io API. 

*note: Beware, this is work in progress and without warranty! It might be broken.*

## Features:

- [x] Create user
- [x] Delete user
- [x] Search user
- [ ] Everything else

## Installation:

Get the binary from the release page or build it yourself using go get.

### Configuration:

Get the pageid and apikey in your statuspage.io admin panel and create a .statuspage.yaml file in your home directory with the following content:

```yaml
 pageid: <pageid>
 apikey: <apikey>
 ```

You can also use environment variables with prefix SP_:

```bash
 export SP_PAGEID <pageid>
 export SP_APIKEY <apikey>
 ```

## statuspage options

```
      --config string   config file (default is $HOME/.statuspage.yaml)
      --debug           debug mode
  -h, --help            help for statuspage
      --version         version for statuspage
```

### Create a user

Create a subscriber through email adres. The subsciber *needs* to confirm the email from statuspage to receive notifications. The subscriber is then added to all the components by default.

```
statuspage subscriber create [email address]
```

#### Examples

```
statuspage subscriber create [email@example.org]
```

#### Options

```
  -h, --help   help for create
```

#### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.statuspage.yaml)
      --debug           debug mode
```

### Delete a user

Delete a subscriber through email adres

```
statuspage subscriber delete [email address]
```

#### Examples

```
statuspage subscriber delete [email@example.org]
```

#### Options

```
  -h, --help   help for delete
```

#### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.statuspage.yaml)
      --debug           debug mode
```



### Search subscriber 

Search a subscriber

#### Synopsis

Search a subscriber through email adres

```
statuspage subscriber search [email address]
```

#### Examples

```
statuspage subscriber search [email@example.org]
```

#### Options

```
  -h, --help   help for search
```

#### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.statuspage.yaml)
      --debug           debug mode
```
