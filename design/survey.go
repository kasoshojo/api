package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("surveys", func() {
	BasePath("/surveys")
	DefaultMedia(SurveyMedia)

	Security(JWT, func() {
		Scope("api:user")
	})

	Action("get", func() {
		Description("Get news list")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer, "Survey ID")
		})
		Response(OK, SurveyMedia)
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("list", func() {
		Description("Get news list")
		Routing(GET("/"))
		Response(OK, CollectionOf(SurveyMedia))
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("vote", func() {
		Description("Register an answer")
		Routing(POST("/:id/results"))
		Params(func() {
			Param("id", Integer, "Survey ID")
		})
		Payload(ResultPayload)
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
	})
})

var ResultPayload = Type("SurveyResultPayload", func() {
	Attribute("question_id", Integer, "ID of question")
	Attribute("answer_id", Integer, "ID of answer")
	Attribute("weight", Integer, "Weight")
	Attribute("text", String, "Answer text")
	Required("question_id")
	Required("answer_id")
})

var SurveyAnswerMedia = MediaType("application/vnd.goa.surveyanswer+json", func() {
	Description("A survey answer")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("text_en", String, "Text")
		Attribute("text_ja", String, "Text")
		Attribute("image_url", String, "image URL")
		Attribute("index", Integer, "Index")
		Attribute("open_answer", Boolean, "Open answer")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "text_en", "text_ja", "image_url", "index", "open_answer", "created_at", "updated_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("text_en")
		Attribute("text_ja")
		Attribute("image_url")
		Attribute("index")
		Attribute("open_answer")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var SurveyQuestionMedia = MediaType("application/vnd.goa.surveyquestion+json", func() {
	Description("A survey question")
	Attributes(func() {
		Attribute("id", Integer, "Unique ID")
		Attribute("text_en", String, "Text")
		Attribute("text_ja", String, "Text")
		Attribute("image_url", String, "image URL")
		Attribute("index", Integer, "Index")
		Attribute("weighted", Boolean, "Weighted")
		Attribute("num_answers", Integer, "Number of answers")
		Attribute("preselected_index", Integer, "Index of preselected answer")
		Attribute("answers", ArrayOf(SurveyAnswerMedia), "Answers")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "text_en", "text_ja", "image_url", "index", "weighted",
			"num_answers", "preselected_index", "created_at", "updated_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("text_en")
		Attribute("text_ja")
		Attribute("image_url")
		Attribute("index")
		Attribute("weighted")
		Attribute("num_answers")
		Attribute("preselected_index")
		Attribute("answers")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var SurveyMedia = MediaType("application/vnd.goa.survey+json", func() {
	Description("A survey")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique ID")
		Attribute("type", Integer, "Survey type")
		Attribute("title_en", String, "Title")
		Attribute("title_ja", String, "Title")
		Attribute("image_url", String, "image URL")
		Attribute("num_questions", Integer, "Number of questions")
		Attribute("author", String, "Author")
		Attribute("reward_id", Integer, "Reward ID")
		Attribute("insertion_code", String, "Insertion code")
		Attribute("start_date", DateTime, "Start")
		Attribute("end_date", DateTime, "End")
		Attribute("questions", ArrayOf(SurveyQuestionMedia), "Questions")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "type", "title_en", "title_ja", "image_url", "num_questions", "author",
			"insertion_code", "start_date", "end_date", "created_at", "updated_at")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")
		Attribute("type")
		Attribute("title_en")
		Attribute("title_ja")
		Attribute("image_url")
		Attribute("num_questions")
		Attribute("author")
		Attribute("reward_id")
		Attribute("insertion_code")
		Attribute("start_date")
		Attribute("end_date")
		Attribute("questions")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
