package command

func GetAppDir() (string, error) {
	return ".", nil //TODO
}

func NotInAppDirectoryMessage() string {
	helpText := `
polka expects to be run in an app directory created by 'polka new'.  E.g the following:

(1) $ pwd
~/code
(2) $ polka new todo
(3) $ cd todo
(4) $ polka [command]

You likely forgot step (3)

`
	return helpText
}
