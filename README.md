# 1pwsafe

Convert [pwsafe][] to [1password][].

[pwsafe]: http://pwsafe.info/
[1password]: https://agilebits.com/onepassword


## Usage

Grab the latest [release][]. Then, export your pwsafe; this produces a CSV file. Then:

```console
1pwsafe < path/to/pwsafe-export.csv > 1password-import.csv
```

Load `1password-import.csv` in your 1password install. Presto.

[release]: https://github.com/whilp/1pwsafe/releases/latest
