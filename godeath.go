package godeath

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/gookit/color"
)

func death() {
	doc := heredoc.Doc(`


	...
   ;::::;
  ;::::; :;
;:::::'   :;
;:::::;     ;.
,:::::'       ;           OOO\
::::::;       ;          OOOOO\
;:::::;       ;         OOOOOOOO
,;::::::;     ;'         / OOOOOOO
;:::::::::'. ,,,;.        /  / DOOOOOO
.';:::::::::::::::::;,     /  /     DOOOO
,::::::;::::::;;;;::::;,   /  /        DOOO
;'::::::''::::::;;;::::: ,#/  /          DOOO
:':::::::';::::::;;::: ;::#  /            DOOO
::':::::::';:::::::: ;::::# /              DOO
':':::::::';:::::: ;::::::#/               DOO
:::':::::::';; ;:::::::::##                OO
::::':::::::';::::::::;:::#                OO
':::::'::::::::::::;'':;::#                O
':::::'::::::::;' /  / ':#
::::::':::::;'  /  /   '#  							
`)
	color.Red.Println(doc)
	color.Error.Println(`     D E A T H !   `)
	fmt.Println("\n")

}
func flower() {
	doc := heredoc.Doc(`
	/-_-\
   /  /  \
  /  /    \
  \  \    /
   \__\__/
	  \\
	  -\\   ____
		\\ /   /
 ____   \\/___/
 \   \ -//
  \___\//-
	 -//
	  \\
	  //
	 //-
   -//
   //
   \\
	\\  
   `)

	color.Green.Println(doc)

}
