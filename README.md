# 📝 SAUCE-EDIT
A command-line SAUCE record editor for ANSI art files.

## EXAMPLE
- Use ```-path``` argument to show SAUCE record:
<img src="https://github.com/robbiew/sauce-edit/blob/main/screenshots/screen1.png" width="100%" title="show sauce">

- Use other arguments to edit SAUCE record:
<img src="https://github.com/robbiew/sauce-edit/blob/main/screenshots/screen2.png" width="100%" title="edit title">
  

## INSTALL
- Binaries are in the Release folder.  Windows64 and Linux64 for now -- feel free to build and test on other platforms!
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

- ❗ ***Single quotes*** around the arguments (otherwise Go will interpret special characters -- and that's not good!)

## License / Credit

sauce-edit utilizies [go-ansi](https://github.com/ActiveState/go-ansi) which is released under the BSD 3-Clause license. It was, in turn, based on [ansilove/C](https://github.com/ansilove/ansilove). I thank these folks for making such awesome code. 👍


## TO-DO
- Make Comment record editable
- Optionally save as a new file instead of overwriting
- Make it more of an interactive, term-based GUI form
- Test ARM
- Better theme/layout, colors :)
