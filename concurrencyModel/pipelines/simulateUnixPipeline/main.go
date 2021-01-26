package main

import "io"

type pipefunc func(in io.Reader, out io.Writer)

func bind(app func(in io.Reader, out io.Writer, args []string), args []string) pipefunc {
	return func(in io.Reader, out io.Writer) {
		app(in, out, args)
	}
}

func pipe(apps ...pipefunc) pipefunc {
	if len(apps) == 0 {return nil}
	if len(apps) == 1 {return apps[0]}

	app := apps[0]
	for i := 1; i < len(apps); i++ {
		app1, app2 := app, apps[i]

		app = func(in io.Reader, out io.Writer) {
			pr, pw := io.Pipe()
			defer pw.Close()
			go func() {
				defer pr.Close()
				app2(pr, out)
			}()
			app1(in, pw)
		}
	}
	return app
}

func main() {
	// pipe(bind(app1,args1), bind(app2, args2))
	// tar, gzip for example
	// func tar(in io.Reader, out io.Writer, files []string)
	// func gzip(in io.Reader, out io.Writer)

	pipe(bind(tar, files), gzip)(nil, out)    // 如此优雅
}
