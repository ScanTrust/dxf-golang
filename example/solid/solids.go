package main

import (
	"fmt"
	"github.com/scantrust/dxf-golang"
)

var qr = []string{
	"####### # # ##    #######",
	"#     # #     #   #     #",
	"# ### # #   ## #  # ### #",
	"# ### #  ### #    # ### #",
	"# ### # # # #   # # ### #",
	"#     #  # #  # # #     #",
	"####### # # # # # #######",
	"         #   # ##        ",
	"#  ########## # ##  # ###",
	"#   #  # #     # # ##### ",
	" #  ####   ###### ## #  #",
	"    #  #  #   # #  # ####",
	"####### #  ##  ## #     #",
	"###         ## ###  #  # ",
	"## #### ##   ### ## #####",
	"#  # # # # #  #### # ## #",
	"#  ## ## ## ######### ## ",
	"        ### ##  #   # ## ",
	"####### # ##    # # #   #",
	"#     # #    ####   #    ",
	"# ### # ## ## #######    ",
	"# ### # #  #  ##  #    ##",
	"# ### #   ## ### #  #####",
	"#     #   #   ##   ## ###",
	"####### ##  #   ##   #  #",
}

//goland:noinspection GoUnhandledErrorResult
func main() {
	d := dxf.NewDrawing()
	d.Header().LtScale = 100.0
	_, _ = d.AddLayer("QRCode", 0, dxf.DefaultLineType, true) // 0=black

	for y, row := range qr {
		for x, cell := range row {
			if cell == '#' {
				d.Rectangle(float64(x), float64(len(qr)-1-y), 1, 1)
			}
		}
	}

	err := d.SaveAs("qr.dxf")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
