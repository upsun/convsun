package logic

import (
	"log"
	"path/filepath"

	app "github.com/upsun/convsun"
	detector "github.com/upsun/lib-sun/detector"
	entity "github.com/upsun/lib-sun/entity"
	readers "github.com/upsun/lib-sun/readers"
	utils "github.com/upsun/lib-sun/utility"
	writers "github.com/upsun/lib-sun/writers"
)

func Convert(projectWorkspace string, outputFilePath string) {
	log.Print("Convert Project to Upsun...")

	// Initialize Meta-model
	var metamodel entity.MetaConfig
	absProjectWorkspace := utils.NormalizePath(projectWorkspace)
	absProjectDestination := utils.NormalizePath(outputFilePath)
	absProjectDestinationConfig := filepath.Join(absProjectDestination, "config.yaml")

	configFiles, err := detector.FindConfig(absProjectWorkspace)
	if err != nil {
		log.Fatal("Not found")
	}

	// Read PSH config files
	readers.ReadServices(&metamodel, configFiles[entity.PSH_SERVICE])
	readers.ReadPlatforms(&metamodel, configFiles[entity.PSH_PLATFORM], absProjectWorkspace)
	readers.ReadApplications(&metamodel, configFiles[entity.PSH_APPLICATION], absProjectWorkspace)
	readers.ReadRoutes(&metamodel, configFiles[entity.PSH_ROUTE])

	// Normalize (TODO)
	// - Remove Size
	log.Println("Upsun does not use 'sizes' in its configuration file (config.yml) !!\n\tSizing is defined in the web console.")
	log.Println("Remove all 'size' on " + entity.PSH_SERVICE + "...")
	readers.RemoveAllEntry(&metamodel.Services, "size")
	log.Println("Remove all 'size' on " + entity.PSH_PLATFORM + "/" + entity.PSH_APPLICATION + "...")
	readers.RemoveAllEntry(&metamodel.Applications, "size")

	// - Warning Change mount
	log.Println("Upsun uses different mount types !!\n\tFor more information: https://docs.upsun.com/create-apps/app-reference/single-runtime-image.html#define-a-mount")
	// TODO Change only on mount section
	// entry := readers.FindEntry(&metamodel.Applications, "mounts")
	// if entry.ValueNode != nil {
	log.Println("Replace all mount type on " + entity.PSH_PLATFORM + "/" + entity.PSH_APPLICATION + "...")
	readers.ReplaceAllEntry(&metamodel.Applications, "local", app.ArgsC.TypeMount)
	readers.ReplaceAllEntry(&metamodel.Applications, "shared", app.ArgsC.TypeMount)
	// }

	// - Remove Disk
	log.Println("Upsun configuration files doesn't define 'disk' !!\n\tDisk is defined into web console.\n\tFor more information: https://docs.upsun.com/create-apps/app-reference/single-runtime-image.html#available-disk-space")
	log.Println("Remove all 'disk' on " + entity.PSH_SERVICE + "...")
	readers.RemoveAllEntry(&metamodel.Services, "disk")
	log.Println("Remove all 'disk' on " + entity.PSH_PLATFORM + "/" + entity.PSH_APPLICATION + "...")
	readers.RemoveAllEntry(&metamodel.Applications, "disk")

	// - Remove Resources
	log.Println("Upsun does not use 'resources' in its configuration file (config.yml) !!\n\tResources is defined in the web console.")
	log.Println("Remove all 'resources' on " + entity.PSH_SERVICE + "...")
	readers.RemoveAllEntry(&metamodel.Services, "resources")
	log.Println("Remove all 'resources' on " + entity.PSH_PLATFORM + "/" + entity.PSH_APPLICATION + "...")
	readers.RemoveAllEntry(&metamodel.Applications, "resources")
	// - Application case

	// Generate Upsun config files
	if utils.IsExist(absProjectDestinationConfig) {
		absProjectDestinationConfig += ".new_convert"
		log.Printf("WARNING: Upsun config file already exists. Generate new one on %v", absProjectDestinationConfig)
	}
	writers.GenerateUpsunConfigFile(metamodel, absProjectDestinationConfig)

	// Move extra config
	utils.TransfertConfigCustom(absProjectWorkspace, absProjectDestination)

	// Check git:validate (core-eng)
	log.Println("Upsun configuration files generated !\n\tOPTIONAL: Please run : \"upsun app:config-validate\" in order to validate them.")
}
