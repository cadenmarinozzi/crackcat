package errs

type PathError struct {
	Path string
}

func (err *PathError) Error() string { 
	return "Could not find " + err.Path + ". If crackcat is running from the system %PATH%, make sure any file paths are not relative.";
}  