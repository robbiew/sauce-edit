package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	sauce "github.com/ActiveState/go-ansi"
)

const (
	Esc = "\u001B["
	Osc = "\u001B]"
	Bel = "\u0007"
)

var (
	hasGroup  bool
	hasAuthor bool
	hasTitle  bool
	sGroup    string
	sAuthor   string
	sTitle    string
	artPath   string
	justSAUCE bool

	re = regexp.MustCompile(ansi)

	Black     = Esc + "30m"
	Red       = Esc + "31m"
	Green     = Esc + "32m"
	Yellow    = Esc + "33m"
	Blue      = Esc + "34m"
	Magenta   = Esc + "35m"
	Cyan      = Esc + "36m"
	White     = Esc + "37m"
	BlackHi   = Esc + "30;1m"
	RedHi     = Esc + "31;1m"
	GreenHi   = Esc + "32;1m"
	YellowHi  = Esc + "33;1m"
	BlueHi    = Esc + "34;1m"
	MagentaHi = Esc + "35;1m"
	CyanHi    = Esc + "36;1m"
	WhiteHi   = Esc + "37;1m"

	BgBlack     = Esc + "40m"
	BgRed       = Esc + "41m"
	BgGreen     = Esc + "42m"
	BgYellow    = Esc + "43m"
	BgBlue      = Esc + "44m"
	BgMagenta   = Esc + "45m"
	BgCyan      = Esc + "46m"
	BgWhite     = Esc + "47m"
	BgBlackHi   = Esc + "40;1m"
	BgRedHi     = Esc + "41;1m"
	BgGreenHi   = Esc + "42;1m"
	BgYellowHi  = Esc + "43;1m"
	BgBlueHi    = Esc + "44;1m"
	BgMagentaHi = Esc + "45;1m"
	BgCyanHi    = Esc + "46;1m"
	BgWhiteHi   = Esc + "47;1m"

	Reset = Esc + "0m"
)

// Sanitize input just in case
const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

func main() {

	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	hasTitle = false
	hasAuthor = false
	hasGroup = false
	justSAUCE = false

	// Use FLAG to get command line paramenters
	pathPtr := flag.String("path", "", "path to ANSI file with SAUCE")
	titlePtr := flag.String("title", "", "'New SAUCE Title' (with SIGNLE QUOTES)")
	authorPtr := flag.String("author", "", "'New SAUCE Author' (with SIGNLE QUOTES)")
	groupPtr := flag.String("group", "", "'New SAUCE Group' with (SIGNLE QUOTES)")
	required := []string{"path"}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, Red+"\nerror: -%s argumant is reuired! Exiting..."+Reset+"\n\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}
	artPath = dir + "/" + *pathPtr

	sTitle = removeAnsi(string(*titlePtr))
	sTitle = strings.Replace(sTitle, "%", "%%", -1)

	if len(strings.TrimSpace(sTitle)) != 0 {
		hasTitle = true
	}

	sAuthor = *authorPtr
	if len(strings.TrimSpace(sAuthor)) != 0 {
		hasAuthor = true
	}

	sGroup = *groupPtr
	if len(strings.TrimSpace(sGroup)) != 0 {
		hasGroup = true
	}

	if !hasTitle && !hasAuthor && !hasGroup {
		justSAUCE = true
	}

	// let's check the file for a valid SAUCE record
	record := sauce.GetSauce(artPath)

	// abort if we don't find SAUCE
	if string(record.Sauceinf.ID[:]) == sauce.SauceID {

		fmt.Printf("\n")
		fmt.Printf(BgBlue + White + "     Edit SAUCE v.01 made by aLPHA      " + Reset + "\n")

		fmt.Printf(CyanHi + "----------------------------------------\n" + Reset)
	} else {
		fmt.Println("NO SAUCE RECORD FOUND!")
		os.Exit(0)
	}

	f, err := os.OpenFile(artPath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if justSAUCE {
		showSauce(record)
	} else {

		var offset int64 = -128 // from end of file
		var whence int = 2

		if hasTitle {
			var position int64 = 7
			f.Seek(offset+position, whence)

			v := sTitle

			vLen := len(v)
			if vLen <= 35 {
				sp := 35 - vLen
				data := []byte(fmt.Sprintf(v+"%s", strings.Repeat(" ", sp)))
				_, err2 := f.Write(data)
				if err2 != nil {
					log.Fatal("Error!!! ", err2)
				}
			} else {
				fmt.Println(Red + "Too many characters in title - max 35 - exiting..." + Reset + "\n")
				os.Exit(0)
			}
			f, err := os.Open(artPath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			record := sauce.GetSauce(artPath)
			fmt.Printf(Green+"New Title    : "+GreenHi+"%s"+Reset+"\n", record.Sauceinf.Title)
		}

		if hasAuthor {
			var position int64 = 42
			f.Seek(offset+position, whence)
			v := sAuthor
			vLen := len(v)
			if vLen <= 20 {
				sp := 20 - vLen
				data := []byte(fmt.Sprintf(v+"%s", strings.Repeat(" ", sp)))
				_, err2 := f.Write(data)
				if err2 != nil {
					log.Fatal("Error!!! ", err2)
				}
			} else {
				fmt.Println(Red + "Too many characters in Author - max 20 - exiting..." + Reset + "\n")
				os.Exit(0)
			}
			f, err := os.Open(artPath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			record := sauce.GetSauce(artPath)
			fmt.Printf(Green+"New Author   : "+GreenHi+"%s"+Reset+"\n", record.Sauceinf.Author)

		}
		if hasGroup {
			var position int64 = 62
			f.Seek(offset+position, whence)
			v := sGroup
			vLen := len(v)
			if vLen <= 20 {
				sp := 20 - vLen
				data := []byte(fmt.Sprintf(v+"%s", strings.Repeat(" ", sp)))
				_, err2 := f.Write(data)
				if err2 != nil {
					log.Fatal("Error!!! ", err2)
				}
			} else {
				fmt.Println(Red + "Too many characters in Group - max 20 - exiting..." + Reset + "\n")
				os.Exit(0)
			}
			f, err := os.Open(artPath)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			record := sauce.GetSauce(artPath)
			fmt.Printf(Green+"New Group    : "+GreenHi+"%s"+Reset+"\n", record.Sauceinf.Group)

		}
		fmt.Printf("\r" + CyanHi + "----------------------------------------" + Reset + "\n")
		fmt.Printf(Red+"%s", artPath+RedHi+" // SAUCE SAVED!\n\n")
	}

}

func removeAnsi(str string) string {
	stringB := strings.ToValidUTF8(str, "\b")
	return re.ReplaceAllString(stringB, "")
}

func showSauce(record sauce.Sauce) {
	// fmt.Printf("Id       : %s v%s\n", record.Sauceinf.ID, record.Sauceinf.Version)
	fmt.Printf(Green+"Title    : "+White+"%s"+Reset+"\n", record.Sauceinf.Title)
	fmt.Printf(Green+"Author   : "+White+"%s"+Reset+"\n", record.Sauceinf.Author)
	fmt.Printf(Green+"Group    : "+White+"%s"+Reset+"\n", record.Sauceinf.Group)
	// fmt.Printf("Date     : %s\n", record.Sauceinf.Date)
	// fmt.Printf("Datatype : %d\n", record.Sauceinf.DataType)
	// fmt.Printf("Filetype : %d\n", record.Sauceinf.FileType)
	// if record.Sauceinf.Flags != 0 {
	// 	fmt.Printf("Flags    : %d\n", record.Sauceinf.Flags)
	// }
	// if record.Sauceinf.Tinfo1 != 0 {
	// 	fmt.Printf("Tinfo1   : %d\n", record.Sauceinf.Tinfo1)
	// }
	// if record.Sauceinf.Tinfo2 != 0 {
	// 	fmt.Printf("Tinfo2   : %d\n", record.Sauceinf.Tinfo2)
	// }
	// if record.Sauceinf.Tinfo3 != 0 {
	// 	fmt.Printf("Tinfo3   : %d\n", record.Sauceinf.Tinfo3)
	// }
	// if record.Sauceinf.Tinfo4 != 0 {
	// 	fmt.Printf("Tinfo4   : %d\n", record.Sauceinf.Tinfo4)
	// }

	fmt.Printf(Green+"# Cmts   : "+White+"%d"+Reset+"\n", record.Sauceinf.Comments)
	if record.Sauceinf.Comments > 0 && len(record.CommentLines) > 0 {
		fmt.Printf("Cmts   : ")
		for i := 0; i < int(record.Sauceinf.Comments); i++ {
			fmt.Printf("%s\n", record.CommentLines[i])
		}
	}
	fmt.Printf(CyanHi + "\r----------------------------------------\n" + Reset)
	fmt.Printf(Red+"%s", artPath+RedHi+" // NO CHANGES...\n\n")
}
