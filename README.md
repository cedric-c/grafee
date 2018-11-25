# grafee
grafee (grade + feedback) is a simple application used to make feedback generation easier for teaching assistants.

grafee can produce individual grading receipts for students. These receipts contain the grading scheme, team names, group members, comments for students, and the points achieved on a per-requirement basis.

Microsoft's Excel or Apple's Numbers can be used to enter grades. Following this, assignment sheets are exported to CSV and used as input to produce the receipts.

grafee also supports multiple languages. Each language has the text strings defined in the `config.LANG.json` file. Currently, English and French configuration files are available.

## Guide

In order to run grafee you need to have installed go. You must also be able to generate the CSV files using any spreadsheet application. Personally, I use Numbers on OSX but any should work.

1. First you want to navigate to the correct directory (this is typically your working directory if you've cloned the project)

<center>![][setup]</center>

2. Once this is done, run the command to obtain the output. In this example, our CSVs are located two directories below in `samples/SampleGradingFile`. We will be grading the contents of `Assignment1.csv`. The full command which was used in this example is `go run grafee.go -file samples/SampleGradingFile/Assignment1.csv -lang=en`.

<center>![][command]</center>

3. Students can now receive more feedback on their assignments!

<center>![][result]</center>

[setup]:guide/pre.png
[command]:guide/command.png
[result]:guide/result.png