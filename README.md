
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
	"github.com/AndreiTelteu/wails-configstore"
)

func main() {
	app := NewApp()

	// add this section
	configStore, err := NewConfigStore("My Application Name")
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

## Usage

In your frontend code, you can use the config store like this:

```js
ConfigStore.Get('auth.json', 'null').then(res => {
	data = JSON.parse(res);
	console.log(data); // is either the data from the file or null
});
ConfigStore.Set('auth.json', JSON.stringify({
	username: 'admin', token: 'secret'
});
```

This way you can have multiple config files for different purposes.

### Credits

Thanks to [@ValentinTrinque](https://github.com/ValentinTrinque) for this comment: https://github.com/wailsapp/wails/issues/1956#issuecomment-1279218552