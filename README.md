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

