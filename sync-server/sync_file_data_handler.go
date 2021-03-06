package main

import (
	"github.com/sotter/dovenet/protocol"
	log "github.com/sotter/dovenet/log"
	sbase "github.com/sotter/dovenet/base"
	"sync-files/proto"
	"github.com/golang/protobuf/proto"
	"path/filepath"
	"sync-files/base"
	"os"
)

func SyncFileDataHandler(session *Session, msg *protocol.CommMsg) error{
	// 将收到的数据写文件， 是否保持文件句柄的长时间打开？不建议
	file_data := &sync_proto.SyncFileData{}
	err := proto.Unmarshal(msg.Body, file_data)
	if err != nil {
		log.Println("SyncFileData Unmarshal fail : ", err.Error())
		return err
	}

	//创建目录，以追加的方式在一个文件的末尾写入数据，如果没有这个文件，那么递归的创建文件；
	file_name := GetRootPath() + "/" + file_data.GetFileName()
	dir, _ := filepath.Split(file_name)

	append := true
	if file_data.GetCurrentPacks() == uint32(0) {
		append = false
	}

	file_mode := file_data.GetFileMode()
	if file_mode == uint32(0) {
		file_mode = 0600
	}

	err = base.SaveDataToLocalFile(file_data.Data, dir, file_name, append, os.FileMode(file_mode))
	//判定是最后一个包
	if err == nil && file_data.GetCurrentPacks() + 1 == file_data.GetTotalPacks(){
		log.Println(sbase.FormatOutput(file_data.GetFileName(), "Complete", 60))
		//发送已经完成的
		file_complete := &sync_proto.SyncFileComplete{
			FileName : file_data.FileName,
		}

		body, err := proto.Marshal(file_complete)
		if err != nil {
			log.Println("sync_file_req Marshal fail:", err.Error())
			return err
		}

		resp := protocol.NewCommMsg(uint16(sync_proto.SYNC_Msg_SyncFileComplete), body)
		//TODO：目前Server没有客户端类型的分开管理，先从数据通道回去；-> 1. 控制通道和数据通道，使用不同的端口； 2. Server端分开管理不同的客户端
		session.TcpConn.WriteWouldBlock(resp)
	}

	return nil
}
