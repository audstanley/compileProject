# Compilers Project
Authors: Richard Stanley, Alex Truong, Kenny Chao

Language used: Go

This project was written to take a single language in the from of:

```
program a2018 ;
    var
    qwe , ab1 , cd , e33a , d18 : integer ;
    begin
        ab1 = 3 ; // some comment
        /*
            some block comment
        */
        cd = 4 ;
        e33a = 5 ;
        show ( ab1 ) ;
        qwe = 1 * -8 ;
        d18 = ab1 * ( cd + 2 * e33a ) ;
        show ( d18 );
    end
```
It converts this language into golang, a language written by [Google's Golang](https://golang.org/) and uses a sub processed shell script in order to compile the language into multiple binaries.  Currently, this must be run on a Linux computer.


![boxplot]()

This was generated with 3,500 iterations of transpiling the data from the original format to GoLang.  Each iteration spaced by 100 millisecond delay.  The plot does not include file read time, and binary compilation using the GoLang Compiler.  The box plot is only a measurement of tanspile time in microseconds.


## Tree of the Files:
```

.
|-- CompileTimeBoxPlot.png
|-- README.md
|-- compileTime.csv          // This csv was used to calculate the box plot in microseconds
|-- finalProject             // main.go's compiled binary
|-- format.go                // 
|-- go-executable-build.sh   // Compiles multiple binaries in output/* as subprocess
|-- main.go
|-- mathrhs.go               // checks expression using the table below
|-- originalCode.txt         // source of the code.
|-- output
|   |-- main.go              // This is the output of the transpiled code
|   |-- main.go-darwin-amd64
|   |-- main.go-linux-amd64
|   |-- main.go-windows-386.exe
|   |-- main.go-windows-amd64.exe
|   `-- sanitized.txt
|-- sanitize.go             // cleans up the originalCode.txt before >> output/sanitize.txt
`-- validateDefinition.go   // validates variables, before checking expression with mathrhs.go

```

## Matrix used to complete mathematical expressions:

| State | a  | +  | -  | *  |  / | (  | )  | ;  | =  | b  |
|-------|----|----|----|----|----|----|----|----|----|----|
|  E    | 10 | -1 | -1 | -1 | -1 | 10 | -1 | -1 | 10 | 10 |
|  Q    | -1 | 11 | 12 | -1 | -1 | -1 | ;  |  ; | -1 | -1 |
|  T    | 13 | -1 | -1 | -1 | -1 | 13 | -1 | -1 | -1 | 13 | 
|  R    | -1 |  ; |  ; | 14 | 15 | -1 |  ; |  ; | -1 | -1 | 
|  F    |  a | -1 | -1 | -1 | -1 | 16 | -1 | -1 | -1 | b  | 
|  S    | 17 | -1 | -1 | -1 | -1 | -1 | -1 | -1 | -1 | 17 |    




