package crypto

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Digest) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes((z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Digest) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes((z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Digest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (z)[:])
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Digest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Digest) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (DigestLength * (msgp.ByteSize))
	return
}