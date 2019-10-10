package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const offset = 56
const width = 38
const tophonk = `
                                                          _________________________________________
					                 /                                         \`
const bottomhonk = `
                                                    /   \__________________________________________/
			                          / 
                               ,-""   .         /
                             ,'  _   e )'-._  / 
                            /  ,' '-._<.===-'
                           /  /
                          /  ;
              _          /   ;
 ('._    _.-"" ""--..__,'    |
 <_  '-""                     \
  <'-                          :
   (__   <__.                  ;
     '-.   '-.__.      _.'    /
        \      '-.__,-'    _,'
         '._    ,    /__,-'
            ""._\__,'< <____
                 | |  '----.'.
                 | |        \ '.
                 ; |___      \-''
                 \   --<
                  '.'.<
             hjw    '-'
`

func main() {

	// Width = 40
	// offset = 42
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	prefix := strings.Repeat(" ", offset) + "|"

	fmt.Println(tophonk)
	buffer := strings.Repeat(" ", width/2-len(text)/2)
	fmt.Println(prefix + buffer + text + buffer + strings.Repeat(" ", (len(text)/2)+1) + "|")
	fmt.Println(bottomhonk)
}
