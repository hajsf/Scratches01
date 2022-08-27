package main

import (
	"github.com/matiasinsaurralde/go-dotnet"

	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi, I'll initialize the .NET runtime.")

	/*
		If you don't set the TRUSTED_PLATFORM_ASSEMBLIES, it will use the default tpaList value.
		APP_PATHS & NATIVE_DLL_SEARCH_DIRECTORIES use the path of the current program,
		this makes it easier to load an assembly, just put the DLL in the same folder as your Go binary!
		You're free to override them to fit your needs.
	*/

	properties := map[string]string{
		// "TRUSTED_PLATFORM_ASSEMBLIES": "/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/mscorlib.ni.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Private.CoreLib.ni.dll",
		// "APP_PATHS":                     "/Users/matias/dev/dotnet/lib/HelloWorld",
		// "NATIVE_DLL_SEARCH_DIRECTORIES": "/Users/matias/dev/dotnet/lib/HelloWorld",
	}

	/*
		CLRFilesAbsolutePath sets the SDK path.
		In case you don't set this parameter, this package will try to find the SDK using a list of common paths.
		It seems to find the right paths under Linux & OSX, feel free to override this setting (like the commented line).
	*/

	runtime, err := dotnet.NewRuntime(dotnet.RuntimeParams{
		Properties: properties,
		// CLRFilesAbsolutePath: "/usr/share/dotnet/shared/Microsoft.NETCore.App/1.0.0"
	})
	defer runtime.Shutdown()

	if err != nil {
		fmt.Println("Something bad happened! :(")
		os.Exit(1)
	}

	fmt.Println("Runtime loaded.")

	SayHello := runtime.CreateDelegate("HelloWorld", "HelloWorld.HelloWorld", "Hello")

	// this will call HelloWorld.HelloWorld.Hello :)
	SayHello()
}
