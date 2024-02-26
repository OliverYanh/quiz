// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.4
// source: quiz.proto

package generated

import (
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

// 获取问卷请求消息
type GetQuizRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQuizRequest) Reset() {
	*x = GetQuizRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuizRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuizRequest) ProtoMessage() {}

func (x *GetQuizRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuizRequest.ProtoReflect.Descriptor instead.
func (*GetQuizRequest) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{0}
}

// 获取问卷响应消息
type GetQuizResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quiz *Quiz `protobuf:"bytes,1,opt,name=quiz,proto3" json:"quiz,omitempty"`
}

func (x *GetQuizResponse) Reset() {
	*x = GetQuizResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuizResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuizResponse) ProtoMessage() {}

func (x *GetQuizResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuizResponse.ProtoReflect.Descriptor instead.
func (*GetQuizResponse) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuizResponse) GetQuiz() *Quiz {
	if x != nil {
		return x.Quiz
	}
	return nil
}

// 提交问卷答案并获取结果请求消息
type SubmitQuizRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionAnswers []*QuestionAnswer `protobuf:"bytes,1,rep,name=question_answers,json=questionAnswers,proto3" json:"question_answers,omitempty"`
}

func (x *SubmitQuizRequest) Reset() {
	*x = SubmitQuizRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitQuizRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitQuizRequest) ProtoMessage() {}

func (x *SubmitQuizRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitQuizRequest.ProtoReflect.Descriptor instead.
func (*SubmitQuizRequest) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitQuizRequest) GetQuestionAnswers() []*QuestionAnswer {
	if x != nil {
		return x.QuestionAnswers
	}
	return nil
}

// 提交问卷答案并获取结果响应消息
type SubmitQuizResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score  int32  `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SubmitQuizResponse) Reset() {
	*x = SubmitQuizResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitQuizResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitQuizResponse) ProtoMessage() {}

func (x *SubmitQuizResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitQuizResponse.ProtoReflect.Descriptor instead.
func (*SubmitQuizResponse) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitQuizResponse) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SubmitQuizResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// 问卷消息类型
type Quiz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *Quiz) Reset() {
	*x = Quiz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quiz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quiz) ProtoMessage() {}

func (x *Quiz) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quiz.ProtoReflect.Descriptor instead.
func (*Quiz) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{4}
}

func (x *Quiz) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

// 问题消息类型
type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text          string    `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Answers       []*Answer `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
	CorrectAnswer string    `protobuf:"bytes,4,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[5]
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
	return file_quiz_proto_rawDescGZIP(), []int{5}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Question) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

func (x *Question) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

// 答案消息类型
type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{6}
}

func (x *Answer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Answer) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// 用户提交的问卷答案消息类型
type QuestionAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestionId string `protobuf:"bytes,1,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	AnswerId   string `protobuf:"bytes,2,opt,name=answer_id,json=answerId,proto3" json:"answer_id,omitempty"`
}

func (x *QuestionAnswer) Reset() {
	*x = QuestionAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quiz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionAnswer) ProtoMessage() {}

func (x *QuestionAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_quiz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionAnswer.ProtoReflect.Descriptor instead.
func (*QuestionAnswer) Descriptor() ([]byte, []int) {
	return file_quiz_proto_rawDescGZIP(), []int{7}
}

func (x *QuestionAnswer) GetQuestionId() string {
	if x != nil {
		return x.QuestionId
	}
	return ""
}

func (x *QuestionAnswer) GetAnswerId() string {
	if x != nil {
		return x.AnswerId
	}
	return ""
}

var File_quiz_proto protoreflect.FileDescriptor

var file_quiz_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x71, 0x75,
	0x69, 0x7a, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x71, 0x75, 0x69, 0x7a, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x69,
	0x7a, 0x52, 0x04, 0x71, 0x75, 0x69, 0x7a, 0x22, 0x54, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x10,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x0f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a,
	0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x34, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x2c, 0x0a, 0x09, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7d, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x4e, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x49, 0x64, 0x32, 0x8a, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x69, 0x7a, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x12,
	0x14, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x12, 0x17, 0x2e, 0x71,
	0x75, 0x69, 0x7a, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x71, 0x75, 0x69, 0x7a, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x61, 0x6e, 0x68, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_quiz_proto_rawDescOnce sync.Once
	file_quiz_proto_rawDescData = file_quiz_proto_rawDesc
)

func file_quiz_proto_rawDescGZIP() []byte {
	file_quiz_proto_rawDescOnce.Do(func() {
		file_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_quiz_proto_rawDescData)
	})
	return file_quiz_proto_rawDescData
}

var file_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_quiz_proto_goTypes = []interface{}{
	(*GetQuizRequest)(nil),     // 0: quiz.GetQuizRequest
	(*GetQuizResponse)(nil),    // 1: quiz.GetQuizResponse
	(*SubmitQuizRequest)(nil),  // 2: quiz.SubmitQuizRequest
	(*SubmitQuizResponse)(nil), // 3: quiz.SubmitQuizResponse
	(*Quiz)(nil),               // 4: quiz.Quiz
	(*Question)(nil),           // 5: quiz.Question
	(*Answer)(nil),             // 6: quiz.Answer
	(*QuestionAnswer)(nil),     // 7: quiz.QuestionAnswer
}
var file_quiz_proto_depIdxs = []int32{
	4, // 0: quiz.GetQuizResponse.quiz:type_name -> quiz.Quiz
	7, // 1: quiz.SubmitQuizRequest.question_answers:type_name -> quiz.QuestionAnswer
	5, // 2: quiz.Quiz.questions:type_name -> quiz.Question
	6, // 3: quiz.Question.answers:type_name -> quiz.Answer
	0, // 4: quiz.QuizService.GetQuiz:input_type -> quiz.GetQuizRequest
	2, // 5: quiz.QuizService.SubmitQuiz:input_type -> quiz.SubmitQuizRequest
	1, // 6: quiz.QuizService.GetQuiz:output_type -> quiz.GetQuizResponse
	3, // 7: quiz.QuizService.SubmitQuiz:output_type -> quiz.SubmitQuizResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_quiz_proto_init() }
func file_quiz_proto_init() {
	if File_quiz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quiz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuizRequest); i {
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
		file_quiz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuizResponse); i {
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
		file_quiz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitQuizRequest); i {
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
		file_quiz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitQuizResponse); i {
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
		file_quiz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quiz); i {
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
		file_quiz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_quiz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_quiz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionAnswer); i {
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
			RawDescriptor: file_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quiz_proto_goTypes,
		DependencyIndexes: file_quiz_proto_depIdxs,
		MessageInfos:      file_quiz_proto_msgTypes,
	}.Build()
	File_quiz_proto = out.File
	file_quiz_proto_rawDesc = nil
	file_quiz_proto_goTypes = nil
	file_quiz_proto_depIdxs = nil
}