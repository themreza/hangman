/*
 * Copyright (c) 2019 Mohammad Tomaraei.
 * All rights reserved.
 * https://tomaraei.com
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const hangman = `                                                
 _   _    _    _   _  ____ __  __    _    _   _ 
| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
|_| |_/_/   \_|_| \_|\____|_|  |_/_/   \_|_| \_|
                                                
(c) 2019 Mohammad Tomaraei  https://tomaraei.com`

var screens = [...]string{
	`   _________
    |/        
    |              
    |                
    |                 
    |               
    |                   
    |___                 
    `,
    `   _________
    |/   |      
    |              
    |                
    |                 
    |               
    |                   
    |___                 
    H`,
    `   _________       
    |/   |              
    |   (_)
    |                         
    |                       
    |                         
    |                          
    |___  
    HA`,
	`   ________               
    |/   |                   
    |   (_)                  
    |    |                     
    |    |                    
    |                           
    |                            
    |___                    
    HAN`,
    `   _________             
    |/   |               
    |   (_)                   
    |   /|                     
    |    |                    
    |                        
    |                          
    |___                          
    HANG`,
    `   _________              
    |/   |                     
    |   (_)                     
    |   /|\                    
    |    |                       
    |                             
    |                            
    |___                          
    HANGM`,
    `   ________                   
    |/   |                         
    |   (_)                      
    |   /|\                             
    |    |                          
    |   /                            
    |                                  
    |___                              
    HANGMA`,
    `   ________
    |/   |     
    |   (_)    
    |   /|\           
    |    |        
    |   / \        
    |               
    |___           
    HANGMAN`,
}

var phrases = []string{
	"abruptly",
	"absurd",
	"abyss",
	"affix",
	"askew",
	"avenue",
	"awkward",
	"axiom",
	"azure",
	"bagpipes",
	"bandwagon",
	"banjo",
	"bayou",
	"beekeeper",
	"bikini",
	"blitz",
	"blizzard",
	"boggle",
	"bookworm",
	"boxcar",
	"boxful",
	"buckaroo",
	"buffalo",
	"buffoon",
	"buxom",
	"buzzard",
	"buzzing",
	"buzzwords",
	"caliph",
	"cobweb",
	"cockiness",
	"croquet",
	"crypt",
	"curacao",
	"cycle",
	"daiquiri",
	"dirndl",
	"disavow",
	"dizzying",
	"duplex",
	"dwarves",
	"embezzle",
	"equip",
	"espionage",
	"euouae",
	"exodus",
	"faking",
	"fishhook",
	"fixable",
	"fjord",
	"flapjack",
	"flopping",
	"fluffiness",
	"flyby",
	"foxglove",
	"frazzled",
	"frizzled",
	"fuchsia",
	"funny",
	"gabby",
	"galaxy",
	"galvanize",
	"gazebo",
	"giaour",
	"gizmo",
	"glowworm",
	"glyph",
	"gnarly",
	"gnostic",
	"gossip",
	"grogginess",
	"haiku",
	"haphazard",
	"hyphen",
	"iatrogenic",
	"icebox",
	"injury",
	"ivory",
	"ivy",
	"jackpot",
	"jaundice",
	"jawbreaker",
	"jaywalk",
	"jazziest",
	"jazzy",
	"jelly",
	"jigsaw",
	"jinx",
	"jiujitsu",
	"jockey",
	"jogging",
	"joking",
	"jovial",
	"joyful",
	"juicy",
	"jukebox",
	"jumbo",
	"kayak",
	"kazoo",
	"keyhole",
	"khaki",
	"kilobyte",
	"kiosk",
	"kitsch",
	"kiwifruit",
	"klutz",
	"knapsack",
	"larynx",
	"lengths",
	"lucky",
	"luxury",
	"lymph",
	"marquis",
	"matrix",
	"megahertz",
	"microwave",
	"mnemonic",
	"mystify",
	"naphtha",
	"nightclub",
	"nowadays",
	"numbskull",
	"nymph",
	"onyx",
	"ovary",
	"oxidize",
	"oxygen",
	"pajama",
	"peekaboo",
	"phlegm",
	"pixel",
	"pizazz",
	"pneumonia",
	"polka",
	"pshaw",
	"psyche",
	"puppy",
	"puzzling",
	"quartz",
	"queue",
	"quips",
	"quixotic",
	"quiz",
	"quizzes",
	"quorum",
	"razzmatazz",
	"rhubarb",
	"rhythm",
	"rickshaw",
	"schnapps",
	"scratch",
	"shiv",
	"snazzy",
	"sphinx",
	"spritz",
	"squawk",
	"staff",
	"strength",
	"strengths",
	"stretch",
	"stronghold",
	"stymied",
	"subway",
	"swivel",
	"syndrome",
	"thriftless",
	"thumbscrew",
	"topaz",
	"transcript",
	"transgress",
	"transplant",
	"triphthong",
	"twelfth",
	"twelfths",
	"unknown",
	"unworthy",
	"unzip",
	"uptown",
	"vaporize",
	"vixen",
	"vodka",
	"voodoo",
	"vortex",
	"voyeurism",
	"walkway",
	"waltz",
	"wave",
	"wavy",
	"waxy",
	"wellspring",
	"wheezy",
	"whiskey",
	"whizzing",
	"whomever",
	"wimpy",
	"witchcraft",
	"wizard",
	"woozy",
	"wristwatch",
	"wyvern",
	"xylophone",
	"yachtsman",
	"yippee",
	"yoked",
	"youthful",
	"yummy",
	"zephyr",
	"zigzag",
	"zigzagging",
	"zilch",
	"zipper",
	"zodiac",
	"zombie",
	"Accept yourself",
	"Act justly",
	"Aim high",
	"Alive and well",
	"Amplify hope",
	"Baby steps",
	"Be awesome",
	"Be colorful",
	"Be fearless",
	"Be honest",
	"Be kind",
	"Be spontaneous",
	"Be still",
	"Beautiful chaos",
	"Breathe deeply",
	"But why",
	"Call me",
	"Carpe diem",
	"Cherish today",
	"Chill out",
	"Come back",
	"Crazy beautiful",
	"Dance today",
	"Don’t panic",
	"Don’t stop",
	"Dream big",
	"Dream bird",
	"Enjoy today",
	"two word quoteEverything counts",
	"Explore magic",
	"Fairy dust",
	"Fear not",
	"Feeling groowy",
	"Find balance",
	"Follow through",
	"For real",
	"Forever free",
	"Forget this",
	"Friends forever",
	"Game on",
	"Getting there",
	"Give thanks",
	"Good job",
	"Good vibration",
	"Got love",
	"Hakuna Matata",
	"Happy endings",
	"Have patience",
	"Hello gorgeous",
	"Hold on",
	"How lovely",
	"I can",
	"I remember",
	"I will",
	"Imperfectly perfect",
	"Infinite possibilities",
	"Inhale exhale",
	"Invite tranquility",
	"Just because",
	"Just believe",
	"Just imagine",
	"Just sayin",
	"Keep calm",
	"Keep going",
	"Keep smiling",
	"Laugh today",
	"Laughter heals",
	"Let go",
	"Limited edition",
	"Look up",
	"Look within",
	"Loosen up",
	"Love endures",
	"Love fearlessly",
	"Love you",
	"Miracle happens",
	"Move on",
	"No boundaries",
	"Not yet",
	"Notice things",
	"Oh snap",
	"Oh really",
	"Only believe",
	"Perfectly content",
	"Perfectly fabulous",
	"Pretty awesome",
	"Respect me",
	"Rise above",
	"Shift happens",
	"Shine on",
	"Sing today",
	"Slow down",
	"Start living",
	"Stay beautiful",
	"Stay focused",
	"Stay strong",
	"Stay true",
	"Stay tuned",
	"Take chances",
	"Thank you",
	"Then when",
	"Think different",
	"Think first",
	"Think twice",
	"Tickled pink",
	"Treasure today",
	"Try again",
	"Unconditional love",
	"Wanna play",
	"What if",
	"Why not",
	"Woo hoo",
	"You can",
	"You matter",
	"You sparkle",
	"Just me",
	"Be yourself",
	"Trust me",
	"Have faith",
	"Enjoy life",
	"True love",
	"Be happy",
}

const rules = `
You have 7 chances to guess the target phrase correctly.
Every time you miss, the man gets one step closer to death.
Don't let the hangman hang the man!
`

var targetPhrase []string
var guessedLetter string
var result []string
var usedLetters []string
var maxCorrectGuesses = 0
var correctGuesses = 0
var mistakes = 0
var selectedMode = 1
const maxMistakes = 7

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Println(hangman)
}

func showScreen() {
	fmt.Println(screens[mistakes])
	if(mistakes == maxMistakes) {
		fmt.Println()
		fmt.Println("Oh no, the hangman has hanged the man!")
		fmt.Println("The target phrase was:", strings.Join(targetPhrase, ""))
		fmt.Println()
	} else if(correctGuesses == maxCorrectGuesses) {
		fmt.Println()
		fmt.Println("Bravo! You saved the man from the hangman!")
		fmt.Println("The target phrase was:", strings.Join(targetPhrase, ""))
		fmt.Println()
	}
}

func showRules() {
	fmt.Println(rules)
}

func initState(tempTargetPhrase string) {
	result = make([]string, 0, len(targetPhrase))
	targetPhrase = make([]string, 0, len(targetPhrase))
	for _, c := range tempTargetPhrase {
		targetPhrase = append(targetPhrase, string(c))
		if(string(c) != " ") {
			result = append(result, "_")
			maxCorrectGuesses++
		} else {
			result = append(result, " ")
		}
	}
}

func chooseMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose single player [1] or two player [2] mode: ")
	input, _ :=  reader.ReadString('\n')
	switch mode := strings.TrimSuffix(strings.Trim(input, " "), "\n"); mode {
		case "1":
			selectedMode = 1
		case "2":
			selectedMode = 2
		default:
			fmt.Println("Wrong mode!")
			chooseMode()
	}
}

func getTargetPhrase() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the target phrase: ")
	input, _ :=  reader.ReadString('\n')
	tempTargetPhrase := strings.TrimSuffix(strings.Trim(input, " "), "\n")
	targetRegex := regexp.MustCompile("^(?:[A-Za-z]+)(?:[A-Za-z ]*)$")
	isPhraseAllowed := targetRegex.MatchString(tempTargetPhrase)
	if(!isPhraseAllowed) {
		fmt.Println("Wrong phrase. You can only use alphabets and spaces. Example: Barney Stinson")
		getTargetPhrase()
	}else{
		initState(tempTargetPhrase)
		clearScreen()
	}
}

func ChooseRandomTargetPhrase() {
	rand.Seed(time.Now().Unix())
	initState(phrases[rand.Intn(len(phrases))])
	clearScreen()
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func getGuessedLetter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your guessed letter: ")
	input, _ :=  reader.ReadString('\n')
	tempGuessedLetter := strings.ToLower(strings.TrimSuffix(strings.Trim(string(input[0]), " "), "\n"))
	guessedLetterRegex := regexp.MustCompile("^(?:[A-Za-z]+)$")
	isGuessedLetterAllowed := guessedLetterRegex.MatchString(tempGuessedLetter)
	_, usedLetterFound := Find(usedLetters, tempGuessedLetter)
	if(!isGuessedLetterAllowed) {
		fmt.Println("Wrong guessed letter. You can only use alphabets. Example: a")
		getGuessedLetter()
	} else if(usedLetterFound) {
		fmt.Println("You have already guessed the letter", tempGuessedLetter, ". Try another letter.")
		getGuessedLetter()
	} else {
		guessedLetter = tempGuessedLetter
		usedLetters = append(usedLetters, guessedLetter)
	}
}

func checkGuessedLetter() {
	correctGuess := false
	for i, c := range targetPhrase {
		if(c != " " && strings.ToLower(c) == guessedLetter) {
			result[i] = c
			correctGuesses++
			correctGuess = true
		}
	}
	if(!correctGuess) {
		mistakes++
	}
}

func showResult() {
	fmt.Println("  ", strings.Join(result, " "), "\n")
}

func main() {
	clearScreen()
	showRules()
	chooseMode()
	if(selectedMode == 1) {
		ChooseRandomTargetPhrase()
	} else {
		getTargetPhrase()
	}
	for (mistakes < maxMistakes && correctGuesses != maxCorrectGuesses) {
		clearScreen()
		showResult()
		showScreen()
		getGuessedLetter()
		checkGuessedLetter()
	}

	clearScreen()
	showScreen()
}