/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"encoding/json"

	"github.com/google/uuid"
)

// genDataCmd represents the genData command
var genDataCmd = &cobra.Command{
	Use:   "gen-data",
	Short: "Generate test data",
	Long:  `Generate test data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen-data called")

		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			fmt.Println("Error getting size flag:", err)
			return
		}
		fmt.Println("Generating test data of size", size)

		const dataDir = ".tmp"

		err = createDirIfNotExist(dataDir)
		if err != nil {
			fmt.Println("Error creating data directory:", err)
			return
		}

		const outfile = "testdata.json"

		outpath := filepath.Join(dataDir, outfile)
		outpath = strings.Replace(outpath, ".json", "."+strconv.Itoa(size)+".json", 1)

		err = genData(size, outpath)
		if err != nil {
			fmt.Println("Error generating test data:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(genDataCmd)

	// Here you will define your flags and configuration settings.

	// Add an integer flag called "size" with a default value of 10 and a short description
	genDataCmd.Flags().IntP("size", "s", 10, "Size of test data to generate")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genData(size int, outfile string) error {
	// Create a slice to hold the generated items
	items := make([]map[string]string, size)

	// Generate UUID strings for each item
	for i := 0; i < size; i++ {
		id := uuid.New().String()
		value := uuid.New().String()

		// Create a map to hold the item data
		item := map[string]string{
			"id":    id,
			"value": value,
		}

		// Add the item to the slice
		items[i] = item
	}

	// Create a map to hold the top-level data
	data := map[string]interface{}{
		"items": items,
	}

	// Convert the data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON data: %v", err)
	}

	// Write the JSON data to the output file
	err = os.WriteFile(outfile, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON data to file: %v", err)
	}

	fmt.Printf("Generated test data of size %d and wrote it to file %s\n", size, outfile)
	return nil
}

func createDirIfNotExist(dirname string) error {
	// Check if the directory already exists
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		// Create the directory if it does not exist
		err = os.Mkdir(dirname, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
		fmt.Printf("Created directory %s\n", dirname)
	} else if err != nil {
		return fmt.Errorf("error checking directory: %v", err)
	}
	return nil
}
