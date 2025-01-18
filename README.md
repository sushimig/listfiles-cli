# listfiles-cli
listfiles-cli is a tool which can display a directory name and its file's names. Use this if you want to visualize directory structure and format into some kinds of data to share documents!

## Installation

```
go install github.com/sushimig/listfiles-cli@latest
```

!n case you failed, try to change previous version!

## Usage

```
listfile-cli <path> [<flags>]
```

you can get some list-formats! 
- simple text


```
+---------------+
| LISTFILES-CLI |
+---------------+
| .git          |
| LICENSE       |
| cmd           |
| go.mod        |
| go.sum        |
| internal      |
| main.go       |
+---------------+
```

- markdown

```
| listfiles-cli |
|------|
| .git |
| LICENSE |
| README.md |
| cmd |
| go.mod |
| go.sum |
| internal |
| main.go |
```

- json

```
{
                "directory": "listfiles-cli",
                "files": [
                                ".git",
                                "LICENSE",
                                "README.md",
                                "cmd",
                                "go.mod",
                                "go.sum",
                                "internal",
                                "main.go"
                ]
}
```

Get a help from the following command!

```
listfile-cli -h
```

## Support
If you face a problem, feel free to write detail in issue on Github!

## licensing
MIT License
