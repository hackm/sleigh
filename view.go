package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sethgrid/multibar"
)

func showTextLogo() {
	logo := `
                    _...,
              o_.-"` + "`" + `    ` + "`" + `\
       .--.  _ ` + "`" + `'-._.-'""-;     _
     .'    \` + "`" + `_\_  {_.-'""-}  _ / \              888        d8b        888      
   _/     .-'  '. {c-._o_.){\|` + "`" + `  |              888        Y8P        888      
  (@` + "`" + `-._ /       \{    ^  } \\ _/               888                   888      
    ` + "`" + `~\  '-._      /'.     }  \}  .-.   .d8888b 888 .d88b. 888 .d88b. 88888b.  
     |>:<   '-.__/   '._,} \_/  / ())   88K     888d8P  Y8b888d88P"88b888 "88b 
     |     >:<   ` + "`" + `'---. ____'-.|(` + "`" + `"` + "`" + `    "Y8888b.88888888888888888  888888  888 
     \            >:<  \\_\\_\ | ;           X88888Y8b.    888Y88b 888888  888 
      \                 \\-{}-\/  \      88888P'888 "Y8888 888 "Y88888888  888 
       \                 '._\\'   /)                               888         
        '.                       /(                           Y8b d88P         
          ` + "`" + `-._ _____ _ _____ __.'\ \
            / \     / \     / \   \ \
         _.'/^\'._.'/^\'._.'/^\'.__) \
     ,=='  ` + "`" + `---` + "`" + `   '---'   '---'      )
     ` + "`" + `"""""""""""""""""""""""""""""""` + "`" + `
	`
	c := color.New(color.FgRed)
	c.Println(logo)
}

func showLogo() {
	santaClaus := `
        _
       {_}
       / \
      /   \
     /_____\
   {` + "`" + `_______` + "`" + `}
    // . . \\
   (/(__7__)\)
   |'-' = '-'|
   |         |
   /\       /\
  /  '.   .'  \
 /_/   ` + "`" + `"` + "`" + `   \_\
{__}###[_]###{__}
(_/\_________/\_)
    |___|___|
     |--|--|
    (__)` + "`" + `(__)
	`
	// continue doing other work
	fmt.Println(santaClaus)
}

// barName is ProgressBarName, progressChannel is return 0..100 percentage int channel.
func showProgress(barName string, progressChannel chan int) {
	// create the multibar container
	// this allows our bars to work together without stomping on one another
	progressBars, _ := multibar.New()
	progressBars.Println("We can separate bars with blocks of text, or have them grouped.\n")

	// we will update the progress down below in the mock work section with barProgress1(int)
	barProgress := progressBars.MakeBar(100, barName)

	// listen in for changes on the progress bars
	// I should be able to move this into the constructor at some point
	go progressBars.Listen()

	// myChからの送信を受け取る
	for progress := range progressChannel {
		barProgress(progress)
	}
	// continue doing other work
	fmt.Println("All Bars Complete")
}
