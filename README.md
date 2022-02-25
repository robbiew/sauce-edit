# 📝 SAUCE-EDIT
A command-line SAUCE record editor for ANSI art files.

<img src="/robbiew/sauce-edit/raw/main/screenshots/screen1.png" alt="Show Sauce" style="width: 100%;">


## INSTALL
- Binaries are in the Release folder.  Only Linux64 for now.
- Download and move somewhere on your path, rename, etc.

## USAGE
```go
sauce-edit -path /path/to/file.ans
```
- No quotes around the path name; relative to your current working directory. If no other arguments are used, it will just output existing SAUCE record.

## OPTIONAL ARGUMENTS
```
-title 'My New title'
-group 'Cool Art Group'
-author 'aLPHA'
```

- ***Single quotes*** around the arguments!

## License / Credit

sauce-edit utilizies [go-ansi](https://github.com/ActiveState/go-ansi) which is released under the BSD 3-Clause license. It was, in turn, based on [ansilove/C](https://github.com/ansilove/ansilove). Refer to the go-ansi license for legal stuff! I thank these folks for making such awesome code. 👍


## TO-DO
- Make Comment record editable
- Optionally save as a new file instead of overwriting
- Make it more of an interactive, term-based GUI form
- Test Windows, ARM
- Better theme/layout, colors :)
