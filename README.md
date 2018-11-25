# grafee
grafee (grade + feedback) is a simple application used to make feedback generation easier for teaching assistants.

grafee can produce individual grading receipts for students. These receipts contain the grading scheme, team names, group members, comments for students, and the points achieved on a per-requirement basis.

Microsoft's Excel or Apple's Numbers can be used to enter grades. Following this, assignment sheets are exported to CSV and used as input to produce the receipts.

grafee also supports multiple languages. Each language has the text strings defined in the `config.LANG.json` file. Currently, English and French configuration files are available.

## Guide

In order to run grafee you need to have installed go. You must also be able to generate the CSV files using any spreadsheet application. Personally, I use Numbers on OSX but any should work.

