// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: question.proto

package question

import (
	_ "github.com/tongque0/edupal/hertz_gen/cwgo/http/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenQuestionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject     string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" form:"subject" query:"subject"`
	Class       string `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty" form:"class" query:"class"`
	Difficulty  string `protobuf:"bytes,3,opt,name=difficulty,proto3" json:"difficulty,omitempty" form:"difficulty" query:"difficulty"`
	ChoiceNum   int32  `protobuf:"varint,4,opt,name=choice_num,json=choiceNum,proto3" json:"choice_num,omitempty" form:"choice_num" query:"choice_num"`               // 选择题数量
	CalcNum     int32  `protobuf:"varint,5,opt,name=calc_num,json=calcNum,proto3" json:"calc_num,omitempty" form:"calc_num" query:"calc_num"`                         // 计算题数量
	ShortAnsNum int32  `protobuf:"varint,6,opt,name=short_ans_num,json=shortAnsNum,proto3" json:"short_ans_num,omitempty" form:"short_ans_num" query:"short_ans_num"` // 简答题数量
	Knowledge   string `protobuf:"bytes,7,opt,name=knowledge,proto3" json:"knowledge,omitempty" form:"knowledge" query:"knowledge"`                                   // 知识点
}

func (x *GenQuestionReq) Reset() {
	*x = GenQuestionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenQuestionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenQuestionReq) ProtoMessage() {}

func (x *GenQuestionReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenQuestionReq.ProtoReflect.Descriptor instead.
func (*GenQuestionReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{0}
}

func (x *GenQuestionReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *GenQuestionReq) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *GenQuestionReq) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *GenQuestionReq) GetChoiceNum() int32 {
	if x != nil {
		return x.ChoiceNum
	}
	return 0
}

func (x *GenQuestionReq) GetCalcNum() int32 {
	if x != nil {
		return x.CalcNum
	}
	return 0
}

func (x *GenQuestionReq) GetShortAnsNum() int32 {
	if x != nil {
		return x.ShortAnsNum
	}
	return 0
}

func (x *GenQuestionReq) GetKnowledge() string {
	if x != nil {
		return x.Knowledge
	}
	return ""
}

type GenQuestionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genid string `protobuf:"bytes,1,opt,name=genid,proto3" json:"genid,omitempty" form:"genid" query:"genid"` // 生成id，可以根据该id查询生成题目列表
}

func (x *GenQuestionResp) Reset() {
	*x = GenQuestionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenQuestionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenQuestionResp) ProtoMessage() {}

func (x *GenQuestionResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenQuestionResp.ProtoReflect.Descriptor instead.
func (*GenQuestionResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{1}
}

func (x *GenQuestionResp) GetGenid() string {
	if x != nil {
		return x.Genid
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`                                                  // 题目类型
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`                                              // 题目标题
	Question   string `protobuf:"bytes,4,opt,name=question,proto3" json:"question,omitempty" form:"question" query:"question"`                                  // 题目内容
	Answer     string `protobuf:"bytes,5,opt,name=answer,proto3" json:"answer,omitempty" form:"answer" query:"answer"`                                          // 答案
	Tip        string `protobuf:"bytes,6,opt,name=tip,proto3" json:"tip,omitempty" form:"tip" query:"tip"`                                                      // 提示
	Parse      string `protobuf:"bytes,7,opt,name=parse,proto3" json:"parse,omitempty" form:"parse" query:"parse"`                                              // 解析
	Subject    string `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject,omitempty" form:"subject" query:"subject"`                                      // 科目
	Class      string `protobuf:"bytes,9,opt,name=class,proto3" json:"class,omitempty" form:"class" query:"class"`                                              // 年级
	Difficulty string `protobuf:"bytes,10,opt,name=difficulty,proto3" json:"difficulty,omitempty" form:"difficulty" query:"difficulty"`                         // 难度
	OwnerId    string `protobuf:"bytes,11,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty" form:"owner_id" query:"owner_id"`                    // 拥有者的用户ID
	OwnerName  string `protobuf:"bytes,12,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty" form:"owner_name" query:"owner_name"`          // 拥有者的用户名
	QuesId     string `protobuf:"bytes,13,opt,name=ques_id,json=quesId,proto3" json:"ques_id,omitempty" form:"ques_id" query:"ques_id"`                         // 题目ID
	QuesBankId string `protobuf:"bytes,14,opt,name=ques_bank_id,json=quesBankId,proto3" json:"ques_bank_id,omitempty" form:"ques_bank_id" query:"ques_bank_id"` // 题库ID
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{2}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Question) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetTip() string {
	if x != nil {
		return x.Tip
	}
	return ""
}

func (x *Question) GetParse() string {
	if x != nil {
		return x.Parse
	}
	return ""
}

func (x *Question) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Question) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *Question) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Question) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Question) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Question) GetQuesId() string {
	if x != nil {
		return x.QuesId
	}
	return ""
}

func (x *Question) GetQuesBankId() string {
	if x != nil {
		return x.QuesBankId
	}
	return ""
}

type GetQuestionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ownerid    string `protobuf:"bytes,1,opt,name=ownerid,proto3" json:"ownerid,omitempty" form:"ownerid" query:"ownerid"`             // 生成题目的id
	Quesbankid string `protobuf:"bytes,2,opt,name=quesbankid,proto3" json:"quesbankid,omitempty" form:"quesbankid" query:"quesbankid"` // 题库id
	Quesid     string `protobuf:"bytes,3,opt,name=quesid,proto3" json:"quesid,omitempty" form:"quesid" query:"quesid"`                 // 题目id
	Subject    string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty" form:"subject" query:"subject"`             // 科目
	Type       string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty" form:"type" query:"type"`                         // 题目类型
	Genid      string `protobuf:"bytes,6,opt,name=genid,proto3" json:"genid,omitempty" form:"genid" query:"genid"`                     // 生成题目的id
}

func (x *GetQuestionReq) Reset() {
	*x = GetQuestionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionReq) ProtoMessage() {}

func (x *GetQuestionReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionReq.ProtoReflect.Descriptor instead.
func (*GetQuestionReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{3}
}

func (x *GetQuestionReq) GetOwnerid() string {
	if x != nil {
		return x.Ownerid
	}
	return ""
}

func (x *GetQuestionReq) GetQuesbankid() string {
	if x != nil {
		return x.Quesbankid
	}
	return ""
}

func (x *GetQuestionReq) GetQuesid() string {
	if x != nil {
		return x.Quesid
	}
	return ""
}

func (x *GetQuestionReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *GetQuestionReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetQuestionReq) GetGenid() string {
	if x != nil {
		return x.Genid
	}
	return ""
}

type GetQuestionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty" form:"questions" query:"questions"`
	HasMore   bool        `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty" form:"has_more" query:"has_more"` // 是否还有更多题目
}

func (x *GetQuestionResp) Reset() {
	*x = GetQuestionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionResp) ProtoMessage() {}

func (x *GetQuestionResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionResp.ProtoReflect.Descriptor instead.
func (*GetQuestionResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{4}
}

func (x *GetQuestionResp) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

func (x *GetQuestionResp) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type CheckAnswerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quesid string `protobuf:"bytes,1,opt,name=quesid,proto3" json:"quesid,omitempty" form:"quesid" query:"quesid"` // 题目id
	Answer string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty" form:"answer" query:"answer"` // 答案
}

func (x *CheckAnswerReq) Reset() {
	*x = CheckAnswerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAnswerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAnswerReq) ProtoMessage() {}

func (x *CheckAnswerReq) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAnswerReq.ProtoReflect.Descriptor instead.
func (*CheckAnswerReq) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{5}
}

func (x *CheckAnswerReq) GetQuesid() string {
	if x != nil {
		return x.Quesid
	}
	return ""
}

func (x *CheckAnswerReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type CheckResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty" form:"result" query:"result"`                                // 答案是否正确
	ResultMsg string `protobuf:"bytes,2,opt,name=result_msg,json=resultMsg,proto3" json:"result_msg,omitempty" form:"result_msg" query:"result_msg"` // 答案结果信息
}

func (x *CheckResp) Reset() {
	*x = CheckResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResp) ProtoMessage() {}

func (x *CheckResp) ProtoReflect() protoreflect.Message {
	mi := &file_question_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResp.ProtoReflect.Descriptor instead.
func (*CheckResp) Descriptor() ([]byte, []int) {
	return file_question_proto_rawDescGZIP(), []int{6}
}

func (x *CheckResp) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *CheckResp) GetResultMsg() string {
	if x != nil {
		return x.ResultMsg
	}
	return ""
}

var File_question_proto protoreflect.FileDescriptor

var file_question_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x63, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x63, 0x4e,
	0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x41, 0x6e, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x69, 0x64, 0x22, 0xe5, 0x02,
	0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x75,
	0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x71, 0x75, 0x65,
	0x73, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x6e, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x42,
	0x61, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x62, 0x61, 0x6e, 0x6b,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x71, 0x75, 0x65, 0x73, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x69, 0x64, 0x22, 0x5e,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x30, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x40,
	0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x71, 0x75, 0x65, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x22, 0x42, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x4d, 0x73, 0x67, 0x32, 0x90, 0x02, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x6e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x12, 0x56, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x67, 0x65, 0x74, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x13, 0xd2, 0xc1, 0x18, 0x0f, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6e, 0x67, 0x71, 0x75, 0x65, 0x30, 0x2f, 0x65,
	0x64, 0x75, 0x70, 0x61, 0x6c, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_question_proto_rawDescOnce sync.Once
	file_question_proto_rawDescData = file_question_proto_rawDesc
)

func file_question_proto_rawDescGZIP() []byte {
	file_question_proto_rawDescOnce.Do(func() {
		file_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_proto_rawDescData)
	})
	return file_question_proto_rawDescData
}

var file_question_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_question_proto_goTypes = []interface{}{
	(*GenQuestionReq)(nil),  // 0: question.GenQuestionReq
	(*GenQuestionResp)(nil), // 1: question.GenQuestionResp
	(*Question)(nil),        // 2: question.Question
	(*GetQuestionReq)(nil),  // 3: question.GetQuestionReq
	(*GetQuestionResp)(nil), // 4: question.GetQuestionResp
	(*CheckAnswerReq)(nil),  // 5: question.CheckAnswerReq
	(*CheckResp)(nil),       // 6: question.CheckResp
}
var file_question_proto_depIdxs = []int32{
	2, // 0: question.GetQuestionResp.questions:type_name -> question.Question
	0, // 1: question.QuesService.GenQuestions:input_type -> question.GenQuestionReq
	3, // 2: question.QuesService.GetQuestions:input_type -> question.GetQuestionReq
	5, // 3: question.QuesService.CheckAnswer:input_type -> question.CheckAnswerReq
	1, // 4: question.QuesService.GenQuestions:output_type -> question.GenQuestionResp
	4, // 5: question.QuesService.GetQuestions:output_type -> question.GetQuestionResp
	6, // 6: question.QuesService.CheckAnswer:output_type -> question.CheckResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_question_proto_init() }
func file_question_proto_init() {
	if File_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenQuestionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenQuestionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_question_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_question_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuestionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_question_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuestionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_question_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAnswerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_question_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_question_proto_goTypes,
		DependencyIndexes: file_question_proto_depIdxs,
		MessageInfos:      file_question_proto_msgTypes,
	}.Build()
	File_question_proto = out.File
	file_question_proto_rawDesc = nil
	file_question_proto_goTypes = nil
	file_question_proto_depIdxs = nil
}
