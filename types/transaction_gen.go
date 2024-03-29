package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Transaction) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "txContent":
			err = z.TxContent.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "TxContent")
				return
			}
		case "signature":
			z.Signature, err = dc.ReadBytes(z.Signature)
			if err != nil {
				err = msgp.WrapError(err, "Signature")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Transaction) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "txContent"
	err = en.Append(0x82, 0xa9, 0x74, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = z.TxContent.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "TxContent")
		return
	}
	// write "signature"
	err = en.Append(0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Signature)
	if err != nil {
		err = msgp.WrapError(err, "Signature")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Transaction) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "txContent"
	o = append(o, 0x82, 0xa9, 0x74, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o, err = z.TxContent.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "TxContent")
		return
	}
	// string "signature"
	o = append(o, 0xa9, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65)
	o = msgp.AppendBytes(o, z.Signature)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Transaction) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "txContent":
			bts, err = z.TxContent.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "TxContent")
				return
			}
		case "signature":
			z.Signature, bts, err = msgp.ReadBytesBytes(bts, z.Signature)
			if err != nil {
				err = msgp.WrapError(err, "Signature")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Transaction) Msgsize() (s int) {
	s = 1 + 10 + z.TxContent.Msgsize() + 10 + msgp.BytesPrefixSize + len(z.Signature)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TxContent) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "chainID":
			z.ChainID, err = dc.ReadByte()
			if err != nil {
				err = msgp.WrapError(err, "ChainID")
				return
			}
		case "nonce":
			z.Nonce, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Nonce")
				return
			}
		case "sender":
			err = z.Sender.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Sender")
				return
			}
		case "recipient":
			err = z.Recipient.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Recipient")
				return
			}
		case "value":
			z.Value, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		case "fee":
			z.Fee, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Fee")
				return
			}
		case "data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TxContent) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "chainID"
	err = en.Append(0x87, 0xa7, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44)
	if err != nil {
		return
	}
	err = en.WriteByte(z.ChainID)
	if err != nil {
		err = msgp.WrapError(err, "ChainID")
		return
	}
	// write "nonce"
	err = en.Append(0xa5, 0x6e, 0x6f, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Nonce)
	if err != nil {
		err = msgp.WrapError(err, "Nonce")
		return
	}
	// write "sender"
	err = en.Append(0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.Sender.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Sender")
		return
	}
	// write "recipient"
	err = en.Append(0xa9, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = z.Recipient.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Recipient")
		return
	}
	// write "value"
	err = en.Append(0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	// write "fee"
	err = en.Append(0xa3, 0x66, 0x65, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Fee)
	if err != nil {
		err = msgp.WrapError(err, "Fee")
		return
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TxContent) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "chainID"
	o = append(o, 0x87, 0xa7, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44)
	o = msgp.AppendByte(o, z.ChainID)
	// string "nonce"
	o = append(o, 0xa5, 0x6e, 0x6f, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint64(o, z.Nonce)
	// string "sender"
	o = append(o, 0xa6, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72)
	o, err = z.Sender.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Sender")
		return
	}
	// string "recipient"
	o = append(o, 0xa9, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74)
	o, err = z.Recipient.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Recipient")
		return
	}
	// string "value"
	o = append(o, 0xa5, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.Value)
	// string "fee"
	o = append(o, 0xa3, 0x66, 0x65, 0x65)
	o = msgp.AppendInt64(o, z.Fee)
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TxContent) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "chainID":
			z.ChainID, bts, err = msgp.ReadByteBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ChainID")
				return
			}
		case "nonce":
			z.Nonce, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Nonce")
				return
			}
		case "sender":
			bts, err = z.Sender.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sender")
				return
			}
		case "recipient":
			bts, err = z.Recipient.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Recipient")
				return
			}
		case "value":
			z.Value, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Value")
				return
			}
		case "fee":
			z.Fee, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Fee")
				return
			}
		case "data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TxContent) Msgsize() (s int) {
	s = 1 + 8 + msgp.ByteSize + 6 + msgp.Uint64Size + 7 + z.Sender.Msgsize() + 10 + z.Recipient.Msgsize() + 6 + msgp.Int64Size + 4 + msgp.Int64Size + 5 + msgp.BytesPrefixSize + len(z.Data)
	return
}
