package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "vRedMLauncher",
		Width:  1550,
		Height: 735,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		DisableResize:    true,
		Frameless:        true,
		CSSDragProperty:  "widows",
		CSSDragValue:     "1",
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func launchRedM() {
	fmt.Println("Starting launchRedM function")
	localAppData := os.Getenv("LOCALAPPDATA")
	redMPath := filepath.Join(localAppData, "RedM")
	redMExePath := filepath.Join(redMPath, "RedM.exe")
	bundledRedMExePath := `C:\Program Files (x86)\RWE Labs\RedM Launcher\RedM.exe`
	dataPath := filepath.Join(redMPath, "RedM Application Data", "data")

	// Vérifier si RedM.exe existe
	if _, err := os.Stat(redMExePath); os.IsNotExist(err) {
		fmt.Println("RedM.exe n'existe pas à l'emplacement :", redMExePath)
		// Copier RedM.exe s'il n'existe pas
		err := os.Rename(bundledRedMExePath, redMExePath)
		if err != nil {
			fmt.Println("Erreur lors de la copie de RedM.exe:", err)
			return
		}
	} else {
		fmt.Println("RedM.exe trouvé à l'emplacement :", redMExePath)
	}

	// Vérifier et supprimer les répertoires server-cache et server-cache-priv
	serverCachePath := filepath.Join(dataPath, "server-cache")
	serverCachePrivPath := filepath.Join(dataPath, "server-cache-priv")

	if _, err := os.Stat(serverCachePath); !os.IsNotExist(err) {
		os.RemoveAll(serverCachePath)
	}

	if _, err := os.Stat(serverCachePrivPath); !os.IsNotExist(err) {
		os.RemoveAll(serverCachePrivPath)
	}

	// Créer le raccourci sur le bureau
	shortcutPath := filepath.Join(os.Getenv("USERPROFILE"), "Desktop", "RedM - Connexion Rapide.lnk")
	powershellCommand := fmt.Sprintf(`$ws = New-Object -ComObject WScript.Shell; $s = $ws.CreateShortcut('%s'); $s.TargetPath = '%s'; $s.Arguments = '+connect rdr.lawless-street.fr'; $s.Save()`, shortcutPath, redMExePath)

	// Exécuter la commande PowerShell pour créer le raccourci
	cmd := exec.Command("powershell", "-Command", powershellCommand)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erreur lors de la création du raccourci:", err)
		return
	}

	fmt.Println("Raccourci créé avec succès sur le bureau.")

	// Lancer RedM.exe avec l'argument de connexion
	fmt.Println("Tentative de connexion à Lawless...")
	cmd = exec.Command(redMExePath, "+connect", "rdr.lawless-street.fr")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Erreur lors du lancement de RedM.exe:", err)
		return
	}

	fmt.Println("RedM.exe lancé avec succès. Connexion à Lawless en cours...")
}
