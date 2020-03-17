package database

import (
	"database/sql"
	"log"
	"fmt"
         "unsafe"
)

func main() {
    // An array of contiguous uint32 values stored in memory.
    arr := []uint32{1, 2, 3}

    // The number of bytes each uint32 occupies: 4.
    const size = unsafe.Sizeof(uint32(0))

    // Take the initial memory address of the array and begin iteration.
    p := uintptr(unsafe.Pointer(&arr[0]))
    for i := 0; i < len(arr); i++ {
        // Print the integer that resides at the current address and then
        // increment the pointer to the next value in the array.
        fmt.Printf("%d ", (*(*uint32)(unsafe.Pointer(p))))
        p += size
    }
}

func Status(username string) string {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	sql := "SELECT * FROM user WHERE username='" + username + "'"
	row := db.QueryRow(sql)

	var isLogged bool
	row.Scan(&isLogged)
	if isLogged {
		return "online"
	}

	return "offline"
}


func (p TagPauseParams) Valid() bool {
	empty := 0
	if p.Hashes == "" {
		empty += 1
	}
	if p.AdTagIDs == "" {
		empty += 1
	}
	if p.Name == "" {
		empty += 1
	}
	return empty == 2
}
