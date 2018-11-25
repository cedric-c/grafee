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

type Messages struct {
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
    
    cliArgs := os.Args[1:]
    // jsonPreferences, err := os.Open("config.json")
    file := cliArgs[0]
    fmt.Println(file)
    
    requirements := initRequirements(file, 3)
    groups := initGroups(file, 3)

    prefs, _ := os.Open("config.json")
    defer prefs.Close()
    fmt.Printf("%+v", prefs)
    
    
    byteVals, _ := ioutil.ReadAll(prefs)
    var results map[string]interface{}
    
    // fmt.Printf("%+v", byteVals)

    // var corrector Corrector
    // var messages Messages
    
    json.Unmarshal([]byte(byteVals), &results)
    fmt.Printf("%+v", results["user"])
    // json.Unmarshal([]byte(byteVals), &messages)
    // fmt.Println(corrector.Name)
    // fmt.Println(messages.By)

    
    
    var buf bytes.Buffer
    for _, group := range groups {
        
        notice1 := "Correction par Cédric (cclem054@uottawa.ca). SVP, me contacter pour poser vos questions.\n\n"
        notice2 := "Vous voyez une erreur dans la correction? Envoyez-moi vos questions!\n\nCédric (cclem054@uottawa.ca)"
        
        
        
        filename := fmt.Sprintf("%v.txt", group.Name)
        
        buf.WriteString(fmt.Sprintf(notice1))
        buf.WriteString(fmt.Sprintf("Nom du groupe: %v\n", group.Name))
        buf.WriteString(fmt.Sprintf("Membres du groupe: %v\n", group.Members))
        buf.WriteString(fmt.Sprintf("Note finale: (%v)\n", group.Mark))
        buf.WriteString(fmt.Sprintf("\nCommentaires:\n%v\n\n", group.Comments))
        
        
        buf.WriteString(fmt.Sprintf("\n=== Breakdown ===\n"))
        for i, grade := range group.Grades {
            buf.WriteString(fmt.Sprintf("(%+v/%+v) ", grade, requirements[i].Weight))
            buf.WriteString(fmt.Sprintf("%v\t", requirements[i].Name))
            buf.WriteString(fmt.Sprintf("%v\n", requirements[i].Description))
        }
        buf.WriteString(fmt.Sprintf("\n=================\n\n"))
        buf.WriteString(fmt.Sprintf(notice2))

        f,_ := os.Create(filename)
        f.WriteString(buf.String())
        buf.Reset()
    }
    
}