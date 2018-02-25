:: COMPILE RUN BATCH
:: -> type "cr" in cmd to compile and run with default configuration

:: Build
::
:: Note: To Suppress console in windows use:
:: go build -o gol.exe -ldflags -H=windowsgui
go build -o gol.exe

:: Launch
gol -fps=10 -seeder=random.default -density=0.1 -debug
:: Alternative: gol -seeder=random.gliders -debug

::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:: Valid Params and defaults for launch (type "gol -h" to view):
::   -debug
::         enable debug logs
::   -density float
::         seed density, if applicable (default 0.1)
::   -fps int
::         frames per second (default 20)
::   -height int
::         height of the panel (default 600)
::   -scale int
::         pixel scaling (default 4)
::   -seeder string
::         seed method (default "random.default")
::   -width int
::         width of the panel (default 600)
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
