// Copyright © 2016 Chu Duc Minh <chu.ducminh@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	hwinfo "github.com/PiScale/hwinfo-lib"
	"github.com/spf13/cobra"
)

// ramCmd represents the ram command
var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "Show information about RAM",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		ram, err := hwinfo.Get_ram()
		if err == nil {
			fmt.Println("Ram Total Capacity:", ram.TotalCapacity)
			fmt.Println("")
			for slot, _ := range ram.Items {
				fmt.Println("Ram Slot:", slot)
				fmt.Println("Ram size:", ram.Items[slot].Size)
				fmt.Println("Ram bank-location:", ram.Items[slot].BankLocator)
				fmt.Println("Ram clock-speed:", ram.Items[slot].ClockSpeed)
				fmt.Println("Ram manufacturer:", ram.Items[slot].Manufacturer)
				fmt.Println("Ram serial-number:", ram.Items[slot].SerialNumber)
				fmt.Println("Ram part-number:", ram.Items[slot].PartNumber)
				fmt.Println("")
			}
		} else {
			fmt.Println("Permission denied (you must be root)!", err)
		}
	},
}

func init() {
	showCmd.AddCommand(ramCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ramCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ramCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
