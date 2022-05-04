package colored

type Color	string

const(
	BLACK Color = "\u001B[30m"
	RED Color = "\u001B[31m"
	GREEN Color = "\u001B[32m"
	YELLOW Color = "\u001B[33m"
	BLUE Color = "\u001B[34m"
	MAGENTA Color = "\u001B[35m"
	CYAN           Color = "\u001B[36m"
	WHITE          Color = "\u001B[37m"
	BLACK_B        Color = "\u001B[40m"
	RED_B          Color = "\u001B[41m"
	GREEN_B        Color = "\u001B[42m"
	YELLOW_B       Color = "\u001B[43m"
	BLUE_B         Color = "\u001B[44m"
	MAGENTA_B      Color = "\u001B[45m"
	CYAN_B         Color = "\u001B[46m"
	WHITE_B        Color = "\u001B[47m"
	ResetPrevColor Color = "\u001B[0m"
	BOLD           Color = "\u001B[1m"
	UNDERLINED     Color = "\u001B[4m"
	BLINK          Color = "\u001B[5m"
	REVERSED       Color = "\u001B[7m"
	INVISIBLE      Color = "\u001B[8m"
	NoColor        Color = ""
)

