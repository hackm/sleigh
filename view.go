package main

// import (
// 	"fmt"

// 	"github.com/sethgrid/multibar"
// )

// func showLogo() {
// 	santaClaus := `
//         _
//        {_}
//        / \
//       /   \
//      /_____\
//    {` + "`" + `_______` + "`" + `}
//     // . . \\
//    (/(__7__)\)
//    |'-' = '-'|
//    |         |
//    /\       /\
//   /  '.   .'  \
//  /_/   ` + "`" + `"` + "`" + `   \_\
// {__}###[_]###{__}
// (_/\_________/\_)
//     |___|___|
//      |--|--|
//     (__)` + "`" + `(__)
// 	`
// 	// continue doing other work
// 	fmt.Println(santaClaus)
// }

// // barName is ProgressBarName, progressChannel is return 0..100 percentage int channel.
// func showProgress(barName string, progressChannel chan int) {
// 	// create the multibar container
// 	// this allows our bars to work together without stomping on one another
// 	progressBars, _ := multibar.New()
// 	progressBars.Println("We can separate bars with blocks of text, or have them grouped.\n")

// 	// we will update the progress down below in the mock work section with barProgress1(int)
// 	barProgress := progressBars.MakeBar(100, barName)

// 	// listen in for changes on the progress bars
// 	// I should be able to move this into the constructor at some point
// 	go progressBars.Listen()

// 	// myChからの送信を受け取る
// 	for progress := range progressChannel {
// 		barProgress(progress)
// 	}
// 	// continue doing other work
// 	fmt.Println("All Bars Complete")
// }
