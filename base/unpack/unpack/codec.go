package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const MsgHeader = "00000000"

func Encode(byteBuffer io.Writer, content string) error {
	// msg header + content_len + content
	//  8				4			content_len
	// 1、写入消息头
	if err := binary.Write(byteBuffer, binary.BigEndian, []byte(MsgHeader)); err != nil {
		return err
	}
	contentLength := int32(len([]byte(content)))
	// 2、写入content length
	if err := binary.Write(byteBuffer, binary.BigEndian, contentLength); err != nil {
		return err
	}
	// 3、 写入content
	if err := binary.Write(byteBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

func Decode(byteBuffer io.Reader) (Buf []byte, err error) {
	// strings.Reader struct中记录了当前reader读取的位置， 所以第一次调用io.ReadFull之后，会继续从上一次的读取开始读
	// 读取消息头
	MessageBuf := make([]byte, len(MsgHeader))
	if _, err = io.ReadFull(byteBuffer, MessageBuf); err != nil {
		return nil, err
	}
	// 比较消息头
	if string(MessageBuf) != MsgHeader {
		return nil, errors.New("msg header error")
	}
	lengthBuffer := make([]byte, 4)
	if _, err = io.ReadFull(byteBuffer, lengthBuffer); err != nil {
		return nil, err
	}
	// 转换大端字节序
	length := binary.BigEndian.Uint32(lengthBuffer)
	Buf = make([]byte, length)
	if _, err = io.ReadFull(byteBuffer, Buf); err != nil {
		return nil, err
	}
	return
}
