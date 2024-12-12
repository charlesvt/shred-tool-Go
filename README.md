# Shred Tool (Go)
## Description
Will remove a given file (by default, "randomfile.txt") within the same directory [main.go](https://github.com/charlesvt/shred-tool-Go/blob/main/main.go) is located in.
## Dependencies
Please ensure Go is installed on the device and is added to PATH. Please follow these [instructions](https://go.dev/doc/install) to do so.
## Features
Due to the nature of file deletion, there is no guarantee whether the deleted data is protected from others. Thus, as a precaution, the tool will overwrite a file with data at least three times to ensure the full erasure of original data before deleting the file.
1.  Detects if file exists or not
2.  Determines original file size
3.  Overwrites file with randomized data of the same of the same size as the original file
4. Deletes file once it is overwritten  with random data

## Using tool
To use the tool, please place the script within the same directory of the file that will be shredded. The target file can be changed by updating filePath in the main function. By default, the file path is set to "randomfile.txt". Running ```go run main.go``` should result in the overwrite and deletion of the given file.

## Additional functions
Additional functions are included in the script as debug functions.
- createRandFile(path)
    - Creates a text file by passing the file name as a string through "path"
    - Fills the text file with some filler text as well
- randomDataGen(path)
    - Detects the size of a given file by passing the file name as a string through "path"
    - Overwrites the data of the given file three times
- shredOnly(path)
    - Deletes given file by passing the file name as a string through "path"
