package main
import (
    "fmt"
    "encoding/csv"
    "encoding/json"
    "os"
    "io"
    "io/ioutil"
    "bufio"
    "strconv"
    "bytes"
    "flag"
)
/*
Requirement (s) are parsed from the columns starting from 2 (i=0)  and ending at n-2 (i=0).
*/
type Requirement struct {
    Name string
    Description string
    Weight float64
}


type Deliverable struct {
    Requirements[] Requirement
    Title string
}

type Group struct {
    Name string
    Members string
    Grades []string
    Comments string
    Mark string
}

type Corrector struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

type Users struct {
    Correctors []Corrector `json:"corrector"`
}

type Messages struct {
    Messages []Message `json:"content"`
}

type Message struct {
    By string `json:"by"`
    Contact string `json:"contact"`
    Error string `json:"error"`
    GroupName string `json:"group"`
    Members string `json:"members"`
    Grade string `json:"grade"`
    Comments string `json:"comments"`
    Breakdown1 string `json:"breakdown1"`
    Breakdown2 string `json:"breakdown2"`
}

// limit is the last value in the requirements struct
func initRequirements(file string, limit int) []Requirement{
    var requirements []Requirement
    f, _ := os.Open(file)
    reader := csv.NewReader(bufio.NewReader(f))
    
    var names, descriptions []string
    var weights []float64
    
    
    recordLength := 0
    for lineIndex := 0; lineIndex < limit; lineIndex++ {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {}
        
        recordLength = 0
        for _, value := range line {
                if lineIndex == 0 {
                    names = append(names, value)
                } else if lineIndex == 1 {
                    descriptions = append(descriptions, value)
                } else if lineIndex == 2 {
                    v, _ := strconv.ParseFloat(value, 64)
                    weights = append(weights, v)
                } else {}
            recordLength++
        }
    }
    
    // First values are simply Feature, Description, and Weight, remove
    names = names[1:len(names)]
    descriptions = descriptions[1:len(descriptions)]
    weights = weights[1:len(weights)]
    
    // Create Requirements
    for i :=0; i < recordLength-1; i++ {
        requirements = append(requirements, Requirement{
            Name: names[i],
            Description: descriptions[i],
            Weight: weights[i],
            })
    }
    return requirements[1:len(requirements)]
}

func initGroups(file string, limit int) []Group{
    f, _ := os.Open(file)
    r := csv.NewReader(f)
    lines, _ := r.ReadAll()
    var groups []Group
    
    // start from line 3, since headers stop on line 2
    for _, value := range lines[limit:len(lines)]{
        
        markIndex := len(value) - 1
        commentIndex := len(value) - 2
        lastGradeIndex := len(value) - 4
        
        groupeName := value[0]
        members := value[1]
        grades := value[2:lastGradeIndex]
        comments := value[commentIndex]
        finalMark := value[markIndex] // should be sum of grades
        
        groups = append(groups, Group{
            Name: groupeName,
            Members: members,
            Grades: grades,
            Comments: comments,
            Mark: finalMark,
            })
        
    }
    
    return groups

}

func main(){
    
    filePtrFlag := flag.String("file", "", "A relative path to the file on which to generate the receipts.")
    langPtrFlag := flag.String("lang", "en","The language with which to generate the receipts. Available options are (en/fr).")
    flag.Parse()
    
    requirements := initRequirements(*filePtrFlag, 3)
    groups := initGroups(*filePtrFlag, 3)
    var fileName string
    if(*langPtrFlag == "en"){
        fileName = "config.en.json"
    } else if (*langPtrFlag == "fr") {
        fileName = "config.fr.json"
    } else {
        fmt.Println("Invalid language. Choices are (en/fr). Type grafee.go -h for help.")
        os.Exit(1)
    }
    
    
    prefs, _ := os.Open(fileName)
    defer prefs.Close()
    
    byteVals, _ := ioutil.ReadAll(prefs)

    var corrector Users
    var messages Messages
    
    json.Unmarshal(byteVals, &corrector)
    json.Unmarshal(byteVals, &messages)

    var buf bytes.Buffer
    for _, group := range groups {
                
        msg := messages.Messages[0]
        ta  := corrector.Correctors[0]
        
        notice1 := fmt.Sprintf("%v %v (%v). %v\n\n",msg.By,
            ta.Name,
            ta.Email,
            msg.Contact)
        notice2 := fmt.Sprintf("%v\n\n%v (%v)",
            msg.Error,
            ta.Name,
            ta.Email)
        
        filename := fmt.Sprintf("%v.txt", group.Name)
        
        buf.WriteString(fmt.Sprintf(notice1))
        buf.WriteString(fmt.Sprintf("%v: %v\n",msg.GroupName, group.Name))
        buf.WriteString(fmt.Sprintf("%v: %v\n",msg.Members, group.Members))
        buf.WriteString(fmt.Sprintf("%v: (%v)\n",msg.Grade, group.Mark))
        buf.WriteString(fmt.Sprintf("\n%v:\n%v\n\n",msg.Comments, group.Comments))
        
        
        buf.WriteString(fmt.Sprintf("\n%v\n", msg.Breakdown1))
        for i, grade := range group.Grades {
            buf.WriteString(fmt.Sprintf("(%+v/%+v) ", grade, requirements[i].Weight))
            buf.WriteString(fmt.Sprintf("%v\t", requirements[i].Name))
            buf.WriteString(fmt.Sprintf("%v\n", requirements[i].Description))
        }
        buf.WriteString(fmt.Sprintf("\n%v\n\n", msg.Breakdown2))
        buf.WriteString(fmt.Sprintf(notice2))

        f,_ := os.Create(filename)
        f.WriteString(buf.String())
        buf.Reset()
    }
    
}