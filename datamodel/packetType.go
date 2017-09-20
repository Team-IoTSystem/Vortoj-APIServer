package datamodel

const (
	LOCALPATH string = "./PacketVortoj.db"
	TABLENAME string = "packet"
	DBTYPE    string = "sqlite3"
)

type JsonPacket struct {
	ID        int16  `json:"id"`
	DeviceID  string `json:"deviceid"`
	SrcMAC    string `json:"src_mac"`
	DstMAC    string `json:"dst_mac"`
	SrcIP     string `json:"src_ip"`
	DstIP     string `json:"dst_ip"`
	SrcPort   string `json:"src_port"`
	DstPort   string `json:"dst_port"`
	SYN       bool   `json:"syn"`
	ACK       bool   `json:"ack"`
	Sequence  int64  `json:"sequence"`
	Protocol  string `json:"protocol"`
	Length    int64  `json:"length"`
	DataChank []byte `json:"datachank"`
}
type DBPacket struct {
	ID        int64  `db:"id"`
	DeviceID  string `db:"deviceid"`
	SrcMAC    string `db:"src_mac"`
	DstMAC    string `db:"dst_mac"`
	SrcIP     string `db:"src_ip"`
	DstIP     string `db:"dst_ip"`
	SrcPort   string `db:"src_port"`
	DstPort   string `db:"dst_port"`
	SYN       bool   `db:"syn"`
	ACK       bool   `db:"ack"`
	Sequence  int64  `db:"sequence"`
	Protocol  string `db:"protocol"`
	Length    int64  `db:"length"`
	DataChank []byte `db:"datachank"`
}
