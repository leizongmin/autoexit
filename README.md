# autoexit

Execute command and automatic exit after specified seconds

## Compile

```bash
go get -u github.com/leizongmin/autoexit
```

## Usage

```text
autoexit v1.0 by Zongmin Lei <leizongmin@gmail.com> Copyright 2018
Project: https://github.com/leizongmin/autoexit
Usage:   autoexit <seconds> <command>
Example:
         autoexit 3600 cmd arg1 arg2 ...           # auto exit after 3600 seconds
         autoexit 3600s cmd arg1 arg2 ...          # auto exit after 3600 seconds
         autoexit 2d+3h+4m+5s cmd arg1 arg2 ...    # auto exit after 3 days and 3 hours and 4 minutes and 5 seconds

```

## License

```text
MIT License

Copyright (c) 2018 Zongmin Lei <leizongmin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
