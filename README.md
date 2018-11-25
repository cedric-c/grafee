# grafee
grafee (grade + feedback) is a simple application used to make feedback generation easier for teaching assistants.

grafee can produce individual grading receipts for students. These receipts contain the grading scheme, team names, group members, comments for students, and the points achieved on a per-requirement basis.

Microsoft's Excel or Apple's Numbers can be used to enter grades. Following this, assignment sheets are exported to CSV and used as input to produce the receipts.

grafee also supports multiple languages. Each language has the text strings defined in the `config.LANG.json` file. Currently, English and French configuration files are available.

## Guide

In order to run grafee you need to have installed go. You must also be able to generate the CSV files using any spreadsheet application. Personally, I use Numbers on OSX but any should work.

1. First you want to navigate to the correct directory (this is typically your working directory if you've cloned the project)

![Setup][setup]

2. Once this is done, run the command to obtain the output. In this example, our CSVs are located two directories below in `samples/SampleGradingFile`. We will be grading the contents of `Assignment1.csv`. The full command which was used in this example is `go run grafee.go -file samples/SampleGradingFile/Assignment1.csv -lang=en`.

![Command][command]

3. Students can now receive more feedback on their assignments!

![Result][result]

[setup]: guide/pre.png
[command]: guide/command.png
[result]: guide/result.png


## Sample Output

```
Corrected by Cédric C (me@myinstitution.ca). Please contact me for all questions and inquiries.

Group name: Team404
Group members: Ad Lehman, 200810.
Final grade: (49.5)

Comments: :
You have a nice UML, everything looks nice but your APK does not compile. Upon further inspection of the source, it appears as though every class does not have actual implementations. (-10). ——— You have no unit tests (-10). ——— You submitted half a day late (-0.5).


=== Breakdown ===
(10/10) UML The student submitted a UML diagram with the required classes (-1 per missing class) (-5 for broken ISA).
(0/10) APK  The student submitted an APK which compiles. (-1 per minute of debugging required by the TA to get the APK working).
(0/15) Tests    5 Unit test cases (simple local tests). No need to include instrumentation or Espresso Tests (UI).
(19.5/20) Submission Date   The student submitted the assignment on time. (-1 per day missed).
(20/20) Required Submission The students submitted the required files. (-1 for each missing file). (-10 if the student submitted an assignment from another course). (-20 if the submission is a syllabus).

=================

Found an error in my correction? Send me an email!

Cédric C (me@myinstitution.ca)
```