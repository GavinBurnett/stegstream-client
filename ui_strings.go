package main

const (
	UI_InvalidArgs                   = `ERROR: Invalid arguments`
	UI_FileNotFound                  = `ERROR: File not found: `
	UI_NoParametersGiven             = `ERROR: No parameters specified`
	UI_ParameterInvalid              = `ERROR: Invalid parameter: %s`
	UI_NoBufferMemory                = `ERROR: Failed to allocate buffer memory`
	UI_HiddenFileNotCreated          = `ERROR: Hidden file: %s not created. %s`
	UI_FileReadError                 = `ERROR: File: %s read error %s`
	UI_FileWriteError                = `ERROR: File: %s write error %s`
	UI_FileExistsError               = `ERROR: File already exists: `
	UI_FileOpenError                 = `ERROR: File open error: `
	UI_FileCreateError               = `ERROR: File create error: `
	UI_FileCopyError                 = `ERROR: File copy error: `
	UI_DeleteError                   = `ERROR: File delete error: `
	UI_HiddenFileDataNotFound        = `ERROR: Hidden file data not found`
	UI_HiddenFileNameNotFound        = `ERROR: Hidden file name not found`
	UI_ReadMagicNumberFailed         = `ERROR: Read magic number failed: `
	UI_ReadFileNameFailed            = `ERROR: Read file name failed: `
	UI_ReadStepsFailed               = `ERROR: Read steps failed: `
	UI_ReadSpacingFailed             = `ERROR: Read spacing failed: `
	UI_ServerError                   = `ERROR: Server error: `
	UI_FailedToExtractHiddenFile     = `ERROR: Failed to extract hidden file`
	UI_InvalidURL                    = `ERROR: Invalid URL`
	UI_FailedToDownloadContainerFile = `ERROR: Failed to download container file`
	UI_EmptyFile                     = `ERROR: Empty file`
	UI_InvalidFileSize               = `ERROR: Invalid file size: `
	UI_FileTooBig                    = `ERROR: File too big`
	UI_NoFileSize                    = `ERROR: Can't get file size: %s %s`
	UI_FileDeleteError               = `ERROR: File delete error: `
	UI_RandomDataError               = `ERROR: Random data not generated`
	UI_WipeFileError                 = `ERROR: Failed to wipe file: %s`
	UI_FileSize                      = `%s File size: %v`
	UI_BytesWritten                  = `Bytes written: %v`
	UI_DataWriteComplete             = `Data write complete`
	UI_EOF                           = `EOF: %s`
	UI_Arguments                     = `Arguments: `
	UI_FileFound                     = `File found: `
	UI_FileDeleted                   = `File deleted: `
	UI_MagicNumber                   = "Magic number: "
	UI_MagicNumberFound              = "Magic number found"
	UI_MagicNumberNotFound           = "Magic number not found"
	UI_WipingFile                    = `Wiping file: %s `
	UI_WipedFile                     = `Wiped file: %s`
	UI_BufferSize                    = `Buffer size: %v`
	UI_Parameter                     = `Parameters: %s`
	UI_GettingFile                   = `Getting file from server`
	UI_GotFile                       = `Got file from server`
	UI_ExtractingHiddenFile          = `Extracting hidden file`
	UI_ExtractedHiddenFile           = `Extracted hidden file:`
	UI_CreatingFile                  = `Creating file: `
	UI_Done                          = ` - Done.`
	UI_LocalFile                     = `--local`
	UI_HiddenFileData                = `Hidden file data:
Spacing: %v
Steps: %v
File name: %s`
	UI_Help = `stegstream client v1.3 by gburnett@outlook.com

Arguments:

./stegstream-client <server URL>
./stegstream-client --local <local file>

Examples:

./stegstream-client http://localhost:8080/Audio
./stegstream-client --local Waves.mp3`
)
