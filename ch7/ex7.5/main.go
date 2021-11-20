package main

import "io"

// type ByteLimitReader struct {
// 	reader io.Reader 
// 	limit int64
// }

// func (blr *ByteLimitReader) Read(p []byte)(int, )

func LimitReader(r io.Reader, n int64) io.Reader {
	Read(p []byte) (int, error){
		// length of p
		nToRead := int64(len(p))
		// if too long set length at n
		if nToRead > n {
			nToRead = n
		}
		nRead, err := r.Read(p[:nToRead])
		if(nRead == n){
			return nRead, io.EOF
		}
		return nRead, err
	}
}

/*
type Reader interface {
	Read(p []byte)(n int, err error)
}
*/
