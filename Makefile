.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/partner.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/channel.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/motion.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/carousel.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/third/recommend.proto
