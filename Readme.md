# PDP Parser

Easily parse footsites PDP page.

### Example

```go
    package main

    import (
    	"github.com/Cutty72/Footsites-PDP-Parser"
    	"fmt"
    )

    func main() {
    	// site, sku, stockStatus ("inStock" or "all")
    	sizes, err := Parse("footlocker.com", "W2288111", "inStock")
    	if err != nil {
    		fmt.Println(err)
    	} 
    	// to grab a specific size we would do 
    	// fmt.Println(sizes["08.0"] etc 
    	fmt.Println(sizes) 
    }
```


### Credits
[Cutty72](https://github.com/Cutty72) ([@72_cutty](https://twitter.com/72_cutty))
<br>

## License
MIT License

Copyright (c) 2021 Cutty72

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
