## Installation

### Download

```bash
go get github.com/AndreiTelteu/wails-configstore
```

### Initialize

In your `main.go` file, add the following code:

```go
import (
	// add this import if is not automatically added
	wailsconfigstore "github.com/AndreiTelteu/wails-configstore"
)

func main() {
	app := NewApp()

	// add this section
	configStore, err := wailsconfigstore.NewConfigStore("My Application Name")
	if err != nil {
		fmt.Printf("could not initialize the config store: %v\n", err)
		return
	}

	// in options > Bind add the configStore
	err = wails.Run(&options.App{
		Title:  "My Application Name",
		...
		Bind: []interface{}{
			app,
			configStore,
		},
		...
	})
}
```

## Usage in frontend

In your frontend code, you can use the config store like this:

```js
import * as ConfigStore from "../wailsjs/go/wailsconfigstore/ConfigStore";

ConfigStore.Get("auth.json", "null").then((response) => {
  const data = JSON.parse(response);
  console.log(data); // is either the data from the file or null
});

ConfigStore.Set(
  "auth.json",
  JSON.stringify({
    username: "admin",
    token: "secret",
  })
);

// For plain js you can use:
// window.go.wailsconfigstore.ConfigStore.Get(...)
```

This way you can have multiple config files for different purposes.

When you call `Set` it will create a folder with the name you provided `"My Application Name"` in the following locations:

- Windows: `C:\Users\{username}\AppData\Local\My Application Name`
- macOS: `~/Library/Application Support/My Application Name`
- Linux: `~/.config/My Application Name`

And also a file with the name you provided `"auth.json"` in the folder above.

## Usage in go

If you need to access that json file from Go, use this:

```go
// new file authState.go
package main

import (
	"encoding/json"
	"fmt"
	wailsconfigstore "github.com/AndreiTelteu/wails-configstore"
)

type AuthState struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

func GetAuthState(conf *wailsconfigstore.ConfigStore) *AuthState {
	data, err := conf.Get("servers.json", `{ "username":"", "token":"" }`)
	if err != nil {
		fmt.Println("could not read the servers config file:", err)
		return nil
	}
	var authState AuthState
	err = json.Unmarshal([]byte(data), &authState)
	if err != nil {
		fmt.Println("could not parse servers config data:", err)
		return nil
	}
	return &authState
}

// in main.go:
configStore, err := wailsconfigstore.NewConfigStore("My Application Name")
if err != nil {
	fmt.Printf("could not initialize the config store: %v\n", err)
	return
}
authState := GetAuthState(configStore)
fmt.Println("username", authState.Username)
```

### Credits

Thanks to [@ValentinTrinque](https://github.com/ValentinTrinque) for this comment: https://github.com/wailsapp/wails/issues/1956#issuecomment-1279218552
