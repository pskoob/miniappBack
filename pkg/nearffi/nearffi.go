package nearffi

/*
#cgo LDFLAGS: -L${SRCDIR} -lnear_bridge
#include <stdlib.h>
char* send_ft(const char* from, const char* to, const char* amount, const char* secretKey, const char* tokenContract);
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type FtResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendFT(from, to, amount, secretKey, tokenContract string) (*FtResponse, error) {
	cFrom := C.CString(from)
	cTo := C.CString(to)
	cAmount := C.CString(amount)
	cSecretKey := C.CString(secretKey)
	cTokenContract := C.CString(tokenContract)

	defer C.free(unsafe.Pointer(cFrom))
	defer C.free(unsafe.Pointer(cTo))
	defer C.free(unsafe.Pointer(cAmount))
	defer C.free(unsafe.Pointer(cSecretKey))
	defer C.free(unsafe.Pointer(cTokenContract))

	result := C.send_ft(cFrom, cTo, cAmount, cSecretKey, cTokenContract)
	defer C.free(unsafe.Pointer(result))

	resStr := C.GoString(result)

	var resp FtResponse
	if err := json.Unmarshal([]byte(resStr), &resp); err != nil {
		return nil, err
	}

	if resp.Status != "success" {
		return &resp, fmt.Errorf(resp.Message)
	}

	return &resp, nil
}
