# SAUCE-EDIT
A command-line SAUCE record editor for ANSI art files.

## INSTALL
- Binaries are in the Release folder.  Only Linux64 for now.
- Download and move somewhere on your path, rename, etc.

## USAGE
```go
sauce-edit -path /path/to/file.ans
```
No quotes around the path name; relative to your current working directory. If no other arguments are used, it will just output existing SAUCE record.

## OPTIONAL ARGUMENTS
```
-title 'My New title'
-group 'Cool Art Group'
-author 'aLPHA'
```

Single quotes only on the arguments!

## TO-DO
- Make Comment record editable
- Make a term-based GUI form
- Test Windows, ARM
- Better theme laout and colors :)
