# Fortune Receiver

Script for displaying random fortune for conky app.

## Example usage:

1. Download the latest script from releases block.
2. Create crontab entry `*/15 * * * * ~/fortune-receiver > ~/random.log`
3. You can also use type console attribute from here: http://rzhunemogu.ru/FAQ.aspx
4. Add rule for conky:
```bash
${offset 0}${font Noto Sans:size=10}${execi 500 cat ~/random.log}
```

You can pass length string param in command line

## Splitting long string to multiline
In this script implemented splitString function, that splits a long multiline string into multiple lines by a maximum string length.

Here's how the updated function works:
1. It takes two parameters: the input string and the maximum length of a line.
2. It initializes an empty string called output to store the result.
3. It splits the input string into an array of strings by the newline character using the strings.Split function.
4. It loops through each line of the input string.
5. For each line, it splits the line into an array of words using the space character as the separator.
6. It initializes a variable lineLength to keep track of the length of the current line.
7. It loops through each word in the line.
8. For each word, it checks if adding it to the current line would exceed the maximum length. If so, it adds a newline character and the word to the output string, and sets lineLength to the length of the word plus one. Otherwise, it adds the word to the output string with a space character, and adds the length of the word plus one to lineLength.
9. After processing all the words in the line, it adds a newline character to the output string.
10. Finally, it returns the output string.
