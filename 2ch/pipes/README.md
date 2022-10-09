# Work making use of unix pipe

-unix pipe is useful when send some programs output to another program
```
$ echo "test" case  | wc -1
1
```
- In go application, we can read left arguments to os.Stdin
